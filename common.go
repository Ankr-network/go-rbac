package go_rbac

type Resp struct {
	Code int64
	Desc string
}

type QueryReq struct {
	Page int64
	Size int64
	Type string
}
