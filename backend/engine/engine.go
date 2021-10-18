package engine

import (
	"sort"
	"strconv"
	"strings"
)

/*
	Find farthest vertexes from start and chains that connect each one to start
*/
func bfs(graph Graph, start int, root int) ([]Chain) {
	var queue Queue
	var visited map[int]bool
	var prev, vertexLevelMap map[int]int
	var levelVertexMap map[int][]int
	var level, maxLevel int

	prev = make(map[int]int)
	visited = make(map[int]bool)
	vertexLevelMap = make(map[int]int)		// vertex => level
	levelVertexMap = make(map[int][]int)	// level => vertexes that has a certain level

	visited[start] = true		// To prevent visiting start vertex again
	visited[root] = true		// To prevent going to the root vertex that start is separated

	queue.Push(start)
	maxLevel = -1			// Will store level of farthest vertexes

	for !queue.IsEmpty() {
		/*
			Get the first vertex in queue and pop it
		*/
		u, _ := queue.Pop()

		/*
			Set the level of vertex u to the level of previous vertex plus 1
		*/
		level = vertexLevelMap[prev[u]] + 1
		vertexLevelMap[u] = level

		/*
			Append vertex u to map of level => vertexes
		*/
		_, ok := levelVertexMap[level]
		if ok {
			levelVertexMap[level] = append(levelVertexMap[level], u)
		} else {
			levelVertexMap[level] = []int{u}
		}

		if maxLevel < level {
			maxLevel = level		// replace maxLevel if it's smaller than level
		}

		/*
			Iterate adjancent vertexes and push them into queue
		*/
		vertexes := graph.GetAdjacentVertexes(u)
		for _, v := range vertexes {
			if !visited[v] {
				visited[v] = true	// Set vertex v as visited
				queue.Push(v)
				prev[v] = u				// Set the previous vertex of vertex v to u
			}
		}
	}

	var chains []Chain		// Will store chains that connect each farthest vertex to start

	/*
		Iterate farthest vertexes and find a chain that connects it to start vertex
	*/
	for _, vertex := range levelVertexMap[maxLevel] {
		var chain Chain

		end := vertex
		okay := true

		for okay {
			chain = append(chain, end)
			end, okay = prev[end]
		}

		chains = append(chains, chain)
	}

	return chains
}

/**
* Construct a graph from SMILES and return it
*/
func parseToGraph(smiles string) Graph {
	var graph Graph
	var stack Stack
	var count int
	count = 1

	for i := 0; i < len(smiles); i++ {
		ch := smiles[i]
		switch ch {
			case 'C':
				if !stack.IsEmpty() {
					// Add edge to the graph if the previous vertex exists and pop it
					num, _ := stack.Pop()
					graph.AddEdge(num, count)
				}
				stack.Push(count)		// Push new vertex
				count++							// Move to next vertex
				break
			case '(':
				// Push the last vertex again so that the first vertex on the new branch can join it
				stack.Push(stack.Head())
				break
			case ')':
				// End of branch so pop the last vertex of it
				stack.Pop()
				break
		}
	}

	graph.SetSize(count)	// Set the number of vertexes of a graph

	return graph
}

/*
 * Generate IUPAC representation of a compound (compound or alkyl group)
 *
 * Params
 *	- parentChain: found longest chain of a compound
 *	- locants: map from substituent name to its corresponding positions on the parent chain
 *	- suffix: 'ane' for compound and 'yl' for alkyl group
 * 	- sortedSubstituentNames: keys (substituent names) of locants sorted in alphabetical order
 */
func generateIUPAC(parentChain []int, locants map[string][]string, suffix string, sortedSubstituentNames []string) string {
	var iupac string = ""

	for _, name := range sortedSubstituentNames {
		if iupac != "" {
			iupac = iupac + "-"		// Separate substituents by hyphen
		}

		/*
			Concatenate substituent to iupac
			e.g. if name is 'methyl' and its locants are [2, 3], then '2,3-dimethyl' is concatenated
		*/
		if name[0] != '(' {
			iupac = iupac + strings.Join(locants[name], ",") + "-" + prefixes[len(locants[name]) - 1] + name
		} else {
			iupac = iupac + strings.Join(locants[name], ",") + "-" + complexPrefixes[len(locants[name]) - 1] + name
		}
	}

	iupac = iupac + names[len(parentChain) - 1] + suffix	// concatenate parent chain name to iupac. (e.g. methane, ethane... or methyl, ethyl...)

	return iupac
}

/*
 * Get all possible longest chains
 *
 * Params
 *	- graph: graph that is constructed from SMILES
 *	- start: start vertex of the chain
 *	- root: vertex on the parent chain where the chain separated
 *	- isMainChain: true if working on outest chain of SMILES, otherwise false.
 *			e.g. CC(CCC)CC(C)C		the main chain is [1, 2, 6, 7, 9], other side chains are [3, 4, 5] and [8]
 *					 12 345 67 8 9
 */
func getLongestChains(graph Graph, start int, root int, isMainChain bool) []Chain {
	chains := bfs(graph, start, root)		// find the farthest vertexes from the start vertex and get chains that connect each vertex to start

	if isMainChain {
		/*
		 If working on the main chain, apply bfs on each found vertex again to find longest chains
		*/
		var longestChains = []Chain{}

		for _, chain := range chains {
			newChains := bfs(graph, chain[0], root)
			longestChains = append(longestChains, newChains...)
		}

		return longestChains
	}	else {
		return chains
	}
}

/*
 * Recursively generate compound and return it
 *
 * Params
 *	- graph: graph that is constructed from SMILES
 *	- start: start vertex of the chain
 *	- root: vertex on the parent chain where the chain separated
 *	- isMainChain: true if working on outest chain of SMILES, otherwise false.
 *			e.g. CC(CCC)CC(C)C		the main chain is [1, 2, 6, 7, 9], other side chains are [3, 4, 5] and [8]
 *					 12 345 67 8 9
 */
func getIUPACRecursively(graph Graph, start int, root int, isMainChain bool) Compound {
	longestChains := getLongestChains(graph, start, root, isMainChain)	// Find possible parent chains
	compounds := []Compound{}		// Will store candidate compounds. Each compound is generated from each chain

	for _, chain := range longestChains {
		var locants1, locants2 map[string][]string
		var sum1, sum2 int
		var compound Compound

		compound.SetIUPAC("")

		/*
			Variables locants1 and locants2 are maps where
			key is IUPAC of a substituent and value is its position array.
			Positions are calculated from left and right, that's why two variables are needed
		*/
		locants1 = make(map[string][]string)	
		locants2 = make(map[string][]string)

		sum1, sum2 = 0, 0		// Sum of positions of substituents. Calculated from left and right

		for i, u := range chain {
			vertexes := graph.GetAdjacentVertexes(u)

			for _, v := range vertexes {
				if !isInArray(chain, v) && v != root {		// Check if vertex is a start of new substituent (side chain)
					child := getIUPACRecursively(graph, v, u, false)	// Get new substituent
					compound.AppendSubstituent(child)					// Append new substituent

					iupac := child.GetIUPAC()
					_, ok := locants1[iupac]

					/*
					 Append position of new substituent to locants map
					 Add position to the sum
					*/
					if ok {		// Check if the substituent of iupac has already appeared
						locants1[iupac] = append(locants1[iupac], strconv.Itoa(i + 1))
						sum1 += (i + 1)
						locants2[iupac] = append([]string{strconv.Itoa(len(chain) - i)}, locants2[iupac]...)
						sum2 += (len(chain) - i)
					} else {
						locants1[iupac] = []string{strconv.Itoa(i + 1)}
						sum1 += (i + 1)
						locants2[iupac] = []string{strconv.Itoa(len(chain) - i)}
						sum2 += (len(chain) - i)
					}
				}
			}
		}

		/*
			Sort substituent IUPACs by representing name.
				e.g. When substituent IUPACs are [methyl, ethyl, propyl, (2-methylpropyl)]
					result is [ethyl, methyl, (2-methylpropyl), propyl]
			
			Representing name is substring of IUPAC which starts from the first letter of IUPAC
				e.g. When IUPAC is "(3-ethyl-2-methylpropyl)", then its representing name is "ethyl-2-methylpropyl)"
		*/
		substituentIUPACs := make([]string, 0, len(locants1))
		for name := range locants1 {
			substituentIUPACs = append(substituentIUPACs, name)
		}

		sort.Slice(substituentIUPACs, func (i, j int) bool {
			str1 := ltrimNonLetter(substituentIUPACs[i])	// Generate representing name of IUPAC
			str2 := ltrimNonLetter(substituentIUPACs[j])

			return str1 < str2
		})

		/*
			Define suffix.
			In case of main chain, it's "ane", otherwise "yl"
		*/
		suffix := "yl"
		if isMainChain {
			suffix = "ane"
		}

		/*
			Choose right one among two candidates which calculated positions from left and right
		*/
		if sum1 < sum2 {	// Case when the sum of positions calculated from left is smaller than sum of those calculated from right
			compound.SetIUPAC(generateIUPAC(chain, locants1, suffix, substituentIUPACs))
		} else if sum1 > sum2 {		// Case when left sum > right sum
			compound.SetIUPAC(generateIUPAC(chain, locants2, suffix, substituentIUPACs))
		} else {	// Case when left sum is same as right sum
			/*
				Iterates IUPAC and choose one where sum of positions of IUPAC is smaller at first
			*/
			for _, name := range substituentIUPACs {
				s1, s2 := sum(locants1[name]), sum(locants2[name])
				if s1 < s2 {
					compound.SetIUPAC(generateIUPAC(chain, locants1, suffix, substituentIUPACs))
					break
				} else if s1 > s2 {
					compound.SetIUPAC(generateIUPAC(chain, locants2, suffix, substituentIUPACs))
					break
				}
			}

			/*
				If two candidates are same, choose the first one as default
			*/
			if compound.GetIUPAC() == "" {
				compound.SetIUPAC(generateIUPAC(chain, locants1, suffix, substituentIUPACs))
			}
		}

		/*
			If the chain is complex and not main chain, add brackets
		*/
		if !isMainChain && len(locants1) > 0 {
			compound.SetIUPAC("(" + compound.GetIUPAC() + ")")
		}

		compounds = append(compounds, compound)		// Append new compound to candidate compounds array
	}

	/*
		Choose the right compound
	*/
	var result Compound
	result = compounds[0]		// Set first compound to result as default

	for _, compound := range compounds {
		if result.NumberOfSubstituents() < compound.NumberOfSubstituents() {
			/*
				If number of substituents of result is smaller than that of compound, then replace it
			*/
			result = compound
		} else if result.NumberOfSubstituents() < compound.NumberOfSubstituents() && 
			result.NumberOfComplexSubstituents() > compound.NumberOfComplexSubstituents() {
			/*
				If number of substituents of result and compound are same, then check the number of complex substituents.
				Choose the one that has bigger number
			*/
			result = compound
		}
	}

	return result
}

// Return IUPAC of SMILES. Entry point
func GetIUPAC(smiles string) string {
	var graph Graph
	var compound Compound

	graph = parseToGraph(smiles)

	compound = getIUPACRecursively(graph, 1, 0, true)

	return compound.GetIUPAC()
}