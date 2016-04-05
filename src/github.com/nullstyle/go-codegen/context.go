package codegen

import (
	"go/token"
	"log"
	"path"
	"path/filepath"
	"text/template"
)

// Context represents the context in which a code generation operation is run.
type Context struct {
	Dir         string
	SearchPaths []string
	Fset        *token.FileSet
	Templates   map[string]*template.Template
	PackageName string
	Imports     map[string]bool
}

// NewContext initializes a new code generation context.
func NewContext(dir string, searchPaths []string) (*Context, error) {
	result := &Context{
		Dir:         dir,
		SearchPaths: searchPaths,
		Fset:        token.NewFileSet(),
		PackageName: "main", // default to main
	}
	return result, result.Populate()
}

// Populate fills in the rest of the context based upon the context's
// config.
func (ctx *Context) Populate() error {
	ctx.Templates = make(map[string]*template.Template)
	ctx.Imports = make(map[string]bool)

	for _, dir := range ctx.SearchPaths {
		err := ctx.searchDir(dir)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ctx *Context) searchDir(dir string) error {
	// search directory for every template in the package
	pat := path.Join(dir, "*.tmpl")
	paths, err := filepath.Glob(pat)

	if err != nil {
		return err
	}

	for _, p := range paths {
		base := path.Base(p)
		name := base[:len(base)-len(".tmpl")]

		t, err := template.New(base).ParseFiles(p)
		if err != nil {
			return err
		}

		ctx.Templates[name] = t
	}

	log.Printf("found %d templates in %s", len(paths), dir)

	return nil
}
