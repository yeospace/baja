package baja

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type NodeParams struct{}

type TreeNode struct {
	Name  string
	Leafs []TreeNode
	Type  string
}

type NodeMeta struct {
	Title string
}

type Node struct {
	Meta *NodeMeta
	Body string

	Params *NodeParams

	Raw  string
	Path string
}

func NewNode(path string) *Node {
	n := Node{Path: path}

	return &n
}

func (n *Node) Parse() {
	content, err := ioutil.ReadFile(n.Path)
	if err != nil {
		log.Fatal("Cannot parse", n.Path)
	}

	part := strings.Split(string(content), "+++")
	if len(part) < 3 {
		log.Fatal("Not enough header/body", n.Path)
	}

	n.Meta = &NodeMeta{}
	toml.Decode(string(part[1]), n.Meta)

	n.Body = string(content[2])
}

func (n *Node) FindTheme(c *Config) {
	// Find theme
	pathComponents := strings.Split(n.Path, "/")
	pathComponents[0] = c.Theme
	pathComponents = append([]string{"theme"}, pathComponents)

}

type visitor func(path string, f os.FileInfo, err error) error

func visit(node *TreeNode) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		fmt.Printf("Visited: %s\n", path)

		if f.IsDir() {
			os.MkdirAll("./public/"+path, os.ModePerm)
			return nil
		}

		//Super simple parsing
		n := NewNode(path)
		n.Parse()
		n.FindTheme(DefaultConfig())

		return nil
	}
}

func BuildNodeTree(config *Config) *TreeNode {
	n := &TreeNode{}
	_ = filepath.Walk("./content", visit(n))
	return nil
}

func (t *TreeNode) Compile() {

}

func _template(layout, path string) error {
	out, err := ioutil.ReadFile(layout)
	if err != nil {
		return err
	}
	t, err := template.New("layoyt").Parse(string(out))
	if err != nil {
		return err
	}

	cluster, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	t, err = t.Parse(string(cluster))
	return err
}

func render(tpl *template.Template, n *Node) {
	//tpl.Execute(buf, n)
}
