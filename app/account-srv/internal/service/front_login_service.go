package service

import (
	"context"
	"shopx/app/account-srv/internal/biz"

	pb "shopx/api/gen/account/v1"
)

type FrontLoginService struct {
	pb.UnimplementedFrontLoginServer
	uc *biz.UserUsecase
}

func NewFrontLoginService(uc *biz.UserUsecase) *FrontLoginService {
	return &FrontLoginService{uc: uc}
}

func (s *FrontLoginService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginLoginReply, error) {
	return s.uc.Login(ctx, req)
}
func (s *FrontLoginService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.LoginLoginReply, error) {
	return &pb.LoginLoginReply{}, nil
}
func (s *FrontLoginService) Logout(ctx context.Context, req *pb.RegisterRequest) (*pb.LoginLoginReply, error) {
	return &pb.LoginLoginReply{}, nil
}
