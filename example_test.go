package go_rbac

import (
	"context"
	"fmt"
)

// this example shows how to use this package to authorize sth.
func ExampleUse_authorize() {
	srvAddr := "127.0.0.1:6801"
	c, err := New(srvAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	rsp, err := c.Authorize(context.Background(), &Request{
		Subject:  "ankr",
		Resource: "/v1/path/to/rc",
		Action:   "read",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", rsp.OK)
	//Output:
	//true
}
