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


func compoundName(name string) string {
  return strings.ReplaceAll(name, "ane", "yl")
}

func parseToTree(smiles string) (*Node, *Node) {
  var root *Node = nil
  var stack Stack
  var longestBranchLeaf *Node = nil

  for _, char := range smiles {
    charStr := string(char)
    if charStr == "(" {
      stack.Push(stack.Head())
    } else if charStr == ")" {
      stack.Pop()
    } else if charStr == "C" {
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

func buildIUPACNomenclature(longestBranchLeaf *Node) string {
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

func GetIUPACNomenclature(smiles string) string {
  _, longestBranchLeaf := parseToTree(smiles)

  return buildIUPACNomenclature(longestBranchLeaf)
}
