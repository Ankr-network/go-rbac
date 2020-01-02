package go_rbac

import (
	"context"
	"runtime"
	"time"

	"github.com/Ankr-network/go-rbac/rbac"
)

type RoleMgr interface {
	Add(ctx context.Context, req *RoleAddRequest) (*RoleAddResponse, error)
	Del(ctx context.Context, req *RoleDelRequest) (*RoleDelResponse, error)
	Mod(ctx context.Context, req *RoleModRequest) (*RoleModResponse, error)
	Qry(ctx context.Context, req *RoleQryRequest) (*RoleQryResponse, error)
}

type role struct {
	c rbac.RoleSrvClient
}

func roleFinalizer(r *role) {
	if err := r.c.Close(); err != nil {
		println(err.Error())
	}
}

// NewRoleMgr create role manager handler
// feature
// 1. weather to choose keep-live feature, I think about it for long time, at last I decide to start it,
// because as usual, our network is the internal and stable, if appear error when use role manager,
// please recreate it
// 2. if the network is not stable, the action of create role manager will go on
// until create it successfully.
func NewRoleMgr(cfg *ConfigMgr) RoleMgr {
	var err error
	r := &role{}
OUT:
	for {
		r.c, err = rbac.NewAnkrRoleSrvClient(cfg.Addr)
		if err != nil {
			println(err.Error())
			time.Sleep(3 * time.Second)
			continue
		} else {
			break OUT
		}
	}
	runtime.SetFinalizer(&r, roleFinalizer)
	return r
}

type RoleAddRequest struct {
	Type string
	Name string
	Memo string
}

type RoleAddResponse struct {
	Resp
}

func (r *role) Add(ctx context.Context, req *RoleAddRequest) (*RoleAddResponse, error) {
	var (
		reqBackEnd = &rbac.RoleAddRequest{}
		rsp        = &RoleAddResponse{}
	)
	reqBackEnd.Type = req.Type
	reqBackEnd.Name = req.Name
	reqBackEnd.Memo = req.Memo
	rspBackEnd, err := r.c.Add(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Desc = rspBackEnd.Rsp.Desc
	return rsp, nil
}

type RoleDelRequest struct {
	ID int64
}

type RoleDelResponse struct {
	Resp
}

func (r *role) Del(ctx context.Context, req *RoleDelRequest) (*RoleDelResponse, error) {
	var (
		reqBackEnd = &rbac.RoleDelRequest{}
		rsp        = &RoleDelResponse{}
	)
	reqBackEnd.Id = req.ID
	rspBackEnd, err := r.c.Del(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Desc = rspBackEnd.Rsp.Desc
	return rsp, nil
}

type RoleModRequest struct {
	Role
}

type RoleModResponse struct {
	Resp
}

func (r *role) Mod(ctx context.Context, req *RoleModRequest) (*RoleModResponse, error) {
	var (
		reqBackEnd = &rbac.RoleModRequest{}
		rsp        = &RoleModResponse{}
	)
	reqBackEnd.Role.Id = req.Id
	reqBackEnd.Role.Memo = req.Memo
	reqBackEnd.Role.Name = req.Name
	reqBackEnd.Role.Type = req.Type
	reqBackEnd.Role.Status = req.Status
	rspBackEnd, err := r.c.Mod(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Desc = rspBackEnd.Rsp.Desc
	return rsp, nil
}

type RoleQryRequest struct {
	QueryReq
}

type Role struct {
	Id     int64
	Name   string
	Type   string
	Memo   string
	Status int64
}

type RoleQryResponse struct {
	Resp
	Total int64
	Data  []*Role
}

func (r *role) Qry(ctx context.Context, req *RoleQryRequest) (*RoleQryResponse, error) {
	var (
		reqBackEnd = &rbac.RoleQryRequest{}
		rsp        = &RoleQryResponse{}
	)
	reqBackEnd.Type = req.Type
	reqBackEnd.Page = req.Page
	reqBackEnd.Size = req.Size
	rspBackEnd, err := r.c.Qry(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Desc = rspBackEnd.Rsp.Desc
	rsp.Total = rspBackEnd.Total
	rsp.Data = make([]*Role, len(rspBackEnd.Data))
	for i, v := range rspBackEnd.Data {
		d := &Role{}
		d.Type = v.Type
		d.Status = v.Status
		d.Name = v.Name
		d.Memo = v.Memo
		d.Id = v.Id
		rsp.Data[i] = d
	}
	return rsp, nil
}
