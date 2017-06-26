/*
 * bones
 *     Copyright (c) 2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE
 */

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"

	"github.com/tsuna/gorewrite"
)

var fs = token.NewFileSet() // positions are relative to fset
var srcs = make(map[string][]byte)

// file represents a file being analyzed
type file struct {
	f        *ast.File
	pkg      *ast.Package
	lastGen  *ast.GenDecl // last GenDecl entered
	filename string
	al       annotationList
}
type walker func(ast.Node) (ast.Node, bool)

func (w walker) Rewrite(node ast.Node) (ast.Node, gorewrite.Rewriter) {
	n, v := w(node)
	if v {
		return n, w
	}
	return n, nil
}

func (f *file) walk(fn func(ast.Node) (ast.Node, bool)) ast.Node {
	return gorewrite.Rewrite(walker(fn), f.f)
}

func (f *file) nodeWalk(node ast.Node) (ast.Node, bool) {
	switch v := node.(type) {
	case *ast.GenDecl:
		if v.Tok == token.IMPORT {
			return node, false
		}
		// token.CONST, token.TYPE or token.VAR
		f.lastGen = v
		return node, true
	case *ast.FuncDecl:
		f.checkFuncDoc(v)
		// Don't proceed inside funcs
		return node, false
	case *ast.TypeSpec:
		// inside a GenDecl, which usually has the doc
		if v.Doc == nil {
			v.Doc = f.lastGen.Doc
		}

		f.checkTypeDoc(v)
		node = v
		// Don't proceed inside types
		return node, false
	case *ast.ValueSpec:
		// f.lintValueSpecDoc(v, lastGen, genDeclMissingComments)
		return node, false
	}
	return node, true
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	if len(root) == 0 {
		root = "."
	}

	pkgs, err := parser.ParseDir(fs, root, func(fi os.FileInfo) bool {
		// exclude tests
		return !strings.HasSuffix(fi.Name(), "_test.go")
	}, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	astf := make([]file, 0)
	for _, pkg := range pkgs {
		fmt.Printf("package %v\n", pkg.Name)
		for fn, f := range pkg.Files {
			fmt.Printf("file %v\n", fn)
			fl := file{
				f:        f,
				filename: fn,
			}
			srcs[fn] = readSource(fn)
			astf = append(astf, fl)
		}
	}

	// Ideally we'd use an ast.CommentMap to create the comments for us
	// In the meantime though:
	for _, fl := range astf {
		fl.walk(fl.nodeWalk)

		insertAnnotations(fl.filename, fl.al)

		// Save the modified AST
		/* var buf bytes.Buffer
		if err := format.Node(&buf, fs, fl.f); err != nil {
			panic(err)
		}

		if !bytes.Equal(buf.Bytes(), srcs[fl.filename][:len(srcs[fl.filename])-1]) {
			if err := ioutil.WriteFile(fl.filename+".fixed", buf.Bytes(), 0644); err != nil {
				log.Fatal(err)
			}
		} */
	}
}
