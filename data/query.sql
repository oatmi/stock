-- name: ListStocks :many
SELECT * FROM stocks
WHERE
    ($1 IS NULL OR name LIKE $1)
ORDER BY id;

-- name: CountStocks :one
SELECT count(*) FROM stocks;
