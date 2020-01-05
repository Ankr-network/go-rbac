package go_rbac

import (
	"context"
	"runtime"
	"time"

	"github.com/Ankr-network/go-rbac/internal/rbac"
)

type ResourceMgr interface {
	Add(ctx context.Context, req *ResourceAddRequest) (*ResourceAddResponse, error)
	Del(ctx context.Context, req *ResourceDelRequest) (*ResourceDelResponse, error)
	Mod(ctx context.Context, req *ResourceModRequest) (*ResourceModResponse, error)
	Qry(ctx context.Context, req *ResourceQryRequest) (*ResourceQryResponse, error)
}

type resource struct {
	c rbac.ResourceSrvClient
}

func resourceFinalizer(r *resource) {
	if err := r.c.Close(); err != nil {
		println(err.Error())
	}
}

// NewResourceMgr
// feature
// 1. weather to choose keep-live feature, I think about it for long time, at last I decide to take it,
// because as usual, our network is the internal and stable, if appear error when use role manager,
// please recreate it
// 2. if the network is not stable, the action of create role manager will go on
// until create it successfully.
func NewResourceMgr(cfg ConfigMgr) ResourceMgr {
	var err error
	r := &resource{}
OUT:
	for {
		r.c, err = rbac.NewAnkrResourceSrvClient(cfg.Addr)
		if err != nil {
			println(err.Error())
			time.Sleep(3 * time.Second)
			continue
		} else {
			break OUT
		}
	}
	runtime.SetFinalizer(&r, resourceFinalizer)
	return r
}

type ResourceAddRequest struct {
	Type   string
	Name   string
	Value  string
	Source string
	Memo   string
}

type ResourceAddResponse struct {
	Resp
}

func (r *resource) Add(ctx context.Context, req *ResourceAddRequest) (*ResourceAddResponse, error) {
	var (
		reqBackEnd = &rbac.ResourceAddRequest{}
		rsp        = &ResourceAddResponse{}
	)
	reqBackEnd.Memo = req.Memo
	reqBackEnd.Name = req.Name
	reqBackEnd.Type = req.Type
	reqBackEnd.Source = req.Source
	reqBackEnd.Value = req.Value
	rspBackEnd, err := r.c.Add(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Desc = rspBackEnd.Rsp.Desc
	return rsp, nil
}

type ResourceDelRequest struct {
	ID int64
}

type ResourceDelResponse struct {
	Resp
}

func (r *resource) Del(ctx context.Context, req *ResourceDelRequest) (*ResourceDelResponse, error) {
	var (
		reqBackEnd = &rbac.ResourceDelRequest{}
		rsp        = &ResourceDelResponse{}
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

type Resource struct {
	Id     int64
	Type   string
	Name   string
	Value  string
	Source string
	Memo   string
	Status int64
}

type ResourceModRequest struct {
	Resource
}

type ResourceModResponse struct {
	Resp
}

func (r *resource) Mod(ctx context.Context, req *ResourceModRequest) (*ResourceModResponse, error) {
	var (
		reqBackEnd = &rbac.ResourceModRequest{}
		rsp        = &ResourceModResponse{}
	)
	reqBackEnd.Rc.Id = req.Id
	reqBackEnd.Rc.Value = req.Value
	reqBackEnd.Rc.Type = req.Type
	reqBackEnd.Rc.Source = req.Source
	reqBackEnd.Rc.Name = req.Name
	reqBackEnd.Rc.Memo = req.Memo
	reqBackEnd.Rc.Status = req.Status
	rspBackEnd, err := r.c.Mod(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Desc = rspBackEnd.Rsp.Desc
	return rsp, nil
}

type ResourceQryRequest struct {
	QueryReq
}

type ResourceQryResponse struct {
	Resp
	Total int64
	Data  []*Resource
}

func (r *resource) Qry(ctx context.Context, req *ResourceQryRequest) (*ResourceQryResponse, error) {
	var (
		reqBackEnd = &rbac.ResourceQryRequest{}
		rsp        = &ResourceQryResponse{}
	)
	reqBackEnd.Type = req.Type
	reqBackEnd.Page = req.Page
	reqBackEnd.Size = req.Size
	rspBackEnd, err := r.c.Qry(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Desc = rspBackEnd.Rsp.Desc
	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Total = rspBackEnd.Total
	rsp.Data = make([]*Resource, len(rspBackEnd.Data))
	for i, v := range rspBackEnd.Data {
		d := &Resource{}
		d.Type = v.Type
		d.Status = v.Status
		d.Memo = v.Memo
		d.Name = v.Name
		d.Source = v.Source
		d.Value = v.Value
		d.Id = v.Id
		rsp.Data[i] = d
	}
	return rsp, nil
}
