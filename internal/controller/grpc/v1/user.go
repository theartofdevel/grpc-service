package v1

import (
	"context"

	proto_user_model "github.com/theartofdevel/grpc-contracts/gen/go/user_service/model/v1"
	proto_user_service "github.com/theartofdevel/grpc-contracts/gen/go/user_service/service/v1"
	"github.com/theartofdevel/grpc-service/internal/domain/entity"
)

type userServer struct {
	proto_user_service.UnimplementedUserServiceServer
}

func NewUserServer(unimplementedUserServiceServer proto_user_service.UnimplementedUserServiceServer) *userServer {
	return &userServer{UnimplementedUserServiceServer: unimplementedUserServiceServer}
}

func (s *userServer) GetUsers(
	ctx context.Context,
	req *proto_user_service.GetUsersRequest,
) (*proto_user_service.GetUsersResponse, error) {
	u := entity.User{
		ID:    "id",
		Name:  "name",
		Age:   123,
		Email: "user@mail.com",
	}

	return &proto_user_service.GetUsersResponse{
		Users: []*proto_user_model.User{
			u.ToProto(),
		},
	}, nil
}

func (s *userServer) UpdateUser(
	ctx context.Context,
	req *proto_user_service.UpdateUserRequest,
) (*proto_user_service.UpdateUserResponse, error) {
	return nil, nil
}
