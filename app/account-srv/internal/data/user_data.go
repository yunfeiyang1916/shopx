package data

import (
	"context"
	"shopx/app/account-srv/internal/biz"
	"shopx/app/account-srv/internal/biz/entity"
	"shopx/app/account-srv/internal/biz/query"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u *userRepo) GetUser(ctx context.Context, id int64) (*entity.AccountUserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) GetUserByPassword(ctx context.Context, account, pwd string) (*entity.AccountUserInfo, error) {
	q := query.Use(u.data.db)
	a := q.AccountUserInfo
	return a.WithContext(ctx).Where(a.UserAccount.Eq(account)).First()
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
