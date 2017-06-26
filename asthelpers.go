/*
 * bones
 *     Copyright (c) 2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE
 */

package main

import (
	"fmt"
	"go/ast"
	"strings"
)

func typeName(field *ast.Field, includePtr bool) string {
	typeExpr := field.Type
	start := typeExpr.Pos() - 1
	end := typeExpr.End() - 1

	pos := fs.Position(typeExpr.Pos())
	src := srcs[pos.Filename]

	t := string(src[pos.Offset : pos.Offset+int(end-start)])
	if !includePtr && strings.HasPrefix(t, "*") {
		return t[1:]
	}

	return t
}

func methodNameWithRecv(fn *ast.FuncDecl) string {
	var s string
	if fn.Recv != nil && len(fn.Recv.List) > 0 {
		s = typeName(fn.Recv.List[0], true) + "."
	}

	return s + fn.Name.Name
}

func printFunctionHeader(fn *ast.FuncDecl) {
	s := fmt.Sprintf("Checking method %s", methodNameWithRecv(fn))
	fmt.Printf("\n%s\n", strings.Repeat("=", len(s)))
	fmt.Println(s)
	fmt.Printf("%s\n", strings.Repeat("=", len(s)))
}

func printTypeHeader(ty *ast.TypeSpec) {
	s := fmt.Sprintf("Checking type %s", ty.Name.Name)
	fmt.Printf("\n%s\n", strings.Repeat("=", len(s)))
	fmt.Println(s)
	fmt.Printf("%s\n", strings.Repeat("=", len(s)))
}
