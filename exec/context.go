package exec

type Context struct {
	data   map[string]any
	parent *Context
}

func NewContext(data map[string]any) *Context {
	return &Context{data: data}
}

func EmptyContext() *Context {
	return &Context{data: map[string]any{}}
}

func (ctx *Context) Has(name string) bool {
	_, exists := ctx.data[name]
	if !exists && ctx.parent != nil {
		return ctx.parent.Has(name)
	}
	return exists
}

func (ctx *Context) Get(name string) (any, bool) {
	value, exists := ctx.data[name]
	if exists {
		return value, true
	} else if ctx.parent != nil {
		return ctx.parent.Get(name)
	} else {
		return nil, false
	}
}

func (ctx *Context) Set(name string, value any) {
	ctx.data[name] = value
}

func (ctx *Context) Inherit() *Context {
	return &Context{
		data:   map[string]any{},
		parent: ctx,
	}
}

// Update updates this context with the key/value pairs from a map.
func (ctx *Context) Update(other map[string]any) *Context {
	for k, v := range other {
		ctx.data[k] = v
	}
	return ctx
}

// Merge updates this context with the key/value pairs from another context.
func (ctx *Context) Merge(other *Context) *Context {
	return ctx.Update(other.data)
}
