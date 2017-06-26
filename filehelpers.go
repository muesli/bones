package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func readSource(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = append(b, '\n')
	return b
}

func insertAnnotations(filename string, annotations []annotation) {
	if len(annotations) > 0 {
		input, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalln(err)
		}
		lines := strings.Split(string(input), "\n")

		for x, c := range annotations {
			// fmt.Println("New Comments:", c.Text, fs.Position(c.Line).Line)
			newline := fs.Position(c.Pos).Line - 1
			lines = append(lines[:newline+x], append([]string{c.Text}, lines[newline+x:]...)...)
		}

		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(filename, []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%s has been fixed\n", filename)
	}
}
