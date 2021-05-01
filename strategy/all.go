package strategy

import (
	"context"
	"fmt"
)

//- All is a generic method to used at example.
type All struct {
	code int
}

func (a *All) service(ctx context.Context, b []byte) {
	str := string(b)
	fmt.Printf("This is: %s. With code equals %d \n", str, a.code)
}
