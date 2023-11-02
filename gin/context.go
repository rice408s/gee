package gin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H is a shortcut for map[string]interface{}  //gin.H
type H map[string]interface{}

type Context struct {
	//origin objects
	Writer http.ResponseWriter
	Req    *http.Request

	//request info
	Path   string
	Method string
	Params map[string]string
	//response info
	StatusCode int
}

// Param returns the value of the URL param.
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// newContext is a constructor of Context
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

// PostForm returns the first value for the named component of the POST or PUT
func (c *Context) PostForm(key string) string {
	fmt.Println(key)
	fmt.Println(c.Req.FormValue(key))
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code

	//write a status code to the response
	c.Writer.WriteHeader(code)
}

// SetHeader sets a response header with a given key and value.
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String writes a string to the response body.
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON serializes the given struct as JSON into the response body.
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	//encode the obj to json
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data writes some data into the body stream and updates the HTTP code.
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML serializes the given struct as HTML into the response body.
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
