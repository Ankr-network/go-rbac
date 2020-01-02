package go_rbac

import (
	"context"
	"testing"
)

func TestRoleBindResource_List(t *testing.T) {
	rc := NewRoleBindResourceMgr(ConfigMgr{Addr: "127.0.0.1:6802"})

	if rc == nil {
		return
	}

	rsp, err := rc.List(context.TODO(), &RoleBindingResourceListRequest{Type: "hub"})
	if err != nil {
		t.Error(err)
		return
	}
	for i, v := range rsp.Data {
		t.Logf("index: %d value: %+v \n", i, v)
	}
}
