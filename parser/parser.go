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

package parser

import "github.com/biodebox/yaastr/ast"

//go:generate mockery -name "Parser"

type (
	Parser interface {
		AddProcessor(processors ...Processor)
		Parse([]byte) (ast.Node, error)
	}
	Processor func(node ast.ParentNode, data []byte, parser func(ast.ParentNode, []byte) error) (int, error)
	parser    struct {
		processors []Processor
	}
)

func New(processors ...Processor) Parser {
	return &parser{processors: processors}
}

func (p *parser) AddProcessor(processors ...Processor) {
	p.processors = append(p.processors, processors...)
}

func (p *parser) Parse(data []byte) (ast.Node, error) {
	doc := &ast.Document{}
	return doc, p.parse(doc, data)
}

func (p *parser) parse(node ast.ParentNode, data []byte) error {
	var text []byte
	index := len(node.GetChildren())
LoopMain:
	for len(data) > 0 {
		for _, processor := range p.processors {
			if offset, err := processor(node, data, p.parse); err != nil {
				return err
			} else if offset != 0 {
				if offset > len(data) {
					data = nil
				} else {
					data = data[offset:]
				}
				if len(text) > 0 {
					node.InsertNode(index, ast.NewText(text...))
					text = []byte{}
				}
				index = len(node.GetChildren())
				continue LoopMain
			}
		}
		if len(data) > 0 {
			text = append(text, data[0])
			if len(data) > 1 {
				data = data[1:]
			} else {
				data = nil
			}
		}
	}
	if len(text) > 0 {
		node.InsertNode(index, ast.NewText(text...))
	}
	return nil

}
