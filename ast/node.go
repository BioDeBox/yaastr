// Copyright © 2019. Vladislav Karpenko <recyger@gmail.com>
//
// All rights reserved.
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
// 3. Neither the name of the copyright holder nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package ast

//go:generate mockery -name "Node|ParentNode"

type (
	Node interface {
		GetParent() ParentNode
		SetParent(node ParentNode)
	}
	ParentNode interface {
		Node
		GetChildren() []Node
		AppendNode(...Node)
		PrependNode(...Node)
		DeleteNode(index int)
		InsertNode(int, ...Node)
	}
	Child struct {
		Parent ParentNode
	}
	Container struct {
		Child
		Children []Node
	}
	Text struct {
		Child
		Content []byte
	}
	Document struct {
		Container
	}
)

//GetParent
func (c *Child) GetParent() ParentNode {
	return c.Parent
}

//SetParent
func (c *Child) SetParent(node ParentNode) {
	c.Parent = node
}

//NewText
func NewText(content ...byte) *Text {
	return &Text{Content: content}
}

//NewContainer
func NewContainer(children ...Node) *Container {
	container := &Container{Children: children}
	setParent(container, children...)
	return container
}

//GetChildren
func (c *Container) GetChildren() []Node {
	return c.Children
}

//InsertNode
func (c *Container) InsertNode(index int, nodes ...Node) {
	setParent(c, nodes...)
	if len(c.Children) == 0 {
		c.Children = nodes
		return
	}
	if index >= len(c.Children) {
		c.Children = append(c.Children, nodes...)
		return
	}
	if index < 1 {
		c.Children = append(nodes, c.Children...)
		return
	}
	c.Children = append(c.Children[:index], append(nodes, c.Children[index:]...)...)
}

//AppendNode
func (c *Container) AppendNode(nodes ...Node) {
	setParent(c, nodes...)
	if len(c.Children) == 0 {
		c.Children = nodes
		return
	}
	c.Children = append(c.Children, nodes...)
}

//PrependNode
func (c *Container) PrependNode(nodes ...Node) {
	setParent(c, nodes...)
	if len(c.Children) == 0 {
		c.Children = nodes
		return
	}
	c.Children = append(nodes, c.Children...)
}

//DeleteNode
func (c *Container) DeleteNode(index int) {
	length := len(c.Children)
	if length == 0 || index < 0{
		return
	}
	if index == 0 {
		c.Children = c.Children[1:]
		return
	}
	if length <= index {
		return
	}
	if length - 1 == index {
		c.Children = c.Children[:index]
		return
	}
	c.Children = append(c.Children[:index], c.Children[index+1:]...)
}

func setParent(parent ParentNode, nodes ...Node) {
	for _, node := range nodes {
		node.SetParent(parent)
	}
}