// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseCall(fn ast.Node) (ast.Node, error) {
	err := p.require(p.peek.Type, types.LPAREN)
	if err != nil {
		return nil, err
	}
	p.next() // '('
	if p.peek.Type == types.EOL {
		p.next()
	}
	node := ast.Call{}.Init(p.peek.Line, fn)
	args, err := p.parseList(types.RPAREN)
	if err != nil {
		return nil, err
	}
	node.SetArguments(args)
	if p.peek.Type == types.EOL {
		p.next()
	}
	err = p.require(p.peek.Type, types.RPAREN)
	if err != nil {
		return nil, err
	}
	p.next() // ')'
	return node, nil
}
