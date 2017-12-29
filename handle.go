package sun

import (
	"net/http"
)

type GetHandle interface {
	Get(w http.ResponseWriter, r *http.Request, ps Params)
}

type PostHandle interface {
	Post(w http.ResponseWriter, r *http.Request, ps Params)
}

type DeleteHandle interface {
	Delete(w http.ResponseWriter, r *http.Request, ps Params)
}

type PutHandle interface {
	Put(w http.ResponseWriter, r *http.Request, ps Params)
}

type OptionsHandle interface {
	Options(w http.ResponseWriter, r *http.Request, ps Params)
}

type HeadHandle interface {
	Head(w http.ResponseWriter, r *http.Request, ps Params)
}

type TraceHandle interface {
	Trace(w http.ResponseWriter, r *http.Request, ps Params)
}