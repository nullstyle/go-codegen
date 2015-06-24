package codegen

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"path"
)

func Output(ctx *Context, p string, data string) error {
	op := path.Join(ctx.Dir, p)
	var out bytes.Buffer

	fmt.Fprintf(&out, "package %s\n", ctx.PackageName)
	outputImports(ctx, &out)

	out.WriteString(data)

	return ioutil.WriteFile(op, out.Bytes(), 0644)
}

func outputImports(ctx *Context, w io.Writer) {
	if len(ctx.Imports) == 0 {
		return
	}

	fmt.Fprint(w, "import (\n")

	for _, i := range ctx.Imports {
		fmt.Fprintf(w, "\t\"%s\"\n", i)
	}
	fmt.Fprint(w, ")\n")

}
