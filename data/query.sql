-- name: ListStocks :many
SELECT *
FROM stocks
WHERE
  (name LIKE sqlc.narg('name') OR sqlc.narg('name') IS NULL) AND
  (product_type = sqlc.narg('product_type') OR sqlc.narg('product_type') IS NULL) AND
  (product_attr = sqlc.narg('product_attr') OR sqlc.narg('product_attr') IS NULL) AND
  (status = sqlc.narg('status') OR sqlc.narg('status') IS NULL)
ORDER BY id DESC 
LIMIT sqlc.narg('limit')
OFFSET sqlc.narg('offset');

-- name: CreateStock :one
INSERT INTO stocks (status, name, product_type, product_attr,
    supplier, model, unit, price, batch_no_in, way_in, location,
    batch_no_produce, produce_date, disinfection_no, disinfection_date,
    stock_date, stock_num, current_num, value)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)
RETURNING *;

-- name: UpdateStocks :exec
UPDATE stocks
SET status = $1
WHERE batch_no_in = $2 and status = $3;

-- name: UpdateStockNumber :exec
UPDATE stocks
SET current_num = $1, value = $2
WHERE id = $3;

-- name: StocksByID :one
SELECT * FROM stocks
WHERE id = $1;

-- name: UpdateStockStatusByID :exec
UPDATE stocks
SET status = $1, price = $2, value = $3
WHERE id = $4;

-- name: UpdateStockPriceByID :exec
UPDATE stocks
SET price = $1, value = $2
WHERE id = $3;

-- name: CountStocks :one
SELECT count(*) FROM stocks
WHERE
  (name LIKE sqlc.narg('name') OR sqlc.narg('name') IS NULL) AND
  (product_type = sqlc.narg('product_type') OR sqlc.narg('product_type') IS NULL) AND
  (status = sqlc.narg('status') OR sqlc.narg('status') IS NULL);

-- name: ListApplications :many
SELECT *
FROM stock_applications
WHERE
  (application_date >= sqlc.narg('application_date_s') OR sqlc.narg('application_date_s') IS NULL) AND
  (application_date <= sqlc.narg('application_date_e') OR sqlc.narg('application_date_e') IS NULL) AND
  (application_user = sqlc.narg('application_user') OR sqlc.narg('application_user') IS NULL) AND
  (status = sqlc.narg('status') OR sqlc.narg('status') IS NULL)
ORDER BY id DESC 
LIMIT sqlc.narg('limit')
OFFSET sqlc.narg('offset');

-- name: ApplicationByID :one
SELECT * FROM stock_applications WHERE id = $1;

-- name: CountApplications :one
SELECT count(*)
FROM stock_applications
WHERE
  (application_date >= sqlc.narg('application_date_s') OR sqlc.narg('application_date_s') IS NULL) AND
  (application_date <= sqlc.narg('application_date_e') OR sqlc.narg('application_date_e') IS NULL) AND
  (application_user = sqlc.narg('application_user') OR sqlc.narg('application_user') IS NULL) AND
  (status = sqlc.narg('status') OR sqlc.narg('status') IS NULL);

-- name: CreateStockApplication :exec
INSERT INTO stock_applications(stock_id, application_date, batch_no_in, status,
    application_user, approve_date, approve_user, create_date)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8);

-- name: UpdateApplicationIN :exec
UPDATE stock_applications
SET status = $1, approve_user = $2
WHERE id = $3;

-- name: CreateOutApplication :exec
INSERT INTO stock_out_applications (
    stockids, number, status, application_user, approve_user, create_date)
VALUES ($1,$2,$3,$4,$5,$6);


-- name: ListOutApplications :many
SELECT *
FROM stock_out_applications
WHERE
  (application_user = sqlc.narg('application_user') OR sqlc.narg('application_user') IS NULL) AND
  (approve_user = sqlc.narg('approve_user') OR sqlc.narg('approve_user') IS NULL) AND
  (status = sqlc.narg('status') OR sqlc.narg('status') IS NULL)
ORDER BY id DESC 
LIMIT sqlc.narg('limit')
OFFSET sqlc.narg('offset');

-- name: CountOutApplication :one
SELECT COUNT(*)
FROM stock_out_applications
WHERE
  (application_user = sqlc.narg('application_user') OR sqlc.narg('application_user') IS NULL) AND
  (approve_user = sqlc.narg('approve_user') OR sqlc.narg('approve_user') IS NULL) AND
  (status = sqlc.narg('status') OR sqlc.narg('status') IS NULL);

-- name: UpdateApplicationOUT :exec
UPDATE stock_out_applications
SET status = $1, approve_user = $2
WHERE id = $3;

-- name: OutApplicationByID :one
SELECT * FROM stock_out_applications
WHERE id = $1;
