package application

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Router struct {
	echo *echo.Echo
	app  *Application
}

type RouterContext struct {
	app  *Application
	echo echo.Context
	http *http.Request
}

type RouterContextRequest struct {
	ctx *RouterContext
}

type RouterContextRequestParam struct {
	ctx *RouterContext
}

type RouterContextResponse struct {
	ctx *RouterContext
}

type List[T interface{}] struct {
	Hits  []T `json:"hits"`
	Total int `json:"total"`
}

type HandlerRouterFunc func(RouterContext) error
type MiddlewareFunc echo.MiddlewareFunc

func CreateRouter(app *Application) *Router {
	output := Router{}
	output.app = app
	output.echo = app.Echo
	return &output
}

func CreateRouterContext(c echo.Context, r *Router) RouterContext {
	context := RouterContext{echo: c}
	context.app = r.app
	context.http = c.Request()
	return context
}

func (r *Router) GET(path string, handler HandlerRouterFunc) *Router {
	r.echo.GET(path, func(c echo.Context) error {
		return handler(CreateRouterContext(c, r))
	})
	return r
}

func (r *Router) DELETE(path string, handler HandlerRouterFunc) *Router {
	r.echo.DELETE(path, func(c echo.Context) error {
		return handler(CreateRouterContext(c, r))
	})
	return r
}

func (r *Router) POST(path string, handler HandlerRouterFunc) *Router {
	r.echo.POST(path, func(c echo.Context) error {
		return handler(CreateRouterContext(c, r))
	})
	return r
}

func (r *Router) PUT(path string, handler HandlerRouterFunc) *Router {
	r.echo.PUT(path, func(c echo.Context) error {
		return handler(CreateRouterContext(c, r))
	})
	return r
}
