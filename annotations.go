/*
 * bones
 *     Copyright (c) 2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE
 */

package main

import (
	"go/ast"
	"go/token"
)

const (
	MissingDocumentation   = iota
	IncorrectDocumentation = iota
)

const (
	FuncDecl = iota
	TypeDecl = iota
)

// annotationList is a list of annotations
type annotationList []annotation

type annotation struct {
	ast.Comment

	Pos            token.Pos
	AnnotationType int
	DeclType       int
}

func (list *annotationList) add(pos token.Pos, comment string, annotationType, declType int) {
	cmt := ast.Comment{
		Text: comment,
	}
	c := annotation{cmt, pos, annotationType, declType}
	*list = append(*list, c)
}
