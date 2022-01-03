package postgres

var (
	sql_insert_food = `
		insert into foods(name, category, kcal, carbs, protein, fats)
			values(
				$1,
				$2,
				$3,
				$4,
				$5,
				$6
			);
	`

	sql_update_food = `
		update table food set $1 = $2;
	`

	sql_delete_food = `
		delete from food where id=$1;
	`
)
