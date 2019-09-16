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

package parser_test

import (
	"github.com/biodebox/yaastr/ast"
	"github.com/biodebox/yaastr/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

type (
	quote struct {
		*ast.Container
		double bool
	}
)

func TestParser_Parse(t *testing.T) {
	t.Run(`empty`, func(t *testing.T) {
		node, err := parser.New().Parse([]byte{})
		if !assert.NoError(t, err) {
			return
		}
		doc := &ast.Document{}
		assert.Equal(t, doc, node)
	})
	t.Run(`text`, func(t *testing.T) {
		node, err := parser.New().Parse([]byte(`quote`))
		if !assert.NoError(t, err) {
			return
		}
		doc := &ast.Document{}
		doc.AppendNode(ast.NewText([]byte(`quote`)...))
		assert.Equal(t, doc, node)
	})
	t.Run(`processor with text`, func(t *testing.T) {
		node, err := parser.New(processorQuote()).Parse([]byte(`'quote' text`))
		if !assert.NoError(t, err) {
			return
		}
		doc := &ast.Document{}
		doc.AppendNode(&quote{
			Container: ast.NewContainer(ast.NewText([]byte(`quote`)...)),
		})
		doc.AppendNode(ast.NewText([]byte(` text`)...))
		assert.Equal(t, doc, node)
	})
	t.Run(`quote with quote with text`, func(t *testing.T) {
		actualDocument, err := parser.New(processorQuote(), processorDoubleQuote()).Parse([]byte(`'quote after "double quote" before' text`))
		if !assert.NoError(t, err) {
			return
		}
		doc := &ast.Document{}
		doc.AppendNode(&quote{
			Container: ast.NewContainer(
				ast.NewText([]byte(`quote after `)...),
				&quote{
					Container: ast.NewContainer(ast.NewText([]byte(`double quote`)...)),
					double:    true,
				},
				ast.NewText([]byte(` before`)...),
			),
		})
		doc.AppendNode(ast.NewText([]byte(` text`)...))
		assert.Equal(t, doc, actualDocument)
	})
	t.Run(`out of range`, func(t *testing.T) {
		actualDocument, err := parser.New(processorOutRange()).Parse([]byte(`text`))
		if !assert.NoError(t, err) {
			return
		}
		doc := &ast.Document{}
		doc.AppendNode(ast.NewText([]byte(`te`)...))
		assert.Equal(t, doc, actualDocument)
	})
}

func processorQuote() parser.Processor {
	return parser.ProcessorByRune(
		'\'',
		'\'',
		func() ast.ParentNode {
			return &quote{
				Container: ast.NewContainer(),
			}
		},
	)
}

func processorDoubleQuote() parser.Processor {
	return parser.ProcessorByRune(
		'"',
		'"',
		func() ast.ParentNode {
			return &quote{
				Container: ast.NewContainer(),
				double:    true,
			}
		},
	)
}

func processorOutRange() parser.Processor {
	return func(node ast.ParentNode, data []byte, parser func(ast.ParentNode, []byte) error) (int, error) {
		if data[0] != 'x' {
			return 0, nil
		}
		return 10, nil
	}
}
