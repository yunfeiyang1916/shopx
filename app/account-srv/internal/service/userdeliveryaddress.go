package service

import (
	"context"

	pb "shopx/api/gen/account/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserDeliveryAddressService struct {
	pb.UnimplementedUserDeliveryAddressServer
}

func NewUserDeliveryAddressService() *UserDeliveryAddressService {
	return &UserDeliveryAddressService{}
}

func (s *UserDeliveryAddressService) CreateUserDeliveryAddress(ctx context.Context, req *pb.CreateUserDeliveryAddressReq) (*pb.CreateUserDeliveryAddressRes, error) {
	return &pb.CreateUserDeliveryAddressRes{}, nil
}
func (s *UserDeliveryAddressService) UpdateUserDeliveryAddress(ctx context.Context, req *pb.UpdateUserDeliveryAddressReq) (*pb.UpdateUserDeliveryAddressRes, error) {
	return &pb.UpdateUserDeliveryAddressRes{}, nil
}
func (s *UserDeliveryAddressService) GetUserDeliveryAddress(ctx context.Context, req *pb.GetUserDeliveryAddressReq) (*pb.CreateUserDeliveryAddressRes, error) {
	return &pb.CreateUserDeliveryAddressRes{}, nil
}
func (s *UserDeliveryAddressService) DeleteUserDeliveryAddress(ctx context.Context, req *pb.DeleteUserDeliveryAddressReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *UserDeliveryAddressService) ListUserDeliveryAddress(ctx context.Context, req *pb.ListUserDeliveryAddressReq) (*pb.UserDeliveryAddressListRes, error) {
	return &pb.UserDeliveryAddressListRes{}, nil
}
