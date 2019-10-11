package main

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	file := filepath.Join(home, "src", "github.com", "shanexu", "go-playground", "helloworld", "main.go")
	src, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", src, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	spew.Dump(f)
}
