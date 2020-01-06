package template

import "context"

type Template struct {
	name string
	raw []byte
}

// Render interpolates the ctx Context into the template and returns the string
func (t Template) Render(ctx context.Context) string {
	return 
}
