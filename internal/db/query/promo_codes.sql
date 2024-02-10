-- name: CreatePromoCode :one
INSERT INTO promo_codes (
    promo_id, promo_name, promo, active
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: UpdatePromoCode :exec
UPDATE promo_codes
SET active = $2
WHERE promo_name = $1;

-- name: DeletePromoCode :exec
DELETE FROM promo_codes
WHERE promo_name = $1;

-- name: QuerryPromoCodes :many
SELECT *
FROM promo_codes
WHERE active = $1;