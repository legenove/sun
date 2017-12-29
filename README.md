# 例子
`
package main

import (
	"fmt"
	"net/http"
	"sun"
)

type IndexHandler struct{}

func (i *IndexHandler) Get(w http.ResponseWriter, r *http.Request, ps sun.Params) {
	fmt.Fprintf(w, "Hello myroute!")
}

func main() {
	r := sun.New()
	r.Handle("/", new(IndexHandler))
	r.Handle("/2/", new(IndexHandler))
	r.Run("0.0.0.0:8888")
}
`
