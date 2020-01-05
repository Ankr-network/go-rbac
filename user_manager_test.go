package go_rbac

import (
	"context"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestUserMgr_Add(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	uc, err := NewUserMgr(&ConfigMgr{Addr: "192.168.39.113:30008"})
	if err != nil {
		t.Error(err)
		return
	}
	rsp, err := uc.Add(context.Background(), &UserAddRequest{
		Type: "test",
		User: strconv.Itoa(rand.Intn(10000)),
		Role: "Admin",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("code: %v desc: %v\n", rsp.Code, rsp.Desc)
}

func TestUserMgr_List(t *testing.T) {
	uc, err := NewUserMgr(&ConfigMgr{Addr: "192.168.39.113:30008"})
	if err != nil {
		t.Error(err)
		return
	}
	rsp, err := uc.List(context.Background(), &UserListRequest{Type: "hub"})
	if err != nil {
		t.Error(err)
		return
	}

	for i, v := range rsp.Data {
		t.Logf("index: %d value: %+v \n", i, v)
	}
}
