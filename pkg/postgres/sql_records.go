package postgres

var (
	sql_insert_food = `
		insert into food_item(name, label, category, kcal, carbs, sugar, protein, fats)
			values(
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7,
				$8
			);
	`
	sql_get_food = `
	   select id, name, category, kcal, carbs, sugar, protein, fats from food_item where id=$1;
	`
	sql_search_food = `
	   select id, name, category from food_item where label like $1 || '%';
	`

	sql_update_food = `
		update table food set $1 = $2;
	`

	sql_delete_food = `
		delete from food where id=$1;
	`
)
