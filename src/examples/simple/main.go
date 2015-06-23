package foo

import (
	"fmt"
	"net/http"
)

//go:generate go-macro

type Action struct{}

func (action *Action) Prepare(w http.ResponseWriter, r *http.Request) {}
func (action *Action) Execute(a interface{})                          {}

type MyCustomAction struct {
	Action `macro:"arg1,second_atg"`
}

func main() {
	fmt.Println(MyCustomAction{})
}
