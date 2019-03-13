package ctl

import (
	"context"
)

// Common 控制器公共模块
type Common struct {
	MenuCtl  *Menu  `inject:""`
	RoleCtl  *Role  `inject:""`
	UserCtl  *User  `inject:""`
	LoginCtl *Login `inject:""`
	DemoCtl  *Demo  `inject:""`
}

// LoadCasbinPolicyData 加载casbin策略数据，包括角色权限数据、用户角色数据
func (c *Common) LoadCasbinPolicyData(ctx context.Context) error {
	err := c.RoleCtl.RoleBll.LoadAllPolicy(ctx)
	if err != nil {
		return err
	}

	err = c.UserCtl.UserBll.LoadAllPolicy(ctx)
	if err != nil {
		return err
	}
	return nil
}
