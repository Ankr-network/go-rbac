package go_rbac

import (
	"context"
	"runtime"
	"time"

	"github.com/Ankr-network/go-rbac/rbac"
	"github.com/golang/glog"
)

type UserMgr interface {
	Add(context.Context, *UserAddRequest) (*UserAddResponse, error)
	Del(context.Context, *UserDelRequest) (*UserDelResponse, error)
	Mod(context.Context, *UserModRequest) (*UserModResponse, error)
	Qry(context.Context, *UserQryRequest) (*UserQryResponse, error)
	List(context.Context, *UserListRequest) (*UserListResponse, error)
}

type userMgr struct {
	c rbac.UserSrvClient
}

func userFinalizer(u *userMgr) {
	if err := u.c.Close(); err != nil {
		glog.Error(err)
	}
}

func (u *userMgr) Add(ctx context.Context, req *UserAddRequest) (*UserAddResponse, error) {
	var (
		reqBackEnd = &rbac.UserAddRequest{}
		rsp        = &UserAddResponse{}
	)
	reqBackEnd.Role = req.Role
	reqBackEnd.User = req.User
	reqBackEnd.Type = req.Type
	rspBackEnd, err := u.c.Add(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Desc = rspBackEnd.Rsp.Desc

	return rsp, nil
}

func (u *userMgr) Del(ctx context.Context, req *UserDelRequest) (*UserDelResponse, error) {
	var (
		reqBackEnd = &rbac.UserDelRequest{}
		rsp        = &UserDelResponse{}
	)
	reqBackEnd.Id = req.Id
	rspBackEnd, err := u.c.Del(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Desc = rspBackEnd.Rsp.Desc

	return rsp, nil
}

func (u *userMgr) Mod(ctx context.Context, req *UserModRequest) (*UserModResponse, error) {
	var (
		reqBackEnd = &rbac.UserModRequest{}
		rsp        = &UserModResponse{}
	)
	reqBackEnd.User.Id = req.User.Id
	reqBackEnd.User.User = req.User.User
	reqBackEnd.User.Type = req.User.Type
	reqBackEnd.User.Role = req.User.Role
	reqBackEnd.User.Status = req.User.Status

	rspBackEnd, err := u.c.Mod(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Desc = rspBackEnd.Rsp.Desc

	return rsp, nil
}

func (u *userMgr) Qry(ctx context.Context, req *UserQryRequest) (*UserQryResponse, error) {
	var (
		reqBackEnd = &rbac.UserQryRequest{}
		rsp        = &UserQryResponse{}
	)
	reqBackEnd.Type = req.Type
	reqBackEnd.Num = req.Num
	reqBackEnd.Page = req.Page

	rspBackEnd, err := u.c.Qry(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}

	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Desc = rspBackEnd.Rsp.Desc
	rsp.Total = rspBackEnd.Total
	rsp.Data = make([]*User, len(rspBackEnd.Data))
	for i, v := range rspBackEnd.Data {
		usr := &User{}

		usr.User = v.User
		usr.Type = v.Type
		usr.Status = v.Status
		usr.Role = v.Role
		usr.Id = v.Id

		rsp.Data[i] = usr
	}

	return rsp, nil
}

func (u *userMgr) List(ctx context.Context, req *UserListRequest) (*UserListResponse, error) {
	var (
		reqBackEnd = &rbac.UserListRequest{}
		rsp        = &UserListResponse{}
	)
	reqBackEnd.Type = req.Type
	rspBackEnd, err := u.c.List(ctx, reqBackEnd)
	if err != nil {
		return nil, err
	}
	rsp.Code = rspBackEnd.Rsp.Code
	rsp.Desc = rspBackEnd.Rsp.Desc

	rsp.Data = make([]*User, len(rspBackEnd.Data))
	for i, v := range rspBackEnd.Data {
		usr := &User{}

		usr.User = v.User
		usr.Type = v.Type
		usr.Status = v.Status
		usr.Role = v.Role
		usr.Id = v.Id

		rsp.Data[i] = usr
	}

	return rsp, nil
}

func NewUserMgr(cfg *ConfigMgr) (UserMgr, error) {
	var err error
	u := &userMgr{}
OUT:
	for {
		u.c, err = rbac.NewAnkrUserSrvClient(cfg.Addr)
		if err != nil {
			println(err.Error())
			time.Sleep(3 * time.Second)
			continue
		} else {
			break OUT
		}
	}
	runtime.SetFinalizer(u, userFinalizer)
	return u, nil
}

type UserAddRequest struct {
	Type string
	User string
	Role string
}

type UserAddResponse struct {
	Resp
}

type UserDelRequest struct {
	Id int64
}

type UserDelResponse struct {
	Resp
}

type User struct {
	Id     int64
	Type   string
	User   string
	Role   string
	Status int64
}

type UserModRequest struct {
	User *User
}

type UserModResponse struct {
	Resp
}

type UserQryRequest struct {
	Page int64
	Num  int64
	Type string
}

type UserQryResponse struct {
	Resp
	Data  []*User
	Total int64
}

type UserListRequest struct {
	Type string
}

type UserListResponse struct {
	Resp
	Data []*User
}
