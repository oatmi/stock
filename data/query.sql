-- name: ListStocks :many
SELECT *
FROM stocks
WHERE
  (name LIKE sqlc.narg('name') OR sqlc.narg('name') IS NULL) AND
  (product_type = sqlc.narg('product_type') OR sqlc.narg('product_type') IS NULL);

-- name: CountStocks :one
SELECT count(*) FROM stocks;
