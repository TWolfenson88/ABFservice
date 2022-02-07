package api

import (
	"context"
	"errors"
	"github.com/TWolfenson88/ABFservice/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type abfServiceServer struct {
	UnimplementedABFServiceServer
	abfService *core.ABFConfig
}


func NewABFServer(abfs *core.ABFConfig) *abfServiceServer {
	return &abfServiceServer{abfService: abfs}
}

func (a *abfServiceServer) AuthRequest(ctx context.Context, req *UserData) (*AllowUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthRequest not implemented")
}

func (a *abfServiceServer) ResetBucket(ctx context.Context, req *UserData) (*StatusMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetBucket not implemented")
}

func (a *abfServiceServer) AddToList(ctx context.Context, req *IpData) (*StatusMsg, error) {
	_, _, err := net.ParseCIDR(req.GetIpAddr())
	if err != nil {
		return &StatusMsg{
			Error:         status.Error(codes.Internal, err.Error()).Error(),
		}, err
	}

	err = a.abfService.NetLists.AddSubnetToList(req.GetIpAddr(), req.GetList())

	if err != nil {
		return &StatusMsg{
			Error: status.Error(codes.Internal, err.Error()).Error(),
		}, err
	}

	return &StatusMsg{
		Error: errors.New("no error").Error(),
	}, nil
}

func (a *abfServiceServer) RemoveFromList(ctx context.Context, req *IpData) (*StatusMsg, error) {
	_, _, err := net.ParseCIDR(req.GetIpAddr())
	if err != nil {
		return &StatusMsg{
			Error:         status.Error(codes.Internal, err.Error()).Error(),
		}, err
	}

	err = a.abfService.NetLists.RemoveSubnetFromList(req.GetIpAddr(), req.GetList())

	if err != nil {
		return &StatusMsg{
			Error: status.Error(codes.Internal, err.Error()).Error(),
		}, err
	}

	return &StatusMsg{
		Error: errors.New("no error").Error(),
	}, nil
}