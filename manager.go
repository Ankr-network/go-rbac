package go_rbac

import "context"

type Manager interface {
	AddRole(ctx context.Context, roleName string) error
	UserBindRole(ctx context.Context, user, role string) error
	AddPolicy(ctx context.Context, catalog, subject, resource, action, policy string) error
}

func NewManager() Manager {

	return nil
}

type mgr struct {
}
