package codegen

import (
	"go/token"
	"path"
	"path/filepath"
	"text/template"
)

type Context struct {
	Dir         string
	Fset        *token.FileSet
	Macros      map[string]*template.Template
	PackageName string
}

func NewContext(dir string) (*Context, error) {
	result := &Context{
		Dir:         dir,
		Fset:        token.NewFileSet(),
		PackageName: "main", // default to main
	}
	return result, result.Populate()
}

func (ctx *Context) Populate() error {
	// search directory for every template in the package
	pat := path.Join(ctx.Dir, "*.tmpl")
	paths, err := filepath.Glob(pat)

	if err != nil {
		return err
	}

	ctx.Macros = make(map[string]*template.Template)

	for _, p := range paths {
		base := path.Base(p)
		name := base[:len(base)-len(".tmpl")]

		t, err := template.New(base).ParseFiles(p)
		if err != nil {
			return err
		}

		ctx.Macros[name] = t
	}

	return nil
}
