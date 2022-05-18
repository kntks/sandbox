package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
)

var src = `
package mypackage

import (
	"fmt"
	"unsafe"
)

type (
	myStruct struct{
		field1 string
		field2 int
	}
)

var x = myStruct{}
func main() {
	fmt.Println(unsafe.Sizeof(x))
}
`

func example() {
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	output := &ast.File{
		Name: ast.NewIdent("hogehoge"),
	}

	structList := make([]ast.Decl, 0)
	for _, node := range f.Decls {
		// fmt.Printf("%+v\n", node)
		switch node.(type) {

		case *ast.GenDecl:
			genDecl := node.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {

				switch spec.(type) {
				case *ast.TypeSpec:
					typeSpec := spec.(*ast.TypeSpec)

					fmt.Printf("Struct: name=%s\n", typeSpec.Name.Name)

					switch typeSpec.Type.(type) {
					case *ast.StructType:
						structType := typeSpec.Type.(*ast.StructType)

						structList = append(structList, genDecl)

						// format.Node(os.Stdout, fset, structType)
						for _, field := range structType.Fields.List {
							i := field.Type.(*ast.Ident)
							fieldType := i.Name

							for _, name := range field.Names {
								fmt.Printf("\tField: name=%s type=%s\n", name.Name, fieldType)
							}

						}

					}
				}
			}
		}
	}
	output.Decls = structList
	// format.Node(os.Stdout, fset, f)
	format.Node(os.Stdout, fset, output)
}

func main() {
	example()
}
