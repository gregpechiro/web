package web

type Route struct {
	method, path string
	handle       Controller
	secure       bool
}

func RouteInstance(method, path string, handle Controller, secure bool) *Route {
	return &Route{
		method: method,
		path:   path,
		handle: handle,
		secure: secure,
	}
}