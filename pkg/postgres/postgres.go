package postgres

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/lowco/core/dao"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type Conn struct {
	user, password string
	host           string
	port           int
	c              *pgx.Conn
}

func New(user, password, host string, port int) *Conn {
	return &Conn{
		user:     user,
		password: password,
		host:     host,
		port:     port,
	}
}

func (conn *Conn) Connect(dbname string) error {
	c, err := pgx.Connect(context.Background(), conn.uri(dbname))
	if err != nil {
		return errors.Wrap(err, "connect to postgres-db caused an issue")
	}
	conn.c = c

	return nil
}

func (conn *Conn) Close(ctx context.Context) error {
	return conn.c.Close(ctx)
}

func (conn *Conn) uri(dbname string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", conn.user, conn.password, conn.host, conn.port, dbname)
}

// Service-Repo: Record-Serivce Functions

func (conn *Conn) InsertFood(ctx context.Context, food dao.Food) error {

	_, err := conn.c.Exec(
		ctx, sql_insert_food,
		food.Name,
		food.Category,
		food.Carbs,
		food.Protein,
		food.Fats,
	)
	if err != nil {
		return errors.Wrap(err, "pg-sql - insert food")
	}
	return nil
}

func (conn *Conn) UpdateFood(ctx context.Context, column string, value interface{}) error {

	_, err := conn.c.Exec(
		ctx, sql_update_food,
		column,
		value,
	)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("pg-sql - update food: column %q, value: %q", column, value))
	}
	return nil
}

func (conn *Conn) DeleteFood(ctx context.Context, ID int) error {

	_, err := conn.c.Exec(
		ctx, sql_delete_food,
		ID,
	)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("pg-sql - delete food: ID %q", ID))
	}
	return nil
}
