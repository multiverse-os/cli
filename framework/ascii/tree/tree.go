package tree

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

type EdgeType string
type Value interface{}

var (
	EdgeTypeLink EdgeType = "│"
	EdgeTypeMid  EdgeType = "├──"
	EdgeTypeEnd  EdgeType = "└──"
)

type Tree interface {
	AddNode(v Value) Tree
	AddBranch(v Value) Tree
	Branch() Tree
	FindByValue(value Value) Tree
	FindLastNode() Tree
	String() string
	Bytes() []byte
	SetValue(value Value)
}

func New() Tree {
	return &node{Value: "."}
}

type node struct {
	Root  *node
	Value Value
	Nodes []*node
}

func (self *node) String() string       { return string(self.Bytes()) }
func (self *node) SetValue(value Value) { self.Value = value }

func (self *node) FindLastNode() Tree {
	nodes := self.Nodes
	self = nodes[len(nodes)-1]
	return self
}

func (self *node) AddNode(value Value) Tree {
	self.Nodes = append(self.Nodes, &node{
		Root:  self,
		Value: value,
	})
	if self.Root != nil {
		return self.Root
	}
	return self
}

func (self *node) AddBranch(value Value) Tree {
	branch := &node{
		Value: value,
	}
	self.Nodes = append(self.Nodes, branch)
	return branch
}

func (self *node) Branch() Tree {
	self.Root = nil
	return self
}

func (self *node) FindByValue(value Value) Tree {
	for _, node := range self.Nodes {
		if reflect.DeepEqual(node.Value, value) {
			return node
		}
	}
	return nil
}

func (self *node) Bytes() []byte {
	byteBuffer := new(bytes.Buffer)
	level := 0
	var levelsEnded []int
	if self.Root == nil {
		byteBuffer.WriteString(fmt.Sprintf("%v", self.Value))
		byteBuffer.WriteByte('\n')
	} else {
		edge := EdgeTypeMid
		if len(self.Nodes) == 0 {
			edge = EdgeTypeEnd
			levelsEnded = append(levelsEnded, level)
		}
		printValues(byteBuffer, 0, levelsEnded, edge, self.Value)
	}
	if len(self.Nodes) > 0 {
		printNodes(byteBuffer, level, levelsEnded, self.Nodes)
	}
	return byteBuffer.Bytes()
}

func printNodes(writer io.Writer, level int, levelsEnded []int, nodes []*node) {
	for i, node := range nodes {
		edge := EdgeTypeMid
		if i == len(nodes)-1 {
			levelsEnded = append(levelsEnded, level)
			edge = EdgeTypeEnd
		}
		printValues(writer, level, levelsEnded, edge, node.Value)
		if len(node.Nodes) > 0 {
			printNodes(writer, level+1, levelsEnded, node.Nodes)
		}
	}
}

func printValues(writer io.Writer, level int, levelsEnded []int, edge EdgeType, value Value) {
	for i := 0; i < level; i++ {
		if isEnded(levelsEnded, i) {
			fmt.Fprint(writer, "    ")
			continue
		}
		fmt.Fprintf(writer, "%s   ", EdgeTypeLink)
	}
	fmt.Fprintf(writer, "%s %v\n", edge, value)
}

func isEnded(levelsEnded []int, level int) bool {
	for _, l := range levelsEnded {
		if l == level {
			return true
		}
	}
	return false
}
