package biz

import (
	"context"
	pb "shopx/api/gen/account/v1"
	"shopx/app/account-srv/internal/biz/entity"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	GetUser(ctx context.Context, id int64) (*entity.AccountUserInfo, error)
	GetUserByPassword(ctx context.Context, account, pwd string) (*entity.AccountUserInfo, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (u *UserUsecase) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginLoginReply, error) {
	user, err := u.repo.GetUserByPassword(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginLoginReply{
		UserId: strconv.Itoa(int(user.UserID)),
	}, nil
}
