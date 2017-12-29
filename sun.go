package sun

import (
	"fmt"
	"net/http"
	"sync"
)

//Context
type Context struct {
	sunspot *Sunspot
}

//End Context

type Sunspot struct {
	tree *node
	pool sync.Pool
}

func New() *Sunspot {
	sunspot := &Sunspot{tree: new(node)}
	sunspot.pool.New = func() interface{} {
		return sunspot.allocateContext()
	}
	return sunspot
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
}

func (sunspot *Sunspot) allocateContext() *Context {
	return &Context{sunspot: sunspot}
}

func (sunspot *Sunspot) Handle(path string, handle interface{}) {
	sunspot.tree.addRoute(path, handle)
}

func (sunspot *Sunspot) Run(addr ...string) (err error) {
	address := resolveAddress(addr)
	err = http.ListenAndServe(address, sunspot)
	return
}

// ServeHTTP conforms to the http.Handler interface.
func (sunspot *Sunspot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := sunspot.pool.Get().(*Context)
	c.writermem.reset(w)
	c.Request = req
	c.reset()
	path := r.URL.Path
	handle, ps, _ := sunspot.tree.getValue(path)
	if handle != nil {
		switch r.Method {
		case "OPTIONS":
			h, ok := handle.(OptionsHandle)
			if ok {
				h.Options(w, r, ps)
				return
			}
		case "GET":
			h, ok := handle.(GetHandle)
			if ok {
				h.Get(w, r, ps)
				return
			}
		case "HEAD":
			h, ok := handle.(HeadHandle)
			if ok {
				h.Head(w, r, ps)
				return
			}
		case "POST":
			h, ok := handle.(PostHandle)
			if ok {
				h.Post(w, r, ps)
				return
			}
		case "PUT":
			h, ok := handle.(PutHandle)
			if ok {
				h.Put(w, r, ps)
				return
			}
		case "DELETE":
			h, ok := handle.(DeleteHandle)
			if ok {
				h.Delete(w, r, ps)
				return
			}
		case "TRACE":
			h, ok := handle.(TraceHandle)
			if ok {
				h.Trace(w, r, ps)
				return
			}
		case "CONNECT":
			fmt.Println("CONNECT")
		default:
			fmt.Println("default")
		}
		MethodNotAllowed(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
