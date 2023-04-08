package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	rootDir := "/"
	root := tview.NewTreeNode(rootDir).SetColor(tcell.ColorLavender)

	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	add := func(target *tview.TreeNode, path string) {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			node := tview.NewTreeNode(file.Name()).SetReference(filepath.Join(path, file.Name())).SetSelectable(file.IsDir())

			if file.IsDir() {
				node.SetColor(tcell.ColorGreen)
			}
			target.AddChild(node)
		}
	}

	add(root, rootDir)

	if err := tview.NewApplication().SetRoot(tree, true).Run(); err != nil {
		log.Fatal(err)
	}

}
