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

package ast_test

import (
	"github.com/biodebox/yaastr/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainer_InsertNode(t *testing.T) {
	text1 := ast.NewText([]byte(`text1`)...)
	text2 := ast.NewText([]byte(`text2`)...)
	text3 := ast.NewText([]byte(`text3`)...)
	t.Run(`empty`, func(t *testing.T) {
		c := &ast.Container{}
		text := ast.NewText([]byte(`text`)...)
		c.InsertNode(-1, text)
		if !assert.Equal(t, []ast.Node{text}, c.Children) {
			return
		}
	})
	t.Run(`middle`, func(t *testing.T) {
		c := ast.NewContainer(text1,text3)
		c.InsertNode(1, text2)
		if !assert.Equal(t, []ast.Node{text1, text2, text3}, c.Children) {
			return
		}
	})
	t.Run(`push`, func(t *testing.T) {
		c := ast.NewContainer(text1, text2)
		c.InsertNode(2, text3)
		if !assert.Equal(t, []ast.Node{text1, text2, text3}, c.Children) {
			return
		}
	})
	t.Run(`unshift`, func(t *testing.T) {
		c := ast.NewContainer(text2, text3)
		c.InsertNode(0, text1)
		if !assert.Equal(t, []ast.Node{text1, text2, text3}, c.Children) {
			return
		}
	})
	t.Run(`parent`, func(t *testing.T) {
		text := ast.NewText([]byte(`text`)...)
		if !assert.Equal(t, nil, text.GetParent()) {
			return
		}
		c := ast.NewContainer()
		c.InsertNode(0, text)
		if !assert.Equal(t, c, text.GetParent()) {
			return
		}
	})
}

func TestContainer_AppendNode(t *testing.T) {
	t.Run(`parent`, func(t *testing.T) {
		text := ast.NewText([]byte(`text`)...)
		if !assert.Equal(t, nil, text.GetParent()) {
			return
		}
		c := ast.NewContainer()
		c.AppendNode(text)
		if !assert.Equal(t, c, text.GetParent()) {
			return
		}
	})
	t.Run(`empty`, func(t *testing.T) {
		c := &ast.Container{}
		text := ast.NewText([]byte(`text`)...)
		c.AppendNode(text)
		if !assert.Equal(t, []ast.Node{text}, c.Children) {
			return
		}
	})
	t.Run(`normal`, func(t *testing.T) {
		text1 := ast.NewText([]byte(`text1`)...)
		text2 := ast.NewText([]byte(`text2`)...)
		text3 := ast.NewText([]byte(`text3`)...)
		c := ast.NewContainer(text1, text2)
		if !assert.Equal(t, []ast.Node{text1, text2}, c.Children) {
			return
		}
		c.AppendNode(text3)
		if !assert.Equal(t, []ast.Node{text1, text2, text3}, c.Children) {
			return
		}
	})
}

func TestContainer_PrependNode(t *testing.T) {
	t.Run(`parent`, func(t *testing.T) {
		text := ast.NewText([]byte(`text`)...)
		if !assert.Equal(t, nil, text.GetParent()) {
			return
		}
		c := ast.NewContainer()
		c.PrependNode(text)
		if !assert.Equal(t, c, text.GetParent()) {
			return
		}
	})
	t.Run(`empty`, func(t *testing.T) {
		c := &ast.Container{}
		text := ast.NewText([]byte(`text`)...)
		c.PrependNode(text)
		if !assert.Equal(t, []ast.Node{text}, c.Children) {
			return
		}
	})
	t.Run(`normal`, func(t *testing.T) {
		text1 := ast.NewText([]byte(`text1`)...)
		text2 := ast.NewText([]byte(`text2`)...)
		text3 := ast.NewText([]byte(`text3`)...)
		c := ast.NewContainer(text2, text3)
		if !assert.Equal(t, []ast.Node{text2, text3}, c.Children) {
			return
		}
		c.PrependNode(text1)
		if !assert.Equal(t, []ast.Node{text1, text2, text3}, c.Children) {
			return
		}
	})
}

func TestContainer_DeleteNode(t *testing.T) {
	text1 := ast.NewText([]byte(`text1`)...)
	text2 := ast.NewText([]byte(`text2`)...)
	text3 := ast.NewText([]byte(`text3`)...)
	t.Run(`empty`, func(t *testing.T) {
		c := ast.NewContainer()
		c.DeleteNode(0)
		if !assert.Equal(t, []ast.Node(nil), c.Children) {
			return
		}
	})
	t.Run(`first`, func(t *testing.T) {
		c := ast.NewContainer(text1, text2, text3)
		if !assert.Equal(t, []ast.Node{text1, text2, text3}, c.Children) {
			return
		}
		c.DeleteNode(0)
		if !assert.Equal(t, []ast.Node{text2, text3}, c.Children) {
			return
		}
	})
	t.Run(`second`, func(t *testing.T) {
		c := ast.NewContainer(text1, text2, text3)
		if !assert.Equal(t, []ast.Node{text1, text2, text3}, c.Children) {
			return
		}
		c.DeleteNode(1)
		if !assert.Equal(t, []ast.Node{text1, text3}, c.Children) {
			return
		}
	})
	t.Run(`third`, func(t *testing.T) {
		c := ast.NewContainer(text1, text2, text3)
		if !assert.Equal(t, []ast.Node{text1, text2, text3}, c.Children) {
			return
		}
		c.DeleteNode(2)
		if !assert.Equal(t, []ast.Node{text1, text2}, c.Children) {
			return
		}
	})
	t.Run(`forth`, func(t *testing.T) {
		c := ast.NewContainer(text1, text2, text3)
		if !assert.Equal(t, []ast.Node{text1, text2, text3}, c.Children) {
			return
		}
		c.DeleteNode(3)
		if !assert.Equal(t, []ast.Node{text1, text2, text3}, c.Children) {
			return
		}
	})
	t.Run(`minus one`, func(t *testing.T) {
		c := ast.NewContainer(text1, text2, text3)
		if !assert.Equal(t, []ast.Node{text1, text2, text3}, c.Children) {
			return
		}
		c.DeleteNode(-1)
		if !assert.Equal(t, []ast.Node{text1, text2, text3}, c.Children) {
			return
		}
	})
}