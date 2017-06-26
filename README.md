Bones
=====

A Golang Code Doctor

Bones will parse your Go sources (using Golang's own AST & parser) and analyzes it - just like a linter - for common
coding mistakes. Unlike a linter, Bones will update your source files and insert `// FIXME` comment blocks for each
problem it detects.

## Installation

Bones requires Go 1.2 or higher. Make sure you have a working Go environment. See the [install instructions](http://golang.org/doc/install.html).

The recommended way of installation is to simply `go get github.com/muesli/bones`.

## Usage

    bones /path/to/go/project

## Development

API docs can be found [here](http://godoc.org/github.com/muesli/bones).

[![Build Status](https://secure.travis-ci.org/muesli/bones.png)](http://travis-ci.org/muesli/bones)
[![Go ReportCard](http://goreportcard.com/badge/muesli/bones)](http://goreportcard.com/report/muesli/bones)
