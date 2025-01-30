package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

// analyzeFile parses a Go file and detects global environment variable declarations.
func analyzeFile(filePath string) {
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filePath, nil, parser.AllErrors)
	if err != nil {
		fmt.Printf("Error parsing %s: %v\n", filePath, err)
		return
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if decl, ok := n.(*ast.GenDecl); ok && decl.Tok == token.VAR {
			for _, spec := range decl.Specs {
				if valueSpec, ok := spec.(*ast.ValueSpec); ok {
					for _, value := range valueSpec.Values {
						if call, ok := value.(*ast.CallExpr); ok {
							if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
								if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "os" && sel.Sel.Name == "Getenv" {
									for _, name := range valueSpec.Names {
										fmt.Printf("Global env variable detected: %s in file: %s\n", name.Name, filePath)
									}
								}
							}
						}
					}
				}
			}
		}
		return true
	})
}

// scanDirectory recursively scans a directory for .go files and analyzes them.
func scanDirectory(root string) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			analyzeFile(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking directory:", err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path-to-folder>")
		return
	}

	dirPath := os.Args[1]
	fmt.Println("Scanning directory:", dirPath)
	scanDirectory(dirPath)
}
