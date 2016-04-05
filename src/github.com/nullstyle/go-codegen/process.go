package codegen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"io"
	"path"
	"strings"
)

// Process runs the code gen engine against `arg` using `searchPath` to lookup
// templates.
func Process(arg string, searchPath []string) error {
	if strings.HasSuffix(arg, ".go") {
		return ProcessFilePath(arg, searchPath)
	}
	return ProcessDir(arg, searchPath)
}

// ProcessDir runs the code gen engine against all files in `dir` using
// `searchPath` to lookup templates.
func ProcessDir(dir string, searchPath []string) error {
	ctx, err := NewContext(dir, searchPath)
	if err != nil {
		return err
	}

	pkgs, err := parser.ParseDir(ctx.Fset, ctx.Dir, nil, 0)

	if err != nil {
		return err
	}

	var output bytes.Buffer

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			src, err := processFile(ctx, file)

			if err != nil {
				return err
			}

			fmt.Fprintln(&output, src)
		}
	}

	return Output(ctx, "main_generated.go", output.String())
}

// ProcessFilePath runs the code gen engine against a single file, at `p` using
// `searchPath` to lookup templates.
func ProcessFilePath(p string, searchPath []string) error {
	ctx, err := NewContext(path.Dir(p), searchPath)
	if err != nil {
		return err
	}

	file, err := parser.ParseFile(ctx.Fset, p, nil, 0)

	if err != nil {
		return err
	}

	src, err := processFile(ctx, file)

	if err != nil {
		return err
	}

	base := path.Base(p)
	name := base[:len(base)-len(".go")]
	return Output(ctx, name+"_generated.go", src)
}

func processFile(ctx *Context, file *ast.File) (string, error) {
	var result bytes.Buffer
	ctx.PackageName = file.Name.Name

	for _, decl := range file.Decls {
		err := processDecl(ctx, &result, decl)
		if err != nil {
			return "", err
		}
	}

	return result.String(), nil
}

func processDecl(ctx *Context, w io.Writer, decl ast.Decl) error {
	gdp, ok := decl.(*ast.GenDecl)

	if !ok {
		return nil
	}

	for _, spec := range gdp.Specs {
		err := processSpec(ctx, w, spec)
		if err != nil {
			return err
		}
	}

	return nil
}

func processSpec(ctx *Context, w io.Writer, spec ast.Spec) error {
	tsp, ok := spec.(*ast.TypeSpec)

	if !ok {
		return nil
	}

	stp, ok := tsp.Type.(*ast.StructType)

	if !ok {
		return nil
	}

	templates, err := ExtractTemplatesFromType(ctx, stp)
	if err != nil {
		return err
	}

	for _, templateName := range templates {
		err := RunTemplate(ctx, w, templateName, tsp.Name.Name, stp)
		if err != nil {
			return err
		}
	}

	return nil
}
