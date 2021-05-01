package strategy

import "context"

//- Invoke define a behavior default to entities.
type Invoke interface {
	service(context.Context, []byte)
}

//- MyContext knows who is current context to be invoked.
type MyContext struct {
	entity Invoke
}

func (c *MyContext) SetContext(entity Invoke) {
	c.entity = entity
}

func (c *MyContext) Execute(ctx context.Context, b []byte) {
	c.entity.service(ctx, b)
}
