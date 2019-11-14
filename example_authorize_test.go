package go_rbac

import "fmt"

func Example() {
	srvAddr := "127.0.0.1:6801"
	c, err := New(srvAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	rsp, err := c.Authorize(&Request{
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
