package main

import (
	"fmt"
	"go/ast"
	"strings"
)

// checkFuncDoc examines doc comments on functions and methods
// It complains if they are missing, or not of the right form
func (f *file) checkFuncDoc(fn *ast.FuncDecl) {
	if !ast.IsExported(fn.Name.Name) {
		// func is unexported
		return
	}

	if fn.Recv != nil && len(fn.Recv.List) > 0 {
		// this is a method
		if !ast.IsExported(typeName(fn.Recv.List[0], false)) {
			// receiver is unexported
			return
		}
	}

	headerPrinted := false

	if fn.Doc == nil {
		// f.errorf(fn, 1, link(docCommentsLink), category("comments"), "exported %s %s should have comment or be unexported", kind, name)
		if !headerPrinted {
			headerPrinted = true
			printFunctionHeader(fn)
		}
		fmt.Println("Method is undocumented!")

		f.al.add(fn.Pos(), fmt.Sprintf("// FIXME: %s is undocumented", fn.Name.Name), MissingDocumentation, FuncDecl)
	} else {
		s := fn.Doc.Text()
		prefix := fn.Name.Name + " "
		if !strings.HasPrefix(s, prefix) {
			if !headerPrinted {
				headerPrinted = true
				printFunctionHeader(fn)
			}
			fmt.Println("Method is incorrectly documented!")

			f.al.add(fn.Pos(), fmt.Sprintf("// FIXME: %s is incorrectly documented", fn.Name.Name), IncorrectDocumentation, FuncDecl)
			// fmt.Printf(fn.Doc, 1, link(docCommentsLink), category("comments"), `comment on exported %s %s should be of the form "%s...\n"`, kind, name, prefix)
		}
	}

	var x int
	for _, field := range fn.Type.Params.List {
		for _, fieldName := range field.Names {
			if fn.Doc == nil || !strings.Contains(fn.Doc.Text(), fieldName.Name) {
				if !headerPrinted {
					headerPrinted = true
					printFunctionHeader(fn)
				}
				fmt.Printf("Parameter #%d %s (type %s) is undocumented!\n", x, fieldName, typeName(field, true))
				// fn.Doc.List = append(fn.Doc.List, &cmt)

				f.al.add(fn.Pos(), fmt.Sprintf("// FIXME: %s is an undocumented parameter", fieldName), MissingDocumentation, FuncDecl)
			}
			x++
		}
	}
}
