package codegen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
)

func Output(ctx *Context, p string, data string) error {
	op := path.Join(ctx.Dir, p)
	var out bytes.Buffer

	fmt.Fprintf(&out, "package %s\n", ctx.PackageName)
	fmt.Fprint(&out, "import (\n")

	for _, i := range ctx.Imports {
		fmt.Fprintf(&out, "\t\"%s\"\n", i)
	}
	fmt.Fprint(&out, ")\n")

	out.WriteString(data)

	return ioutil.WriteFile(op, out.Bytes(), 0644)
}
