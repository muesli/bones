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

// checkTypeDoc examines doc comments on types
// It complains if they are missing, or not of the right form
func (f *file) checkTypeDoc(ty *ast.TypeSpec) {
	if !ast.IsExported(ty.Name.Name) {
		// type is unexported
		return
	}
	headerPrinted := false

	if ty.Doc == nil {
		// there's no existing documentation
		if !headerPrinted {
			printTypeHeader(ty)
			fmt.Println("Type is undocumented!")
			// ty.Doc = &ast.CommentGroup{}
			// ty.Doc.List = append(ty.Doc.List, &cmt)

			f.al.add(ty.Pos(), fmt.Sprintf("// FIXME: %s is undocumented", ty.Name.Name), MissingDocumentation, TypeDecl)
		}
		return
	}

	// we have previously existing documentation
	s := ty.Doc.Text()
	prefix := ty.Name.Name + " "
	if !strings.HasPrefix(s, prefix) {
		if !headerPrinted {
			printTypeHeader(ty)
			fmt.Println("Type is incorrectly documented!")
			// ty.Doc.List = append(ty.Doc.List, &cmt)

			f.al.add(ty.Pos(), fmt.Sprintf("// FIXME: %s is incorrectly documented", ty.Name.Name), IncorrectDocumentation, TypeDecl)
		}
		// fmt.Printf(fn.Doc, 1, link(docCommentsLink), category("comments"), `comment on exported %s %s should be of the form "%s...\n"`, kind, name, prefix)
	}
}
