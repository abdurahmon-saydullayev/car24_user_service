package storage

import (
	"Projects/Car24/car24_user_service/genproto/client_service"
	"Projects/Car24/car24_user_service/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Client() ClientRepoI
}

type ClientRepoI interface {
	Create(context.Context, *client_service.CreateClient) (*client_service.CLientPrimaryKey, error)
	GetByPK(ctx context.Context, req *client_service.CLientPrimaryKey) (resp *client_service.Client, err error)
	GetList(context.Context, *client_service.GetListClientRequest) (*client_service.GetListClientResponse, error)
	Update(ctx context.Context, req *client_service.UpdateClient) (rowsAffected int64, err error)
	UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *client_service.CLientPrimaryKey) error

	//otp
	CreateOTP(context.Context, *client_service.CreateOTP) error
	VerifyOTP(context.Context, *client_service.VerifyOTP) error

	GetByPhoneNumber(context.Context, *client_service.ClientPhoneNumberReq) (*client_service.Client, error)
}
