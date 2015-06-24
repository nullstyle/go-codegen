# `go-codegen`, a simple code generation system

`go-codegen` is a simple template-based code generation system for go.  By
annotating structs with special-format fields, go-codegen will generate code
based upon templates provided alongside your package.

## Example usage

go-codegen works by building a catalog of available code templates, and then any
anonymous fields on a struct whose type match a template name will be invoked
with a struct that provides information about the struct that the template is
being invoked upon.

Take for example, the following go file:

```go
package main

import "fmt"

 // cmd is a template.  Blank interfaces are good to use for targeting templates
 // as they do not affect the compiled package.  
type cmd interface{
  Execute() interface{}, error
}

type HelloCommand struct {
  // HelloCommand is invoked
  cmd
  Name string
}

func main() {
  cmd := HelloCommand{Name:"You"}
  fmt.Println(cmd.MustExecute())
}
```

TODO


### Arguments

TODO

## Finding templates

At present, a template is searched for within the same directory as a struct
invoking the template, using a name of the form `TemplateName.tmpl`.  For
example, invoking go-codegen in a directory that has a file named "FooBar.tmpl"
will cause that template to be invoked for any struct that has an anonymous
field with a type of `FooBar`.  `FooBar` may be any type: interface, struct or
otherwise.

## Notes on template invocation
