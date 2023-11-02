package gin

import (
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc)}
}

// only one * is allowed in a pattern
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)

	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	// parse pattern
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	// insert node
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

/*
getRoute returns the node and the params if the pattern matches
for example:
r.addRoute("GET","/users/:id)
node,params: = r.getRoute("GET","/users/123")
node.pattern="/users/:id"
params["id"]="123"
*/
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)

		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]

			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

// handle is the real handler of the request
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		//
		c.Params = params
		key := c.Method + "-" + c.Path
		//use the handler
		r.handlers[key](c)
	} else {
		c.String(404, "404 NOT FOUND: %s\n", c.Path)
	}
}
