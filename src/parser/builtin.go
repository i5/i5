// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseBuiltin() (ast.Node, error) {
	err := p.require(p.peek.Type, types.BUILTIN)
	if err != nil {
		return nil, err
	}
	node := ast.Builtin{}.Init(p.peek.Line, p.peek.Value)
	p.next()
	return node, nil
}