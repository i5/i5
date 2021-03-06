// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

import (
	"strings"
)

type Assign struct {
	line     uint32
	operator string
	left     Node
	right    Node
}

func (this Assign) GetType() string {
	return ASSIGN
}

func (this Assign) Debug() string {
	var result strings.Builder
	result.WriteString("(")
	result.WriteString(this.left.Debug())
	result.WriteString(" ")
	result.WriteString(this.operator)
	result.WriteString(" ")
	result.WriteString(this.right.Debug())
	result.WriteString(")")
	return result.String()
}

func (this Assign) GetLine() uint32 {
	return this.line
}

func (this Assign) Init(line uint32, operator string, left Node) Assign {
	this.line = line
	this.operator = operator
	this.left = left
	return this
}

func (this Assign) GetLeft() Node {
	return this.left
}

func (this Assign) GetRight() Node {
	return this.right
}

func (this *Assign) SetRight(right Node) {
	this.right = right
}
