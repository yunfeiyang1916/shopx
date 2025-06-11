package account

import (
	"context"
	v1 "shopx/api/account/v1"
	"shopx/internal/dao"
	"shopx/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	var obj entity.AccountUserBase
	err = dao.AccountUserBase.Ctx(ctx).Where("username=?", req.Username).Scan(&obj)
	res = &v1.LoginRes{
		Token:  "",
		UserId: obj.UserId,
	}
	return
}
func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
