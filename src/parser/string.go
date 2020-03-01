// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseString() (ast.Node, error) {
	err := p.require(p.peek.Type, constants.TOKEN_STRING)
	if err != nil {
		return nil, err
	}
	node := ast.String{}.Init(p.peek.Line, p.peek.Value)
	p.next()
	return node, nil
}
