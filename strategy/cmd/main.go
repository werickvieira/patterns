package main

import (
	"context"
	"v0/strategy"
)

func main() {

	ctx := new(strategy.MyContext)
	strategy := strategy.NewStrategy()

	strategies := strategy.GetAll()
	selected := "A"

	instance := strategies[selected]
	ctx.SetContext(instance)
	ctx.Execute(context.TODO(), []byte(`Strategy A`))

	selected = "B"

	instance = strategies[selected]
	ctx.SetContext(instance)
	ctx.Execute(context.TODO(), []byte(`Strategy B`))

	selected = "A"

	instance = strategies[selected]
	ctx.SetContext(instance)
	ctx.Execute(context.TODO(), []byte(`Strategy A`))

}
