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
	   select id, name, category, kcal, carbs, sugar, protein, fats from food_item where label like $1 || '%' limit $2;
	`

	sql_update_food = `
		update table food set $1 = $2;
	`

	sql_delete_food = `
		delete from food where id=$1;
	`

	sql_insert_recipe = `
		insert into recipe_item(name, category, label) values($1, $2, $3) returning id;
	`
	sql_ref_food_recipe = `
	    insert into recipe_food_items(recipe_id, food_id, amount)
			values($1,$2,$3);
	`

	sql_search_recipe = `
		select
		recipe.id, recipe.name, recipe.category, recipe_food_items.food_id, recipe_food_items.amount,
		food_item.name, food_item.category,
		(sum(food_item.kcal)::float/100)::int*50 as kcal,
		(sum(food_item.carbs)::float/100)*50 as carbs,
		(sum(food_item.sugar)::float/100)*50 as sugar,
		(sum(food_item.protein)::float/100)*50 as protein,
		(sum(food_item.fats)::float/100)*50 as fats
				from recipe_food_items
						inner join  (
								select recipe_item.id, recipe_item.name, recipe_item.category
									from recipe_item
								where recipe_item.label like '%' || $1 || '%') as recipe
						on recipe.id = recipe_food_items.recipe_id

						left join food_item on recipe_food_items.food_id=food_item.id
				group by recipe.id, recipe.name, recipe.category, food_item.name, recipe_food_items.food_id, food_item.category, recipe_food_items.amount
				limit $2;
	`

	sql_get_categories = `select id, category, item_type, emoji from categorys where item_type = $1`

	sql_insert_verify_item = `insert into verfiy_process(food_id, status, convidence) values($1, 0, 0.0)`
)
