package postgres

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/weeat/core/dao"
	"github.com/KonstantinGasser/weeat/core/pkg/unit"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	ErrNotFound = fmt.Errorf("could not find item")
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
	logrus.Infof("[postgres.Conntect] conntecting to pg: %s\n", conn.uri(dbname))

	pgconf, err := pgxpool.ParseConfig(conn.uri(dbname))
	if err != nil {
		return errors.Wrap(err, "pg-sql - parsing pg config")
	}
	c, err := pgxpool.ConnectConfig(context.Background(), pgconf)
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
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?pool_max_conns=10", conn.user, conn.password, conn.host, conn.port, dbname)
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
	// to not be closed even if the batch results are not used
	// and required in the context
	// see: https://github.com/jackc/pgx/issues/635
	rows, _ := conn.c.Query(ctx, sql_get_food, ID)

	var item = dao.Food{}
	var kcal int
	var carbs, protein, fats, sugar float64
	if !rows.Next() {
		return item, ErrNotFound
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

func (conn *Conn) SearchFood(ctx context.Context, query string, limit int) ([]dao.FoodQuery, error) {

	rows, _ := conn.c.Query(ctx, sql_search_food, query, limit)
	var items []dao.FoodQuery
	var kcal int
	var carbs, sugar, protein, fats float64
	for rows.Next() {
		var item = dao.FoodQuery{}
		if err := rows.Scan(&item.ID, &item.Name, &item.Category, &kcal, &carbs, &sugar, &protein, &fats); err != nil {
			return items, errors.Wrap(err, "pg-sql - rows.Next, lookup item")
		}
		item.Kcal = unit.NewKcal(float64(kcal))
		item.Carbs = unit.NewGramm(carbs)
		item.Sugar = unit.NewGramm(sugar)
		item.Protein = unit.NewGramm(protein)
		item.Fats = unit.NewGramm(fats)

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

func (conn *Conn) InsertRecipe(ctx context.Context, recipe dao.Recipe) error {

	lock, err := conn.c.Acquire(ctx)
	if err != nil {
		return err
	}
	defer lock.Release()

	tx, err := lock.Begin(ctx)
	if err != nil {
		fmt.Println("begin: ", err)
		return err
	}

	var rowID int
	if err := tx.QueryRow(ctx, sql_insert_recipe, recipe.Name).Scan(&rowID); err != nil {
		fmt.Println("scan: ", err)
		return err
	}

	batch := pgx.Batch{}

	batch.Queue("insert into recipe_food_items(recipe_id,food_id, amount) values(32, 1, 100)")
	batch.Queue("insert into recipe_food_items(recipe_id,food_id, amount) values(32, 2, 100)")

	// dont forget to close batchResults
	// see: https://github.com/jackc/pgx/issues/610
	tx.SendBatch(ctx, &batch).Close()

	if err := tx.Commit(ctx); err != nil {
		fmt.Println("commit: ", err)
	}
	return nil
}

func (conn *Conn) MapFoodToRecipe(ctx context.Context, recipeID int, foods ...dao.Ingredient) error {

	tx, err := conn.c.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "pg-sql - start transaction")
	}

	batch := pgx.Batch{}

	for _, food := range foods {
		batch.Queue(sql_ref_food_recipe, recipeID, food.ID, food.Amount)
	}

	_ = tx.SendBatch(ctx, &batch)

	if err := tx.Commit(ctx); err != nil {
		return errors.Wrap(tx.Rollback(ctx), "pg-sql - rollback: commit to recipe_food_item")
	}
	return nil
}
