package grpc_api

import (
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertUser(user db.Users) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
