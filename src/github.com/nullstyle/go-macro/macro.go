package codegen

import (
	"errors"
	"fmt"
	"go/ast"
	"io"
	"os"
)

type MacroContext struct {
	Name        string
	MacroName   string
	PackageName string
	Ctx         *Context
	Struct      *ast.StructType
}

func (mc *MacroContext) Args() []string {
	// find the field with our macro's name
	result, err := ExtractArgs(mc.Ctx, mc.Struct, mc.MacroName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "warn: couldn't get args for %s on %s\n", mc.MacroName, mc.Name)
	}

	return result
}

func (mc *MacroContext) AddImport(name string) string {
	mc.Ctx.Imports = append(mc.Ctx.Imports, name)
	return ""
}

func RunMacro(ctx *Context, w io.Writer, macroName string, typeName string, st *ast.StructType) error {
	template, ok := ctx.Macros[macroName]
	if !ok {
		return errors.New("Could not find macro: " + macroName)
	}

	// populate the template object
	mc := &MacroContext{
		Name:        typeName,
		MacroName:   macroName,
		PackageName: ctx.PackageName,
		Ctx:         ctx,
		Struct:      st,
	}

	err := template.Execute(w, mc)

	return err
}
