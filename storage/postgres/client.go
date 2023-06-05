package postgres

import (
	"Projects/Car24/car24_user_service/genproto/client_service"
	"Projects/Car24/car24_user_service/models"
	"Projects/Car24/car24_user_service/pkg/helper"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewClientRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(ctx context.Context, req *client_service.CreateClient) (resp *client_service.CLientPrimaryKey, err error) {
	id := uuid.New().String()

	query := `
	INSERT INTO "client" (
		id,
		first_name,
		last_name,
		address,
		phone_number,
		driving_license_number,
		passport_number,
		photo,
		driving_number_given_place,
		driving_number_given_date,
		driving_number_expired,
		propiska,
		passport_pinfl,
		additional_phone_number,
		created_at,
		updated_at
	) VALUES($1, $2, $3, $4, $5, $6,$7,$8,$9,$10,$11,$12,$13,$14, NOW(), NOW())
	`

	_, err = u.db.Exec(
		ctx,
		query,
		id,
		req.FirstName,
		req.LastName,
		req.Address,
		req.PhoneNumber,
		req.DrivingLicenseNumber,
		req.PassportNumber,
		req.Photo,
		req.DrivingNumberGivenPlace,
		req.DrivingNumberGivenDate,
		req.DrivingNumberExpired,
		req.Propiska,
		req.PassportPinfl,
		req.AdditionalPhoneNumber,
	)
	if err != nil {
		return nil, err
	}

	return &client_service.CLientPrimaryKey{Id: id}, nil
}

func (u *userRepo) GetByPK(ctx context.Context, req *client_service.CLientPrimaryKey) (resp *client_service.Client, err error) {

	query := `
		SELECT 
			id,
			first_name,
			last_name,
			address,
			phone_number,
			driving_license_number,
			passport_number,
			photo,
			driving_number_given_place,
			driving_number_given_date,
			driving_number_expired,
			propiska,
			passport_pinfl,
			additional_phone_number,
			created_at,
			updated_at
		FROM "client"
		WHERE id = $1
	`
	resp = &client_service.Client{}
	var (
		id                         sql.NullString
		first_name                 sql.NullString
		last_name                  sql.NullString
		address                    sql.NullString
		phone_number               sql.NullString
		driving_license_number     sql.NullString
		passport_number            sql.NullString
		photo                      sql.NullString
		driving_number_given_place sql.NullString
		driving_number_given_date  sql.NullString
		driving_number_expired     sql.NullString
		propiska                   sql.NullString
		passport_pinfl             sql.NullString
		additional_phone_number    sql.NullString
		created_at                 sql.NullString
		updated_at                 sql.NullString
	)

	err = u.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&first_name,
		&last_name,
		&address,
		&phone_number,
		&driving_license_number,
		&passport_number,
		&photo,
		&driving_number_given_place,
		&driving_number_given_date,
		&driving_number_expired,
		&propiska,
		&passport_pinfl,
		&additional_phone_number,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return resp, err
	}

	resp = &client_service.Client{
		Id:                      id.String,
		FirstName:               first_name.String,
		LastName:                last_name.String,
		Address:                 address.String,
		PhoneNumber:             phone_number.String,
		DrivingLicenseNumber:    driving_license_number.String,
		PassportNumber:          passport_number.String,
		Photo:                   photo.String,
		DrivingNumberGivenPlace: driving_number_given_place.String,
		DrivingNumberGivenDate:  driving_number_given_date.String,
		DrivingNumberExpired:    driving_number_expired.String,
		Propiska:                propiska.String,
		CreatedAt:               created_at.String,
		UpdatedAt:               updated_at.String,
	}

	return resp, nil
}

func (u *userRepo) GetList(ctx context.Context, req *client_service.GetListClientRequest) (resp *client_service.GetListClientResponse, err error) {
	resp = &client_service.GetListClientResponse{}

	var (
		query  string
		limit  = ""
		offset = " OFFSET 0 "
		params = make(map[string]interface{})
		filter = " WHERE TRUE "
		sort   = " ORDER BY created_at DESC"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			first_name,
			last_name,
			address,
			phone_number,
			driving_license_number,
			passport_number,
			photo,
			driving_number_given_place,
			driving_number_given_date,
			driving_number_expired,
			propiska,
			passport_pinfl,
			additional_phone_number,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "user"
	`
	if len(req.GetSearch()) > 0 {
		filter += " AND (first_name || ' ' || last_name) ILIKE '%' || '" + req.Search + "' || '%' "
	}
	if req.GetLimit() > 0 {
		limit = " LIMIT :limit"
		params["limit"] = req.Limit
	}
	if req.GetOffset() > 0 {
		offset = " OFFSET :offset"
		params["offset"] = req.Offset
	}
	query += filter + sort + offset + limit

	query, args := helper.ReplaceQueryParams(query, params)
	rows, err := u.db.Query(ctx, query, args...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id                         sql.NullString
			first_name                 sql.NullString
			last_name                  sql.NullString
			address                    sql.NullString
			phone_number               sql.NullString
			driving_license_number     sql.NullString
			passport_number            sql.NullString
			photo                      sql.NullString
			driving_number_given_place sql.NullString
			driving_number_given_date  sql.NullString
			driving_number_expired     sql.NullString
			propiska                   sql.NullString
			passport_pinfl             sql.NullString
			additional_phone_number    sql.NullString
			created_at                 sql.NullString
			updated_at                 sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&first_name,
			&last_name,
			&address,
			&phone_number,
			&driving_license_number,
			&passport_number,
			&photo,
			&driving_number_given_place,
			&driving_number_given_date,
			&driving_number_expired,
			&propiska,
			&passport_pinfl,
			&additional_phone_number,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return resp, err
		}

		resp.Clients = append(resp.Clients, &client_service.Client{
			Id:                      id.String,
			FirstName:               first_name.String,
			LastName:                last_name.String,
			Address:                 address.String,
			PhoneNumber:             phone_number.String,
			DrivingLicenseNumber:    driving_license_number.String,
			PassportNumber:          passport_number.String,
			Photo:                   photo.String,
			DrivingNumberGivenPlace: driving_number_given_place.String,
			DrivingNumberGivenDate:  driving_number_given_date.String,
			DrivingNumberExpired:    driving_number_expired.String,
			Propiska:                propiska.String,
			CreatedAt:               created_at.String,
			UpdatedAt:               updated_at.String,
		})
	}

	return resp, nil
}

func (u *userRepo) Update(ctx context.Context, req *client_service.UpdateClient) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			"client"
		SET
			first_name = :first_name,
			last_name = :last_name,
			address = :address,
			phone_number= :phone_number,
			driving_license_number = :driving_license_number,
			passport_number, = : passport_number,
			photo = :photo,
			driving_number_given_place =:driving_number_given_place
			driving_number_given_date = :driving_number_given_date,
			driving_number_expired = :driving_number_expired,
			propiska = :propiska,
			passport_pinfl = :passport_pinfl
			additional_phone_number = :additional_phone_number
		WHERE id = :id
	`
	params = map[string]interface{}{
		"first_name":                 req.GetFirstName(),
		"last_name":                  req.GetLastName(),
		"address":                    req.GetAddress(),
		"phone_number":               req.GetPhoneNumber(),
		"driving_license_number":     req.GetDrivingLicenseNumber(),
		"passport_number":            req.GetPassportNumber(),
		"photo":                      req.GetPhoto(),
		"driving_number_given_place": req.GetDrivingNumberGivenPlace(),
		"driving_number_given_date":  req.GetDrivingNumberGivenDate(),
		"driving_number_expired":     req.GetDrivingNumberExpired(),
		"propiska":                   req.GetPropiska(),
		"passport_pinfl":             req.GetPassportPinfl(),
		"additional_phone_number":    req.GetAdditionalPhoneNumber(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := u.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (u *userRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {

	var (
		set   = " SET "
		ind   = 0
		query string
	)

	if len(req.Fields) == 0 {
		err = errors.New("no updates provided")
		return
	}

	req.Fields["id"] = req.Id

	for key := range req.Fields {
		set += fmt.Sprintf(" %s = :%s ", key, key)
		if ind != len(req.Fields)-1 {
			set += ", "
		}
		ind++
	}

	query = `
		UPDATE
			"client"
	` + set + ` , updated_at = now()
		WHERE
			id = :id
	`

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	result, err := u.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), err
}

func (u *userRepo) Delete(ctx context.Context, req *client_service.CLientPrimaryKey) error {

	query := `DELETE FROM "user" WHERE id = $1`

	_, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil
}
