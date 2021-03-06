// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

import "strings"

type Infix struct {
	line     uint32
	left     Node
	operator string
	right    Node
}

func (this Infix) GetType() string {
	return INFIX
}

func (this Infix) Debug() string {
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

func (this Infix) GetLine() uint32 {
	return this.line
}

func (this Infix) GetLeft() Node {
	return this.left
}

func (this Infix) GetRight() Node {
	return this.right
}

func (this Infix) GetOperator() string {
	return this.operator
}

func (this Infix) Init(line uint32, operator string, left Node) Infix {
	this.line = line
	this.operator = operator
	this.left = left
	return this
}

func (this Infix) Set(line uint32, left Node, operator string, right Node) Infix {
	this.line = line
	this.left = left
	this.operator = operator
	this.right = right
	return this
}

func (this *Infix) SetRight(right Node) {
	this.right = right
}
