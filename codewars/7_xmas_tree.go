// 7kyu
// https://www.codewars.com/kata/577c349edf78c178a1000108/

package kata

func XMasTree(height int) []string {
  widest := height*2 - 1
  tree := []string{}

  for i := 1; i <= height; i++ {
    branches := i*2 - 1
    buffer := (widest - branches) / 2
    branch := ""

    for j := 0; j < buffer; j++ {
      branch += "_"
    }

    for k := 0; k < branches; k++ {
      branch += "#"
    }

    for l := 0; l < buffer; l++ {
      branch += "_"
    }

    tree = append(tree, branch)
  }

  buffer := (widest - 1) / 2
  trunk := ""

  for j := 0; j < buffer; j++ {
    trunk += "_"
  }

  trunk += "#"

  for l := 0; l < buffer; l++ {
    trunk += "_"
  }

  tree = append(tree, trunk)
  tree = append(tree, trunk)

  return tree
}
