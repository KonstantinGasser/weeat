package postgres

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/weeat/core/dao"
	"github.com/KonstantinGasser/weeat/core/pkg/unit"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Conn struct {
	user, password string
	host           string
	port           int
	c              *pgxpool.Pool
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
	logrus.Infof("[postgres.Conntect] conntecting to pg: host=%s,db=%s\n", conn.host, dbname)

	c, err := pgxpool.Connect(context.Background(), conn.uri(dbname))
	if err != nil {
		return errors.Wrap(err, "connect to postgres-db caused an issue")
	}
	conn.c = c

	return nil
}

func (conn *Conn) Close(ctx context.Context) {
	conn.c.Close()
}

func (conn *Conn) uri(dbname string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", conn.user, conn.password, conn.host, conn.port, dbname)
}

// Service-Repo: Record-Serivce Functions

func (conn *Conn) InsertFood(ctx context.Context, food dao.Food) error {

	_, err := conn.c.Exec(
		ctx, sql_insert_food,
		food.Name,
		food.Label,
		food.Category,
		food.Kcal.Value(),
		food.Carbs.Value(),
		food.Sugar.Value(),
		food.Protein.Value(),
		food.Fats.Value(),
	)
	if err != nil {
		return errors.Wrap(err, "pg-sql - insert food")
	}
	return nil
}

func (conn *Conn) GetFood(ctx context.Context, ID string) (dao.Food, error) {
	// cannot use conn.c.QueryRow due to bug which causes the connection
	// to not be closed
	// see: https://github.com/jackc/pgx/issues/635
	rows, _ := conn.c.Query(ctx, sql_get_food, ID)

	var item = dao.Food{}
	var kcal int
	var carbs, protein, fats, sugar float64
	if !rows.Next() {
		return item, nil
	}
	defer rows.Close()

	err := rows.Scan(&item.ID, &item.Name, &item.Category, &kcal, &carbs, &sugar, &protein, &fats)
	if err != nil {
		return dao.Food{}, errors.Wrap(err, fmt.Sprintf("pg-sql could not lookup food with ID: %s", ID))
	}
	item.Kcal = unit.NewKcal(float64(kcal))
	item.Carbs = unit.NewGramm(carbs)
	item.Sugar = unit.NewGramm(sugar)
	item.Protein = unit.NewGramm(protein)
	item.Fats = unit.NewGramm(fats)

	return item, nil
}

func (conn *Conn) SearchFood(ctx context.Context, query string) ([]dao.FoodQuery, error) {

	rows, _ := conn.c.Query(ctx, sql_search_food, query)
	var items []dao.FoodQuery
	for rows.Next() {
		var item = dao.FoodQuery{}
		if err := rows.Scan(&item.ID, &item.Name, &item.Category); err != nil {
			return items, errors.Wrap(err, "pg-sql - rows.Next, lookup item")
		}

		items = append(items, item)
	}
	defer rows.Close()

	return items, rows.Err()
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
