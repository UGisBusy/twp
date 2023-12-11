// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: admin.sql

package db

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgtype"
)

const addCoupon = `-- name: AddCoupon :one
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
    "expire_date"
`

type AddCouponParams struct {
	Type        CouponType         `json:"type"`
	Scope       CouponScope        `json:"scope"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Discount    pgtype.Numeric     `json:"discount" swaggertype:"number"`
	StartDate   pgtype.Timestamptz `json:"start_date" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

type AddCouponRow struct {
	ID          int32              `json:"id" param:"id"`
	Type        CouponType         `json:"type"`
	Scope       CouponScope        `json:"scope"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Discount    pgtype.Numeric     `json:"discount" swaggertype:"number"`
	StartDate   pgtype.Timestamptz `json:"start_date" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

func (q *Queries) AddCoupon(ctx context.Context, arg AddCouponParams) (AddCouponRow, error) {
	row := q.db.QueryRow(ctx, addCoupon,
		arg.Type,
		arg.Scope,
		arg.Name,
		arg.Description,
		arg.Discount,
		arg.StartDate,
		arg.ExpireDate,
	)
	var i AddCouponRow
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Scope,
		&i.Name,
		&i.Description,
		&i.Discount,
		&i.StartDate,
		&i.ExpireDate,
	)
	return i, err
}

const addCouponTags = `-- name: AddCouponTags :execrows
INSERT INTO "coupon_tag"("coupon_id", "tag_id")
VALUES ($1, $2::int []) ON CONFLICT ("coupon_id", "tag_id") DO NOTHING
`

type AddCouponTagsParams struct {
	CouponID int32   `json:"coupon_id" param:"id"`
	TagID    []int32 `json:"tag_id"`
}

func (q *Queries) AddCouponTags(ctx context.Context, arg AddCouponTagsParams) (int64, error) {
	result, err := q.db.Exec(ctx, addCouponTags, arg.CouponID, arg.TagID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const couponExists = `-- name: CouponExists :one
SELECT EXISTS (
        SELECT 1
        FROM "coupon"
        WHERE "id" = $1
    )
`

func (q *Queries) CouponExists(ctx context.Context, id int32) (bool, error) {
	row := q.db.QueryRow(ctx, couponExists, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const deleteCoupon = `-- name: DeleteCoupon :execrows
DELETE FROM "coupon"
WHERE "id" = $1
    AND "scope" = 'global'
`

func (q *Queries) DeleteCoupon(ctx context.Context, id int32) (int64, error) {
	result, err := q.db.Exec(ctx, deleteCoupon, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const disableProductsFromShop = `-- name: DisableProductsFromShop :execrows
UPDATE "product" AS p
SET p."enabled" = FALSE
WHERE p."shop_id" = $1
`

func (q *Queries) DisableProductsFromShop(ctx context.Context, shopID int32) (int64, error) {
	result, err := q.db.Exec(ctx, disableProductsFromShop, shopID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const disableShop = `-- name: DisableShop :execrows
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
    )
`

func (q *Queries) DisableShop(ctx context.Context, sellerName string) (int64, error) {
	result, err := q.db.Exec(ctx, disableShop, sellerName)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const disableUser = `-- name: DisableUser :execrows
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
    )
`

func (q *Queries) DisableUser(ctx context.Context, username string) (int64, error) {
	result, err := q.db.Exec(ctx, disableUser, username)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const editCoupon = `-- name: EditCoupon :one
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
    "expire_date"
`

type EditCouponParams struct {
	ID          int32              `json:"id" param:"id"`
	Type        CouponType         `json:"type"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Discount    pgtype.Numeric     `json:"discount" swaggertype:"number"`
	StartDate   pgtype.Timestamptz `json:"start_date" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

type EditCouponRow struct {
	ID          int32              `json:"id" param:"id"`
	Type        CouponType         `json:"type"`
	Scope       CouponScope        `json:"scope"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Discount    pgtype.Numeric     `json:"discount" swaggertype:"number"`
	StartDate   pgtype.Timestamptz `json:"start_date" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

func (q *Queries) EditCoupon(ctx context.Context, arg EditCouponParams) (EditCouponRow, error) {
	row := q.db.QueryRow(ctx, editCoupon,
		arg.ID,
		arg.Type,
		arg.Name,
		arg.Description,
		arg.Discount,
		arg.StartDate,
		arg.ExpireDate,
	)
	var i EditCouponRow
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Scope,
		&i.Name,
		&i.Description,
		&i.Discount,
		&i.StartDate,
		&i.ExpireDate,
	)
	return i, err
}

const enabledShop = `-- name: EnabledShop :execrows
UPDATE "shop" AS s
SET s."enabled" = TRUE
WHERE s."seller_name" = $1
`

func (q *Queries) EnabledShop(ctx context.Context, sellerName string) (int64, error) {
	result, err := q.db.Exec(ctx, enabledShop, sellerName)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const getCouponTags = `-- name: GetCouponTags :many
SELECT "tag_id",
    "name"
FROM "coupon_tag" AS CT,
    "tag" AS T
WHERE CT."coupon_id" = $1
    AND CT."tag_id" = T."id"
`

type GetCouponTagsRow struct {
	TagID int32  `json:"tag_id"`
	Name  string `json:"name"`
}

func (q *Queries) GetCouponTags(ctx context.Context, couponID int32) ([]GetCouponTagsRow, error) {
	rows, err := q.db.Query(ctx, getCouponTags, couponID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetCouponTagsRow{}
	for rows.Next() {
		var i GetCouponTagsRow
		if err := rows.Scan(&i.TagID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGlobalCouponDetail = `-- name: GetGlobalCouponDetail :one
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
    AND "id" = $1
`

type GetGlobalCouponDetailRow struct {
	ID          int32              `json:"id" param:"id"`
	Type        CouponType         `json:"type"`
	Scope       CouponScope        `json:"scope"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Discount    pgtype.Numeric     `json:"discount" swaggertype:"number"`
	StartDate   pgtype.Timestamptz `json:"start_date" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

func (q *Queries) GetGlobalCouponDetail(ctx context.Context, id int32) (GetGlobalCouponDetailRow, error) {
	row := q.db.QueryRow(ctx, getGlobalCouponDetail, id)
	var i GetGlobalCouponDetailRow
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Scope,
		&i.Name,
		&i.Description,
		&i.Discount,
		&i.StartDate,
		&i.ExpireDate,
	)
	return i, err
}

const getGlobalCoupons = `-- name: GetGlobalCoupons :many
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
LIMIT $1 OFFSET $2
`

type GetGlobalCouponsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetGlobalCouponsRow struct {
	ID          int32              `json:"id" param:"id"`
	Type        CouponType         `json:"type"`
	Scope       CouponScope        `json:"scope"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Discount    pgtype.Numeric     `json:"discount" swaggertype:"number"`
	StartDate   pgtype.Timestamptz `json:"start_date" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

func (q *Queries) GetGlobalCoupons(ctx context.Context, arg GetGlobalCouponsParams) ([]GetGlobalCouponsRow, error) {
	rows, err := q.db.Query(ctx, getGlobalCoupons, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetGlobalCouponsRow{}
	for rows.Next() {
		var i GetGlobalCouponsRow
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Scope,
			&i.Name,
			&i.Description,
			&i.Discount,
			&i.StartDate,
			&i.ExpireDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getShopIDBySellerName = `-- name: GetShopIDBySellerName :one
SELECT "id"
FROM "shop"
WHERE "seller_name" = $1
`

func (q *Queries) GetShopIDBySellerName(ctx context.Context, sellerName string) (int32, error) {
	row := q.db.QueryRow(ctx, getShopIDBySellerName, sellerName)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getTopSeller = `-- name: GetTopSeller :many
SELECT S."seller_name",
    S."name",
    S."image_id",
    SUM(O."total_price") AS "total_sales"
FROM "shop" AS S,
    "order_history" AS O
WHERE S."id" = O."shop_id"
    AND O."status" = 'paid'
    AND O."created_at" < ($1) + INTERVAL '1 month'
    AND O."created_at" >= $1
GROUP BY S."seller_name",
    S."name",
    S."image_id"
ORDER BY "total_sales" DESC
LIMIT 3
`

type GetTopSellerRow struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	Name       string `json:"name"`
	ImageID    string `json:"image_id" swaggertype:"string"`
	TotalSales int64  `json:"total_sales"`
}

func (q *Queries) GetTopSeller(ctx context.Context, date interface{}) ([]GetTopSellerRow, error) {
	rows, err := q.db.Query(ctx, getTopSeller, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetTopSellerRow{}
	for rows.Next() {
		var i GetTopSellerRow
		if err := rows.Scan(
			&i.SellerName,
			&i.Name,
			&i.ImageID,
			&i.TotalSales,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserIDByUsername = `-- name: GetUserIDByUsername :one
SELECT "id"
FROM "user"
WHERE "username" = $1
`

func (q *Queries) GetUserIDByUsername(ctx context.Context, username string) (int32, error) {
	row := q.db.QueryRow(ctx, getUserIDByUsername, username)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getUsers = `-- name: GetUsers :many
SELECT "username",
    "name",
    "email",
    "address",
    "role",
    "credit_card",
    "enabled"
FROM "user"
ORDER BY "id" ASC
LIMIT $1 OFFSET $2
`

type GetUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetUsersRow struct {
	Username   string          `json:"username"`
	Name       string          `json:"name"`
	Email      string          `json:"email"`
	Address    string          `json:"address"`
	Role       RoleType        `json:"role"`
	CreditCard json.RawMessage `json:"credit_card"`
	Enabled    bool            `json:"enabled"`
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]GetUsersRow, error) {
	rows, err := q.db.Query(ctx, getUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUsersRow{}
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.Username,
			&i.Name,
			&i.Email,
			&i.Address,
			&i.Role,
			&i.CreditCard,
			&i.Enabled,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const validateTags = `-- name: ValidateTags :one
SELECT NOT EXISTS (
        SELECT 1
        FROM "tag" AS T,
            "coupon" AS C
        WHERE T."id" != ANY($1::int [])
            AND C."id" = $2
            AND T."shop_id" = C."shop_id"
    )
`

type ValidateTagsParams struct {
	TagID    []int32 `json:"tag_id"`
	CouponID int32   `json:"coupon_id" param:"id"`
}

func (q *Queries) ValidateTags(ctx context.Context, arg ValidateTagsParams) (bool, error) {
	row := q.db.QueryRow(ctx, validateTags, arg.TagID, arg.CouponID)
	var not_exists bool
	err := row.Scan(&not_exists)
	return not_exists, err
}
