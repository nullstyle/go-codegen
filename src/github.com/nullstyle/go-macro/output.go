package macro

import (
	"io/ioutil"
	"path"
)

func Output(ctx *Context, p string, data string) error {
	op := path.Join(ctx.Dir, p)
	return ioutil.WriteFile(op, []byte(data), 0644)
}
