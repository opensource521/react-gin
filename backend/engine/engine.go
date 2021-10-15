package engine

import (
  "fmt"
  "sort"
  "strings"
  "strconv"
)

type NodePointer *Node

type Node struct {
  parent *Node
  pos int
  children *[]NodePointer
}

type NomenclaturePart struct {
  positions string
  countName string
  name string
}

type NomenclatureParts []NomenclaturePart

func (parts NomenclatureParts) Len() int {
    return len(parts)
}

func (parts NomenclatureParts) Swap(i, j int) {
    parts[i], parts[j] = parts[j], parts[i]
}

func (parts NomenclatureParts) Less(i, j int) bool {
    return parts[i].name < parts[j].name
}


var names = [12]string{"methane", "ethane", "propane", "butane", "pentane", "hexane", "heptane", "octane", "nonane", "decane", "nonane", "decane"}

var prefixes = [12]string{"", "di", "tri", "tetra", "penta", "hexa", "hepta", "octa", "nona", "deca", "undeca", "dodeca"}


/*
 * Get compound name, e.g. methyl from methan, ethyl from ethan, ...
 */
func compoundName(name string) string {
  return strings.ReplaceAll(name, "ane", "yl")
}

/*
 * Parse tree structure from SMILES format and returns pointer to
 * root node and longest branch leaf node.
 */
func parseToTree(smiles string) (*Node, *Node) {
  var root *Node = nil
  var stack Stack   // stack will keep last nodes of the branches building in progress
  var longestBranchLeaf *Node = nil

  // Iterate each character in SMILES format and process
  for _, char := range smiles {
    charStr := string(char)
    if charStr == "(" {
      // Start of new branch so use last node as root of new branch
      // e.g. CC(CC) -> node for second character is root of new branch (CC)
      stack.Push(stack.Head())
    } else if charStr == ")" {
      // End of the branch so go back to previous branch
      stack.Pop()
    } else if charStr == "C" {
      // Create new node and link to last node of current branch
      parent := stack.Head()
      var newNode Node
      newNode.parent = parent
      newNode.pos = 1
      if parent != nil {
        newNode.pos = (*parent).pos + 1
      }

      newNodeSlice := make([]NodePointer, 0, 2)
      newNode.children = &newNodeSlice

      if parent != nil {
        newChildren := append(*(parent.children), &newNode)
        parent.children = &newChildren
        stack.Pop()
        stack.Push(&newNode)
      } else {
        stack.Push(&newNode)
      }

      if root == nil {
        root = &newNode
      }

      if (longestBranchLeaf == nil || newNode.pos > (*longestBranchLeaf).pos) {
        longestBranchLeaf = &newNode
      }
    }
  }

  return root, longestBranchLeaf
}

/*
 * Build IUPAC nomenclature for a sub-branch. Gets start position and node
 * of the sub-branch as param and returns nomenclature part for it.
 * This function just traverses to the leaf assuming sub-branch is straight chain.
 */
func buildIUPACNomenclatureFromSubBranch(startPos int, root *Node) string {
  length := 0
  node := root
  for true {
    length++
    if len(*(node.children)) > 0 {
      node = (*node.children)[0]
    } else {
      break
    }
  }

  result := compoundName(names[length - 1])
  if startPos > 1 {
    result = strconv.Itoa(startPos) + "-" + result
  }
  return result
}

/*
 * Build IUPAC nomenclature for a branch. Gets start node
 * of the branch as param and returns nomenclature part for it
 */
func buildIUPACNomenclatureFromBranch(root *Node) string {
  node := root
  length := 0
  nomenclature := ""

  for true {
    length += 1
    for i := 1; i < len(*(node.children)); i++ {
      child := (*node.children)[i]
      nomenclature += "-" + buildIUPACNomenclatureFromSubBranch(child.pos - root.pos, child)
    }

    if len(*(node.children)) > 0 {
      node = (*(node.children))[0]
    } else {
      break
    }
  }

  nomenclature = compoundName(names[length - 1]) + nomenclature

  return nomenclature
}

/*
 * Build IUPAC nomenclature for the tree parsed from SMILES format.
 * Accepts longest branch leaf node as param and returns nomenclature.
 */
func buildIUPACNomenclature(longestBranchLeaf *Node) string {
  // Get the list of longest chain to use as base
  baseBranch := make([]*Node, 0, 10)
  node := longestBranchLeaf
  idx := 0
  for true {
    baseBranch = append(baseBranch, nil)
    if len(baseBranch) > 0 {
      copy(baseBranch[1:], baseBranch)
    }
    baseBranch[0] = node
    idx++;

    if node.parent != nil {
      node = node.parent
    } else {
      break
    }
  }

  /*
   * Create map of compound name => branching positions for nomenclature parts.
   * e.g.: ("methyl" => {2,3,3}, "ethyl" => {7})
   * This map will then be used to sort these parts
   * by alphabetical order of compound name to finally make IUPAC nomenclature
   */
  var branchNomenclatures map[string][]int = make(map[string][]int)
  for i := 0; i < len(baseBranch) - 1; i++ {
    for _, branchNode := range *(baseBranch[i].children) {
      if branchNode != baseBranch[i + 1] {
        nomenclature := buildIUPACNomenclatureFromBranch(branchNode)

        _, included := branchNomenclatures[nomenclature]
        if !included {
          branchNomenclatures[nomenclature] = make([]int, 0, 3)
        }
        branchNomenclatures[nomenclature] = append(branchNomenclatures[nomenclature], baseBranch[i].pos)
      }
    }
  }

  /* Generate nomenclature from map created above */

  parts := make([]NomenclaturePart, 0, 10)
  for branchName := range branchNomenclatures {
    branchPositions := branchNomenclatures[branchName]
    var part NomenclaturePart
    part.positions = strings.Trim(strings.Replace(fmt.Sprint(branchPositions), " ", ",", -1), "[]")
    part.countName = prefixes[len(branchPositions) - 1]
    part.name = branchName
    parts = append(parts, part)
  }

  sort.Sort(NomenclatureParts(parts))

  result := ""
  first := true
  for _, part := range parts {
    if first {
      first = false
    } else {
      result += "-"
    }
    result += part.positions + "-"
    result += part.countName
    result += part.name
  }

  return result + names[longestBranchLeaf.pos - 1]
}

/*
 * public API function to get IUPAC nomenclature from SMILES format.
 */
func GetIUPACNomenclature(smiles string) string {
  _, longestBranchLeaf := parseToTree(smiles)

  return buildIUPACNomenclature(longestBranchLeaf)
}
