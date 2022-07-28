package entity

import proto_user_model "github.com/theartofdevel/grpc-contracts/gen/go/user_service/model/v1"

type User struct {
	ID    string
	Name  string
	Age   uint32
	Email string
}

func (u *User) ToProto() *proto_user_model.User {
	return &proto_user_model.User{
		Id:    u.ID,
		Name:  u.Name,
		Age:   u.Age,
		Email: u.Email,
	}
}
