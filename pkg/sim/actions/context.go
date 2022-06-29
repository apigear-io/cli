package actions

type ActionContext struct {
	InterfaceId string
	Data        map[string]any
}

func (c *ActionContext) Get(key string) any {
	return c.Data[key]
}

func (c *ActionContext) Set(key string, value any) {
	c.Data[key] = value
}
