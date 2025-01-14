package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ UserServiceServer = (*UserService)(nil)

type UserService struct {
	UnimplementedUserServiceServer
	Name string
}

func NewUserService(name string) *UserService {
	return &UserService{
		Name: name,
	}
}

func (svc *UserService) GetByID(ctx context.Context, req *GetByIDReq) (*GetByIDResp, error) {
	uid, _ := ctx.Value("uid").(string)
	return &GetByIDResp{
		User: &User{
			Id:   req.Id,
			Name: "lazywoo",
		},
		Msg: fmt.Sprintf("uid: %v, from %v", uid, svc.Name),
	}, nil
}

type FailService struct {
	UnimplementedUserServiceServer
}

func NewFailService() *FailService {
	return &FailService{}
}

func (svc *FailService) GetByID(ctx context.Context, req *GetByIDReq) (*GetByIDResp, error) {
	return &GetByIDResp{}, status.Error(codes.Unavailable, "mock service fail")
}
