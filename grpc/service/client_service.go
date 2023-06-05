package service

import (
	"Projects/Car24/car24_user_service/config"
	"Projects/Car24/car24_user_service/genproto/client_service"
	"Projects/Car24/car24_user_service/grpc/client"
	"Projects/Car24/car24_user_service/models"
	"Projects/Car24/car24_user_service/pkg/logger"
	"Projects/Car24/car24_user_service/storage"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClientService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*client_service.UnimplementedClientServiceServer
}

func NewClientService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ClientService {
	return &ClientService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (u *ClientService) Create(ctx context.Context, req *client_service.CreateClient) (resp *client_service.Client, err error) {
	u.log.Info("---create user---", logger.Any("req", req))

	pkey, err := u.strg.Client().Create(ctx, req)
	if err != nil {
		return nil, err
	}

	resp, err = u.strg.Client().GetByPK(ctx, pkey)
	if err != nil {
		u.log.Error("!!!GetByPKeyUser->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (u *ClientService) GetByID(ctx context.Context, req *client_service.CLientPrimaryKey) (resp *client_service.Client, err error) {

	u.log.Info("---get userbyid---", logger.Any("req", req))

	client, err := u.strg.Client().GetByPK(ctx, req)
	if err != nil {
		return nil, err
	}

	return client, err
}

func (u *ClientService) GetList(ctx context.Context, req *client_service.GetListClientRequest) (resp *client_service.GetListClientResponse, err error) {

	u.log.Info("---GetUsers------>", logger.Any("req", req))

	resp, err = u.strg.Client().GetList(ctx, req)
	if err != nil {
		u.log.Error("!!!GetUsers->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (i *ClientService) Update(ctx context.Context, req *client_service.UpdateClient) (resp *client_service.Client, err error) {

	i.log.Info("---UpdateUser------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Client().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Client().GetByPK(ctx, &client_service.CLientPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetUser->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ClientService) UpdatePatch(ctx context.Context, req *client_service.UpdatePatchClient) (resp *client_service.Client, err error) {

	i.log.Info("---UpdatePatchUser------>", logger.Any("req", req))

	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Client().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Client().GetByPK(ctx, &client_service.CLientPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetUser->User->Get--->", logger.Error(err))

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ClientService) Delete(ctx context.Context, req *client_service.CLientPrimaryKey) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteUser------>", logger.Any("req", req))

	err = i.strg.Client().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteUser->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}

// otp
func (i *ClientService) CreateUserOTP(ctx context.Context, req *client_service.CreateOTP) (resp *empty.Empty, err error) {

	i.log.Info("---CreateUserOTP------->", logger.Any("req", req))

	err = i.strg.Client().CreateOTP(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateUserOTP->OTP->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}

func (i *ClientService) VerifyUserOTP(ctx context.Context, req *client_service.VerifyOTP) (resp *empty.Empty, err error) {

	i.log.Info("---VerifyUserOTP------->", logger.Any("req", req))

	err = i.strg.Client().VerifyOTP(ctx, req)
	if err != nil {
		i.log.Error("!!!VerifyUserOTP->OTP->Verify--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}

func (i *ClientService) Check(ctx context.Context, req *client_service.ClientPhoneNumberReq) (resp *client_service.Client, err error) {

	i.log.Info("---CheckUser------>", logger.Any("req", req))

	resp, err = i.strg.Client().GetByPhoneNumber(ctx, req)
	if err != nil {
		i.log.Error("!!!GetUserByPhoneNumber->User->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}
