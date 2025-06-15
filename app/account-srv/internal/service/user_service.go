package service

import (
	"context"

	pb "shopx/api/gen/account/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUser(ctx context.Context, req *emptypb.Empty) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) BindMobile(ctx context.Context, req *pb.BindMobileRequest) (*pb.LoginLoginReply, error) {
	return &pb.LoginLoginReply{}, nil
}
func (s *UserService) UnBindMobile(ctx context.Context, req *pb.BindMobileRequest) (*pb.LoginLoginReply, error) {
	return &pb.LoginLoginReply{}, nil
}
func (s *UserService) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *UserService) ChangePassword(ctx context.Context, req *pb.ResetPasswordRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *UserService) Certificate(ctx context.Context, req *pb.CertificateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
