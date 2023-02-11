-- name: ListStocks :many
SELECT *
FROM stocks
WHERE
  (name LIKE sqlc.narg('name') OR sqlc.narg('name') IS NULL) AND
  (product_type = sqlc.narg('product_type') OR sqlc.narg('product_type') IS NULL);

-- name: CountStocks :one
SELECT count(*) FROM stocks;

-- name: ListApplications :many
SELECT *
FROM stock_applications
WHERE
  (application_date >= sqlc.narg('application_date_s') OR sqlc.narg('application_date_s') IS NULL) AND
  (application_date <= sqlc.narg('application_date_e') OR sqlc.narg('application_date_e') IS NULL) AND
  (application_user = sqlc.narg('application_user') OR sqlc.narg('application_user') IS NULL) AND
  (status = sqlc.narg('status') OR sqlc.narg('status') IS NULL);
