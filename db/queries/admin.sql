-- name: GetUsers :many
SELECT "username",
    "name",
    "email",
    "address",
    "role",
    "credit_card",
    "enabled"
FROM "user"
ORDER BY "id" ASC
LIMIT $1 OFFSET $2;

-- name: EnabledShop :execrows
UPDATE "shop" AS s
SET s."enabled" = TRUE
WHERE s."seller_name" = $1;

-- name: DisableUser :execrows
WITH disabled_user AS (
    UPDATE "user"
    SET "enabled" = FALSE
    WHERE "username" = $1
    RETURNING "username"
),
disabled_shop AS (
    UPDATE "shop"
    SET "enabled" = FALSE
    WHERE "seller_name" =(
            SELECT "username"
            FROM disabled_user
        )
    RETURNING "id"
)
UPDATE "product"
SET "enabled" = FALSE
WHERE "shop_id" =(
        SELECT "id"
        FROM disabled_shop
    );

-- name: DisableShop :execrows
WITH disable_shop AS (
    UPDATE "shop" AS s
    SET s."enabled" = FALSE
    WHERE s."seller_name" = $1
    RETURNING s."id"
)
UPDATE "product" AS p
SET p."enabled" = FALSE
WHERE p."shop_id" =(
        SELECT "id"
        FROM disable_shop
    );

-- name: DisableProductsFromShop :execrows
UPDATE "product" AS p
SET p."enabled" = FALSE
WHERE p."shop_id" = $1;

-- name: CouponExists :one
SELECT EXISTS (
        SELECT 1
        FROM "coupon"
        WHERE "id" = $1
    );

-- name: GetGlobalCoupons :many
SELECT "id",
    "type",
    "scope",
    "name",
    "description",
    "discount",
    "start_date",
    "expire_date"
FROM "coupon"
WHERE "scope" = 'global'
ORDER BY "id" ASC
LIMIT $1 OFFSET $2;

-- name: GetGlobalCouponDetail :one
SELECT "id",
    "type",
    "scope",
    "name",
    "description",
    "discount",
    "start_date",
    "expire_date"
FROM "coupon"
WHERE "scope" = 'global'
    AND "id" = $1;

-- name: AddCoupon :one
INSERT INTO "coupon" (
        "type",
        "scope",
        "name",
        "description",
        "discount",
        "start_date",
        "expire_date"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING "id",
    "type",
    "scope",
    "name",
    "description",
    "discount",
    "start_date",
    "expire_date";

-- name: ValidateTags :one
SELECT NOT EXISTS (
        SELECT 1
        FROM "tag" AS T,
            "coupon" AS C
        WHERE T."id" != ANY(@tag_id::int [])
            AND C."id" = @coupon_id
            AND T."shop_id" = C."shop_id"
    );

-- name: AddCouponTags :execrows
INSERT INTO "coupon_tag"("coupon_id", "tag_id")
VALUES (@coupon_id, @tag_id::int []) ON CONFLICT ("coupon_id", "tag_id") DO NOTHING;

-- name: GetCouponTags :many
SELECT "tag_id",
    "name"
FROM "coupon_tag" AS CT,
    "tag" AS T
WHERE CT."coupon_id" = $1
    AND CT."tag_id" = T."id";

-- name: EditCoupon :one
UPDATE "coupon"
SET "type" = COALESCE($2, "type"),
    "name" = COALESCE($3, "name"),
    "description" = COALESCE($4, "description"),
    "discount" = COALESCE($5, "discount"),
    "start_date" = COALESCE($6, "start_date"),
    "expire_date" = COALESCE($7, "expire_date")
WHERE "id" = $1
    AND "scope" = 'global'
RETURNING "id",
    "type",
    "scope",
    "name",
    "description",
    "discount",
    "start_date",
    "expire_date";

-- name: DeleteCoupon :execrows
DELETE FROM "coupon"
WHERE "id" = $1
    AND "scope" = 'global';

-- name: GetUserIDByUsername :one
SELECT "id"
FROM "user"
WHERE "username" = $1;

-- name: GetShopIDBySellerName :one
SELECT "id"
FROM "shop"
WHERE "seller_name" = $1;

-- name: GetTopSeller :many
SELECT S."seller_name",
    S."name",
    S."image_id",
    SUM(O."total_price") AS "total_sales"
FROM "shop" AS S,
    "order_history" AS O
WHERE S."id" = O."shop_id"
    AND O."status" = 'paid'
    AND O."created_at" < (@date) + INTERVAL '1 month'
    AND O."created_at" >= @date
GROUP BY S."seller_name",
    S."name",
    S."image_id"
ORDER BY "total_sales" DESC
LIMIT 3;
