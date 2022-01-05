<template>
  <div class="widget">
    <div class="widget-header">
        <span>New Recipe</span>
        <i class="bi bi-x icon-medium" @click="closeWidget()"></i>
    </div>
    <div class="">
      <div class="row g-2 my-1">
        <div class="col-md">
          <div class="form-floating">
            <input
              v-model="recipe_name"
              type="text"
              class="form-control"
              id="floatingInputGrid1"
              placeholder="Pumkin"
            />
            <label for="floatingInputGrid">Recipe Name</label>
          </div>
        </div>
        <div class="col-md">
          <div class="form-outline">
            <input v-model="searched_ingredient" type="search" id="form1" class="form-control" @keyup.enter="addIngredient()" placeholder="Search for Ingredient"
            aria-label="Search" />
          </div>
        </div>
      </div>
      <div class="row g-2 my-1 ingredients">
        <span class="list_header">Ingredients...</span>
        <div v-if="ingredients.length === 0" class="nothing-selected">
          You have not selected any ingredients for the recipe 🙃
        </div>
        <div class="list_ingredients">
          <div v-for="i in ingredients" :key="i" class="item_ingredient">
            {{i}}
            <i class="bi bi-x" @click="removeIngredient(i)"></i>
          </div>
        </div>
      </div>
      <div class="row g-2 my-1 recipe_info">
        <div class="recipe_tag tag_kcal">245 (kcal)</div>
        <div class="recipe_tag tag_carbs">25g (carbohydrates)</div>
        <div class="recipe_tag tag_fats">2.4g (fat)</div>
        <div class="recipe_tag tag_protein">24.5g (protein)</div>
      </div>
      <div class="row g-2 d-flex justify-end my-2">
        <button class="action-btn" @click="addRecipe()">Add</button>
      </div>
    </div>
  </div>
</template>


<script>
import axios from 'axios'

export default {
  name: "WidgetAddFood",
  data() {
    return {
        emit_widget_name: "widget_close_new_recipe",
        recipe_name: null,
        searched_ingredient: null,
        ingredients: [],
    };
  },
  unmounted() {
    console.log("destroying: new.Recipe")
  },
  methods: {
      closeWidget() {
        this.$emit(this.emit_widget_name)
        this.$options.unmounted()
      },
      addIngredient() {
        this.ingredients.push(this.searched_ingredient)
        this.searched_ingredient = null
        console.log(this.ingredients)
      },
      removeIngredient(item) {
        this.ingredients = this.ingredients.filter(i => (i != item))
      },
      addRecipe() {
        let options = {
            headers: {
                'Content-Type': 'application/json',
            }
        };
        const payload = {
          name: this.recipe_name,
          ingredients: this.ingredients,
        }
        
        axios.post("http://localhost:8000/records/new/recipe", payload, options).then(resp => {
          this.$moshaToast(resp?.data, {type: 'success',position: 'top-center', timeout: 3000})
          this.$emit('widget_close_new_recipe')
        }).catch(err => {
          this.$moshaToast(err, {type:'success', position: 'top-center', timeout: 3000})
        })
      },
  },
};
</script>

<style scoped>

.ingredients .list_header {
  text-align: left;
}

.list_ingredients {
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
}

.list_ingredients .item_ingredient {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 3px 5px;
  width: max-content;
  border-radius: 24px;
  background: #1fcf8025;
  color: #1fcf80;
  border: 1px solid #1fcf80;
}

.item_ingredient .bi {
  padding-top: 2px;
}


.recipe_info {
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
  justify-content: center;
}

.recipe_info .recipe_tag {
  width: max-content;
  padding: 3px 10px;
  font-weight: bolder;
  border-radius: 14px;
}

.recipe_tag.tag_kcal {
  background: #ef233c25;
  color: #ef233c;
  border: 1px solid #ef233c;
}

.recipe_tag.tag_carbs {
  background: #ffadad25;
  color: #ffadad;
  border: 1px solid #ffadad;
}

.recipe_tag.tag_fats {
  background: #ffd6a525;
  color: #ffd6a5;
  border: 1px solid #ffd6a5;
}

.recipe_tag.tag_protein {
  background: #1fcf8025;
  color: #1fcf80;
  border: 1px solid #1fcf80;
}

</style>