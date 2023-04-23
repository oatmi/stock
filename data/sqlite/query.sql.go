// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package sqlite

import (
	"context"
	"database/sql"
)

const countStocks = `-- name: CountStocks :one
SELECT count(*) FROM stocks
`

func (q *Queries) CountStocks(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countStocks)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createOutApplication = `-- name: CreateOutApplication :exec
INSERT INTO stock_out_applications (
    stockids, number, status, application_user, approve_user, create_date)
VALUES ($1,$2,$3,$4,$5,$6)
`

type CreateOutApplicationParams struct {
	Stockids        string `json:"stockids"`
	Number          int32  `json:"number"`
	Status          int32  `json:"status"`
	ApplicationUser string `json:"application_user"`
	ApproveUser     string `json:"approve_user"`
	CreateDate      string `json:"create_date"`
}

func (q *Queries) CreateOutApplication(ctx context.Context, arg CreateOutApplicationParams) error {
	_, err := q.db.ExecContext(ctx, createOutApplication,
		arg.Stockids,
		arg.Number,
		arg.Status,
		arg.ApplicationUser,
		arg.ApproveUser,
		arg.CreateDate,
	)
	return err
}

const createStock = `-- name: CreateStock :exec
INSERT INTO stocks (status, name, product_type, product_attr,
    supplier, model, unit, price, batch_no_in, way_in, location,
    batch_no_produce, produce_date, disinfection_no, disinfection_date,
    stock_date, stock_num, current_num, value)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)
`

type CreateStockParams struct {
	Status           int32  `json:"status"`
	Name             string `json:"name"`
	ProductType      int32  `json:"product_type"`
	ProductAttr      int32  `json:"product_attr"`
	Supplier         string `json:"supplier"`
	Model            string `json:"model"`
	Unit             string `json:"unit"`
	Price            int32  `json:"price"`
	BatchNoIn        string `json:"batch_no_in"`
	WayIn            int32  `json:"way_in"`
	Location         int32  `json:"location"`
	BatchNoProduce   string `json:"batch_no_produce"`
	ProduceDate      int32  `json:"produce_date"`
	DisinfectionNo   string `json:"disinfection_no"`
	DisinfectionDate int32  `json:"disinfection_date"`
	StockDate        int32  `json:"stock_date"`
	StockNum         int32  `json:"stock_num"`
	CurrentNum       int32  `json:"current_num"`
	Value            int32  `json:"value"`
}

func (q *Queries) CreateStock(ctx context.Context, arg CreateStockParams) error {
	_, err := q.db.ExecContext(ctx, createStock,
		arg.Status,
		arg.Name,
		arg.ProductType,
		arg.ProductAttr,
		arg.Supplier,
		arg.Model,
		arg.Unit,
		arg.Price,
		arg.BatchNoIn,
		arg.WayIn,
		arg.Location,
		arg.BatchNoProduce,
		arg.ProduceDate,
		arg.DisinfectionNo,
		arg.DisinfectionDate,
		arg.StockDate,
		arg.StockNum,
		arg.CurrentNum,
		arg.Value,
	)
	return err
}

const createStockApplication = `-- name: CreateStockApplication :exec
INSERT INTO stock_applications(application_date, batch_no_in, status,
    application_user, approve_date, approve_user, create_date)
VALUES ($1,$2,$3,$4,$5,$6,$7)
`

type CreateStockApplicationParams struct {
	ApplicationDate string `json:"application_date"`
	BatchNoIn       string `json:"batch_no_in"`
	Status          int32  `json:"status"`
	ApplicationUser string `json:"application_user"`
	ApproveDate     string `json:"approve_date"`
	ApproveUser     string `json:"approve_user"`
	CreateDate      string `json:"create_date"`
}

func (q *Queries) CreateStockApplication(ctx context.Context, arg CreateStockApplicationParams) error {
	_, err := q.db.ExecContext(ctx, createStockApplication,
		arg.ApplicationDate,
		arg.BatchNoIn,
		arg.Status,
		arg.ApplicationUser,
		arg.ApproveDate,
		arg.ApproveUser,
		arg.CreateDate,
	)
	return err
}

const listApplications = `-- name: ListApplications :many
SELECT id, application_date, batch_no_in, status, application_user, approve_user, approve_date, create_date
FROM stock_applications
WHERE
  (application_date >= $1 OR $1 IS NULL) AND
  (application_date <= $2 OR $2 IS NULL) AND
  (application_user = $3 OR $3 IS NULL) AND
  (status = $4 OR $4 IS NULL)
`

type ListApplicationsParams struct {
	ApplicationDateS sql.NullString `json:"application_date_s"`
	ApplicationDateE sql.NullString `json:"application_date_e"`
	ApplicationUser  sql.NullString `json:"application_user"`
	Status           sql.NullInt32  `json:"status"`
}

func (q *Queries) ListApplications(ctx context.Context, arg ListApplicationsParams) ([]StockApplication, error) {
	rows, err := q.db.QueryContext(ctx, listApplications,
		arg.ApplicationDateS,
		arg.ApplicationDateE,
		arg.ApplicationUser,
		arg.Status,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StockApplication
	for rows.Next() {
		var i StockApplication
		if err := rows.Scan(
			&i.ID,
			&i.ApplicationDate,
			&i.BatchNoIn,
			&i.Status,
			&i.ApplicationUser,
			&i.ApproveUser,
			&i.ApproveDate,
			&i.CreateDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listOutApplications = `-- name: ListOutApplications :many
SELECT id, stockids, number, status, application_user, approve_user, create_date
FROM stock_out_applications
WHERE
  (application_user >= $1 OR $1 IS NULL) AND
  (approve_user <= $2 OR $2 IS NULL) AND
  (status = $3 OR $3 IS NULL)
`

type ListOutApplicationsParams struct {
	ApplicationUser sql.NullString `json:"application_user"`
	ApproveUser     sql.NullString `json:"approve_user"`
	Status          sql.NullInt32  `json:"status"`
}

func (q *Queries) ListOutApplications(ctx context.Context, arg ListOutApplicationsParams) ([]StockOutApplication, error) {
	rows, err := q.db.QueryContext(ctx, listOutApplications, arg.ApplicationUser, arg.ApproveUser, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StockOutApplication
	for rows.Next() {
		var i StockOutApplication
		if err := rows.Scan(
			&i.ID,
			&i.Stockids,
			&i.Number,
			&i.Status,
			&i.ApplicationUser,
			&i.ApproveUser,
			&i.CreateDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listStocks = `-- name: ListStocks :many
SELECT id, status, name, product_type, product_attr, supplier, model, unit, price, batch_no_in, way_in, location, batch_no_produce, produce_date, disinfection_no, disinfection_date, stock_date, stock_num, current_num, value
FROM stocks
WHERE
  (name LIKE $1 OR $1 IS NULL) AND
  (product_type = $2 OR $2 IS NULL) AND
  (status = $3 OR $3 IS NULL)
`

type ListStocksParams struct {
	Name        sql.NullString `json:"name"`
	ProductType sql.NullInt32  `json:"product_type"`
	Status      sql.NullInt32  `json:"status"`
}

func (q *Queries) ListStocks(ctx context.Context, arg ListStocksParams) ([]Stock, error) {
	rows, err := q.db.QueryContext(ctx, listStocks, arg.Name, arg.ProductType, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Stock
	for rows.Next() {
		var i Stock
		if err := rows.Scan(
			&i.ID,
			&i.Status,
			&i.Name,
			&i.ProductType,
			&i.ProductAttr,
			&i.Supplier,
			&i.Model,
			&i.Unit,
			&i.Price,
			&i.BatchNoIn,
			&i.WayIn,
			&i.Location,
			&i.BatchNoProduce,
			&i.ProduceDate,
			&i.DisinfectionNo,
			&i.DisinfectionDate,
			&i.StockDate,
			&i.StockNum,
			&i.CurrentNum,
			&i.Value,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const outApplicationByID = `-- name: OutApplicationByID :one
SELECT id, stockids, number, status, application_user, approve_user, create_date FROM stock_out_applications
WHERE id = $1
`

func (q *Queries) OutApplicationByID(ctx context.Context, id int32) (StockOutApplication, error) {
	row := q.db.QueryRowContext(ctx, outApplicationByID, id)
	var i StockOutApplication
	err := row.Scan(
		&i.ID,
		&i.Stockids,
		&i.Number,
		&i.Status,
		&i.ApplicationUser,
		&i.ApproveUser,
		&i.CreateDate,
	)
	return i, err
}

const stocksByID = `-- name: StocksByID :one
SELECT id, status, name, product_type, product_attr, supplier, model, unit, price, batch_no_in, way_in, location, batch_no_produce, produce_date, disinfection_no, disinfection_date, stock_date, stock_num, current_num, value FROM stocks
WHERE id = $1
`

func (q *Queries) StocksByID(ctx context.Context, id int32) (Stock, error) {
	row := q.db.QueryRowContext(ctx, stocksByID, id)
	var i Stock
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.Name,
		&i.ProductType,
		&i.ProductAttr,
		&i.Supplier,
		&i.Model,
		&i.Unit,
		&i.Price,
		&i.BatchNoIn,
		&i.WayIn,
		&i.Location,
		&i.BatchNoProduce,
		&i.ProduceDate,
		&i.DisinfectionNo,
		&i.DisinfectionDate,
		&i.StockDate,
		&i.StockNum,
		&i.CurrentNum,
		&i.Value,
	)
	return i, err
}

const updateApplicationIN = `-- name: UpdateApplicationIN :exec
UPDATE stock_applications
SET status = $1
WHERE id = $2
`

type UpdateApplicationINParams struct {
	Status int32 `json:"status"`
	ID     int32 `json:"id"`
}

func (q *Queries) UpdateApplicationIN(ctx context.Context, arg UpdateApplicationINParams) error {
	_, err := q.db.ExecContext(ctx, updateApplicationIN, arg.Status, arg.ID)
	return err
}

const updateApplicationOUT = `-- name: UpdateApplicationOUT :exec
UPDATE stock_out_applications
SET status = $1
WHERE id = $2
`

type UpdateApplicationOUTParams struct {
	Status int32 `json:"status"`
	ID     int32 `json:"id"`
}

func (q *Queries) UpdateApplicationOUT(ctx context.Context, arg UpdateApplicationOUTParams) error {
	_, err := q.db.ExecContext(ctx, updateApplicationOUT, arg.Status, arg.ID)
	return err
}

const updateStockNumber = `-- name: UpdateStockNumber :exec
UPDATE stocks
SET current_num = $1
WHERE id = $2
`

type UpdateStockNumberParams struct {
	CurrentNum int32 `json:"current_num"`
	ID         int32 `json:"id"`
}

func (q *Queries) UpdateStockNumber(ctx context.Context, arg UpdateStockNumberParams) error {
	_, err := q.db.ExecContext(ctx, updateStockNumber, arg.CurrentNum, arg.ID)
	return err
}

const updateStockPriceByID = `-- name: UpdateStockPriceByID :exec
UPDATE stocks
SET price = $1, value = $2
WHERE id = $3
`

type UpdateStockPriceByIDParams struct {
	Price int32 `json:"price"`
	Value int32 `json:"value"`
	ID    int32 `json:"id"`
}

func (q *Queries) UpdateStockPriceByID(ctx context.Context, arg UpdateStockPriceByIDParams) error {
	_, err := q.db.ExecContext(ctx, updateStockPriceByID, arg.Price, arg.Value, arg.ID)
	return err
}

const updateStockStatusByID = `-- name: UpdateStockStatusByID :exec
UPDATE stocks
SET status = $1
WHERE id = $2
`

type UpdateStockStatusByIDParams struct {
	Status int32 `json:"status"`
	ID     int32 `json:"id"`
}

func (q *Queries) UpdateStockStatusByID(ctx context.Context, arg UpdateStockStatusByIDParams) error {
	_, err := q.db.ExecContext(ctx, updateStockStatusByID, arg.Status, arg.ID)
	return err
}

const updateStocks = `-- name: UpdateStocks :exec
UPDATE stocks
SET status = $1
WHERE batch_no_in = $2 and status = $3
`

type UpdateStocksParams struct {
	Status    int32  `json:"status"`
	BatchNoIn string `json:"batch_no_in"`
	Status_2  int32  `json:"status_2"`
}

func (q *Queries) UpdateStocks(ctx context.Context, arg UpdateStocksParams) error {
	_, err := q.db.ExecContext(ctx, updateStocks, arg.Status, arg.BatchNoIn, arg.Status_2)
	return err
}
