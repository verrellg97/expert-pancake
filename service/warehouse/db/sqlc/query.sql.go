// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package db

import (
	"context"
)

const getRacks = `-- name: GetRacks :many
SELECT a.id, a.branch_id, a.name
FROM warehouse.racks a
LEFT JOIN warehouse.warehouse_racks b ON a.id = b.rack_id
WHERE a.branch_id = $1 
    AND a.name LIKE $2 
    AND CASE WHEN
        $3::bool THEN b.rack_id IS NULL ELSE TRUE END
`

type GetRacksParams struct {
	BranchID       string `db:"branch_id"`
	Name           string `db:"name"`
	Isgetavailable bool   `db:"isgetavailable"`
}

type GetRacksRow struct {
	ID       string `db:"id"`
	BranchID string `db:"branch_id"`
	Name     string `db:"name"`
}

func (q *Queries) GetRacks(ctx context.Context, arg GetRacksParams) ([]GetRacksRow, error) {
	rows, err := q.db.QueryContext(ctx, getRacks, arg.BranchID, arg.Name, arg.Isgetavailable)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRacksRow
	for rows.Next() {
		var i GetRacksRow
		if err := rows.Scan(&i.ID, &i.BranchID, &i.Name); err != nil {
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

const getWarehouses = `-- name: GetWarehouses :many
SELECT a.id, a.branch_id, a.code, a.name, a.address, a.type, a.is_deleted, a.created_at, a.updated_at
FROM warehouse.warehouses a
WHERE  a.branch_id = $1
AND a.name LIKE $2
AND a.type LIKE $3
AND a.is_deleted = 0
`

type GetWarehousesParams struct {
	BranchID string `db:"branch_id"`
	Name     string `db:"name"`
	Type     string `db:"type"`
}

func (q *Queries) GetWarehouses(ctx context.Context, arg GetWarehousesParams) ([]WarehouseWarehouse, error) {
	rows, err := q.db.QueryContext(ctx, getWarehouses, arg.BranchID, arg.Name, arg.Type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []WarehouseWarehouse
	for rows.Next() {
		var i WarehouseWarehouse
		if err := rows.Scan(
			&i.ID,
			&i.BranchID,
			&i.Code,
			&i.Name,
			&i.Address,
			&i.Type,
			&i.IsDeleted,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const upsertRack = `-- name: UpsertRack :one
INSERT INTO warehouse.racks(id, branch_id, name)
VALUES ($1, $2, $3)
ON CONFLICT (id)
DO UPDATE SET
    name = EXCLUDED.name,
    branch_id = EXCLUDED.branch_id,
    updated_at = NOW()
RETURNING id, branch_id, name, created_at, updated_at
`

type UpsertRackParams struct {
	ID       string `db:"id"`
	BranchID string `db:"branch_id"`
	Name     string `db:"name"`
}

func (q *Queries) UpsertRack(ctx context.Context, arg UpsertRackParams) (WarehouseRack, error) {
	row := q.db.QueryRowContext(ctx, upsertRack, arg.ID, arg.BranchID, arg.Name)
	var i WarehouseRack
	err := row.Scan(
		&i.ID,
		&i.BranchID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
