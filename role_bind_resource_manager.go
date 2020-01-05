package go_rbac

import (
	"context"
	"runtime"
	"time"

	"github.com/Ankr-network/go-rbac/internal/rbac"
)

type RoleBindingResourceListRequest struct {
	Type string
}

type RoleBindingResourceListResponse struct {
	Resp
	Data []*RoleBindingResource
}

type RoleBindingResource struct {
	Id       int64
	Type     string
	Subject  string
	Resource string
	Action   string
	Policy   string
	Status   int64
}

type RoleBindResourceMgr interface {
	Add(ctx context.Context, req *RoleBindingResourceAddRequest) (*RoleBindingResourceAddResponse, error)
	Del(ctx context.Context, req *RoleBindingResourceDelRequest) (*RoleBindingResourceDelResponse, error)
	Mod(ctx context.Context, req *RoleBindingResourceModRequest) (*RoleBindingResourceModResponse, error)
	Qry(ctx context.Context, req *RoleBindingResourceQryRequest) (*RoleBindingResourceQryResponse, error)
	List(ctx context.Context, req *RoleBindingResourceListRequest) (*RoleBindingResourceListResponse, error)
}

type roleBindResource struct {
	c rbac.RoleBindingResourceSrvClient
}

func (r *roleBindResource) List(ctx context.Context, req *RoleBindingResourceListRequest) (*RoleBindingResourceListResponse, error) {
	var (
		reqBackEnd = &rbac.RoleBindingResourceListRequest{}
		rsp        = &RoleBindingResourceListResponse{}
	)
	reqBackEnd.Type = req.Type
	rspBackEnd, err := r.c.List(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Desc = rspBackEnd.Rsp.Desc
	data := make([]*RoleBindingResource, len(rspBackEnd.Data))
	for index, v := range rspBackEnd.Data {
		d := &RoleBindingResource{}
		d.Status = v.Status
		d.Action = v.Action
		d.Id = v.Id
		d.Policy = v.Policy
		d.Resource = v.Resource
		d.Subject = v.Subject
		d.Type = v.Type
		data[index] = d
	}
	rsp.Data = data
	return rsp, nil
}

func rbrFinalizer(r *roleBindResource) {
	if err := r.c.Close(); err != nil {
		println(err.Error())
	}
}

func NewRoleBindResourceMgr(cfg ConfigMgr) RoleBindResourceMgr {
	var (
		err error
		rbr = &roleBindResource{}
	)
	rbr.c, err = rbac.NewAnkrRoleBindingResourceSrvClient(cfg.Addr)

OUT:
	for {
		rbr.c, err = rbac.NewAnkrRoleBindingResourceSrvClient(cfg.Addr)
		if err != nil {
			println(err.Error())
			time.Sleep(3 * time.Second)
			continue
		} else {
			break OUT
		}
	}
	runtime.SetFinalizer(rbr, rbrFinalizer)

	return rbr
}

type RoleBindingResourceAddRequest struct {
	Type     string
	Subject  string
	Resource string
	Action   string
	Policy   string
}

type RoleBindingResourceAddResponse struct {
	Resp
}

func (r *roleBindResource) Add(ctx context.Context, req *RoleBindingResourceAddRequest) (*RoleBindingResourceAddResponse, error) {
	var (
		reqBackEnd = &rbac.RoleBindingResourceAddRequest{}
		rsp        = &RoleBindingResourceAddResponse{}
	)
	reqBackEnd.Type = req.Type
	reqBackEnd.Resource = req.Resource
	reqBackEnd.Policy = req.Policy
	reqBackEnd.Action = req.Action
	reqBackEnd.Subject = req.Subject
	rspBackEnd, err := r.c.Add(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Desc = rspBackEnd.Rsp.Desc

	return rsp, nil
}

type RoleBindingResourceDelRequest struct {
	Id int64
}

type RoleBindingResourceDelResponse struct {
	Resp
}

func (r *roleBindResource) Del(ctx context.Context, req *RoleBindingResourceDelRequest) (*RoleBindingResourceDelResponse, error) {
	var (
		reqBackEnd = &rbac.RoleBindingResourceDelRequest{}
		rsp        = &RoleBindingResourceDelResponse{}
	)
	reqBackEnd.Id = req.Id
	rspBackEnd, err := r.c.Del(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Desc = rspBackEnd.Rsp.Desc

	return rsp, nil
}

type Policy struct {
	Id       int64
	Type     string
	Subject  string
	Resource string
	Action   string
	Permit   string
	Status   int64
}

type RoleBindingResourceModRequest struct {
	Policy
}

type RoleBindingResourceModResponse struct {
	Resp
}

func (r *roleBindResource) Mod(ctx context.Context, req *RoleBindingResourceModRequest) (*RoleBindingResourceModResponse, error) {
	var (
		reqBackEnd = &rbac.RoleBindingResourceModRequest{}
		rsp        = &RoleBindingResourceModResponse{}
	)
	reqBackEnd.Role.Id = req.Id
	reqBackEnd.Role.Subject = req.Subject
	reqBackEnd.Role.Action = req.Action
	reqBackEnd.Role.Policy = req.Permit
	reqBackEnd.Role.Resource = req.Resource
	reqBackEnd.Role.Type = req.Type
	reqBackEnd.Role.Status = req.Status
	rspBackEnd, err := r.c.Mod(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Desc = rspBackEnd.Rsp.Desc
	rsp.Code = rsp.Resp.Code
	return rsp, nil
}

type RoleBindingResourceQryRequest struct {
	QueryReq
}

type RoleBindingResourceQryResponse struct {
	Resp
	Total int64
	Data  []*Policy
}

func (r *roleBindResource) Qry(ctx context.Context, req *RoleBindingResourceQryRequest) (*RoleBindingResourceQryResponse, error) {
	var (
		reqBackEnd = &rbac.RoleBindingResourceQryRequest{}
		rsp        = &RoleBindingResourceQryResponse{}
	)
	reqBackEnd.Type = req.Type
	reqBackEnd.Page = req.Page
	reqBackEnd.Num = req.Size
	rspBackEnd, err := r.c.Qry(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Desc = rspBackEnd.Rsp.Desc
	rsp.Total = rspBackEnd.Total
	rsp.Data = make([]*Policy, len(rspBackEnd.Data))
	for i, v := range rspBackEnd.Data {
		d := &Policy{}
		d.Type = v.Type
		d.Status = v.Status
		d.Resource = v.Resource
		d.Action = v.Action
		d.Id = v.Id
		d.Permit = v.Policy
		d.Subject = v.Subject
		rsp.Data[i] = d
	}
	return rsp, nil
}
