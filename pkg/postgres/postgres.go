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

func (conn *Conn) SearchRecipe(ctx context.Context, query string, limit int) ([]dao.Recipe, error) {
	rows, err := conn.c.Query(ctx, sql_search_recipe, query, limit)
	if err != nil {
		return nil, errors.Wrap(err, "pg-sql - search for recipes")
	}
	defer rows.Close()

	type row struct {
		recID      int
		recName    string
		recCat     int
		ingID      int
		ingAmount  float64
		ingName    string
		ingCat     int
		ingKcal    int
		ingCarbs   float64
		ingSugar   float64
		ingProtein float64
		ingFats    float64
	}

	var recipeIngs = make(map[int]*dao.Recipe)
	for rows.Next() {
		var r = row{}
		if err := rows.Scan(
			&r.recID, &r.recName, &r.recCat,
			&r.ingID, &r.ingAmount, &r.ingName, &r.ingCat,
			&r.ingKcal, &r.ingCarbs, &r.ingSugar, &r.ingProtein, &r.ingFats,
		); err != nil {
			return nil, errors.Wrap(err, "pg-sql - scan searched recipe rows")
		}

		if _, ok := recipeIngs[r.recID]; !ok {
			recipeIngs[r.recID] = &dao.Recipe{
				ID:       r.recID,
				Name:     r.recName,
				Category: r.recCat,
				FoodIDs: []dao.Ingredient{
					{
						ID:     r.ingID,
						Amount: r.ingAmount,
						Food: dao.FoodQuery{
							Name:     r.ingName,
							Category: r.ingCat,
							Kcal:     unit.NewKcal(float64(r.ingKcal)),
							Carbs:    unit.NewGramm(r.ingCarbs),
							Sugar:    unit.NewGramm(r.ingSugar),
							Protein:  unit.NewGramm(r.ingProtein),
							Fats:     unit.NewGramm(r.ingFats),
						},
					},
				},
			}
			continue
		}

		recipeIngs[r.recID].FoodIDs = append(recipeIngs[r.recID].FoodIDs, dao.Ingredient{
			ID:     r.ingID,
			Amount: r.ingAmount,
			Food: dao.FoodQuery{
				Name:     r.ingName,
				Category: r.ingCat,
				Kcal:     unit.NewKcal(float64(r.ingKcal)),
				Carbs:    unit.NewGramm(r.ingCarbs),
				Sugar:    unit.NewGramm(r.ingSugar),
				Protein:  unit.NewGramm(r.ingProtein),
				Fats:     unit.NewGramm(r.ingFats),
			},
		})

	}

	var recipes []dao.Recipe
	for _, item := range recipeIngs {
		recipes = append(recipes, *item)
	}
	return recipes, nil
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
		return errors.Wrap(err, "pg-sql - insert recipe, acquire lock")
	}
	defer lock.Release()

	tx, err := lock.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "pg-sql - insert recipe, begin transaction")
	}

	var rowID int
	if err := tx.QueryRow(ctx, sql_insert_recipe, recipe.Name, recipe.Category, recipe.Label).Scan(&rowID); err != nil {
		return errors.Wrap(err, "pg-sql - insert recipe, scanning returning id")
	}

	batch := pgx.Batch{}

	for _, item := range recipe.FoodIDs {
		batch.Queue(sql_ref_food_recipe, rowID, item.ID, item.Amount)
	}

	// don't forget to close batchResults
	// else the connection will say connective and thus
	// will not be returned to the pgxcool due to "conn busy" error
	// Close() can also not be deferred unless the Close happens
	// before the commit of the transaction
	//
	// see: https://github.com/jackc/pgx/issues/610
	tx.SendBatch(ctx, &batch).Close()

	if err := tx.Commit(ctx); err != nil {
		return errors.Wrap(err, "pg-sql - insert recipe commit transaction")
	}
	return nil
}

func (conn *Conn) GetCategories(ctx context.Context, _type int) ([]dao.Category, error) {

	rows, err := conn.c.Query(ctx, sql_get_categories, _type)
	if err != nil {
		return nil, errors.Wrap(err, "pg-sql - select categories")
	}
	defer rows.Close()

	var cats []dao.Category
	for rows.Next() {
		var cat dao.Category
		if err := rows.Scan(&cat.ID, &cat.Label, &cat.Type, &cat.Emoji); err != nil {
			return nil, errors.Wrap(err, "pg-sql - scan category rows")
		}
		cats = append(cats, cat)
	}
	return cats, nil
}

func (conn *Conn) InitVerifyItem(ctx context.Context, item dao.VerifyItem) error {
	return nil
}
