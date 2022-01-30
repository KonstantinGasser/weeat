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
          <div class="form-floating">
            <select
              v-model="recipe_cat"
              class="form-select"
              id="floatingSelectGrid"
              aria-label="Floating label select example"
            >
              <option value="0" selected>Breakfast 🥞</option>
              <option value="10">Lunch 🌯</option>
              <option value="11">Dinner 🥘</option>
              <option value="12">Snack 🍿</option>
            </select>
            <label for="floatingSelectGrid">Category</label>
          </div>
        </div>
        <div class="col-md">
          <div class="input-group">
            <input
              v-model="query_query"
              @input="searchFood"
              @change="searchFood"
              type="search"
              class="form-control"
              placeholder="Search"
              aria-label="Search"
              aria-describedby="basic-addon2"
            />
            <input
              v-model="scaler"
              type="search"
              class="form-control"
              placeholder="gramm/ml"
              aria-label="gramms"
              aria-describedby="basic-addon2"
            />
          </div>
          <div>
            <ul class="list-query-food">
              <li
                v-for="item in query_food"
                :key="item.id"
                @click="addIngredient(item.id)"
              >
                {{ item.name }}
                <span v-if="item.category === 1">🍒</span>
                <span v-if="item.category === 2">🥦</span>
                <span v-if="item.category === 3">🍗</span>
                <span v-if="item.category === 4">🍣</span>
                <span v-if="item.category === 5">🧀</span>
                <span v-if="item.category === 6">🍞</span>
                <span v-if="item.category === 7">🧃</span>
                <span v-if="item.category === 8">🍹</span>
              </li>
            </ul>
          </div>
        </div>
      </div>
      <div class="row g-2 my-1 ingredients">
        <span class="list_header">Ingredients...</span>
        <div v-if="ingredients.length === 0" class="nothing-selected">
          You have not selected any ingredients for the recipe 🙃
        </div>
        <div class="list_ingredients">
          <div
            v-for="item in ingredients"
            :key="item.id"
            class="item_ingredient"
          >
            <i class="bi bi-x" @click="removeIngredient(item)"></i>
            {{ item.name }} {{ item.amount }}
            <span v-if="item.category < 7">g</span>
            <span v-if="item.category >= 7">ml</span>
            <span v-if="item.category === 1">&nbsp;🍒</span>
            <span v-if="item.category === 2">&nbsp;🥦</span>
            <span v-if="item.category === 3">&nbsp;🍗</span>
            <span v-if="item.category === 4">&nbsp;🍣</span>
            <span v-if="item.category === 5">&nbsp;🧀</span>
            <span v-if="item.category === 6">&nbsp;🍞</span>
            <span v-if="item.category === 7">&nbsp;🧃</span>
            <span v-if="item.category === 8">&nbsp;🍹</span>
          </div>
        </div>
      </div>
      <hr />
      <div class="row g-2 my-1 recipe_info">
        <div class="nutrition_tag tag_kcal">{{ recipe_kcal }} (kcal)</div>
        <div class="nutrition_tag tag_carbs">
          {{ recipe_carbs.toFixed(2) }}g (carbohydrates)
        </div>
        <div class="nutrition_tag tag_fats">
          {{ recipe_fats.toFixed(2) }}g (fat)
        </div>
        <div class="nutrition_tag tag_protein">
          {{ recipe_protein.toFixed(2) }}g (protein)
        </div>
      </div>
      <div class="row g-2 d-flex justify-end my-2">
        <button class="action-btn" @click="addRecipe()">Add</button>
      </div>
    </div>
  </div>
</template>


<script>
import axios from "axios";

export default {
  name: "WidgetAddFood",
  components: {},
  data() {
    return {
      emit_widget_name: "widget_close_new_recipe",
      recipe_name: null,
      query_food: [],
      query_query: "",
      scaler: null,
      recipe_cat: 9,
      ingredients: [],
      recipe_kcal: 0,
      recipe_carbs: 0,
      recipe_protein: 0,
      recipe_fats: 0,
    };
  },
  methods: {
    closeWidget() {
      this.$emit(this.emit_widget_name);
      this.ingredients = [];
      this.query_food = [];
      this.query_query = "";
      this.recipe_kcal = 0;
      this.recipe_carbs = 0;
      this.recipe_protein = 0;
      this.recipe_fats = 0;
      this.scaler = null;
      this.recipe_name = "";
    },
    searchFood() {
      if (this.query_query.length === 0) return;

      axios
        .get(
          process.env.VUE_APP_API +
            `/api/v1/food/search?q=${this.query_query}&l=${process.env.VUE_APP_FOOD_SEARCH_LIMIT}`
        )
        .then((resp) => {
          this.query_food = resp?.data?.data?.filter((item) => {
            this.query_food?.forEach((i) => {
              if (i.id === item.id) return false;
            });
            return true;
          });
        });
    },
    addIngredient(id) {
      // add only if not exists
      if (this.ingredients?.filter((i) => i.ID === id).length > 0) {
        return;
      }
      // error out if unit is not set
      if (this.scaler === null) {
        this.$moshaToast("unit cannot be empty...", {
          type: "danger",
          position: "top-center",
          timeout: 3000,
        });
        return;
      }

      axios
        .get(
          process.env.VUE_APP_API +
            `/api/v1/food?id=${id}&scaler=${this.scaler}`
        )
        .then((resp) => {
          let item = {};
          item = resp?.data?.data;
          item.amount = this.scaler;

          this.ingredients.push(item);

          this.recipe_kcal += item.kcal;
          this.recipe_carbs += item.carbs;
          this.recipe_protein += item.protein;
          this.recipe_fats += item.fats;

          this.query_query = "";
          this.query_food = [];
          this.scaler = null;
        })
        .catch((err) => {
          this.$moshaToast(err, {
            type: "danger",
            position: "top-center",
            timeout: 3000,
          });
          return;
        });
    },
    removeIngredient(item) {
      this.ingredients = this.ingredients.filter((i) => i != item);
      this.recipe_kcal -= item.kcal;
      this.recipe_carbs -= item.carbs;
      this.recipe_protein -= item.protein;
      this.recipe_fats -= item.fats;
    },
    addRecipe() {
      console.log();
      let options = {
        headers: {
          "Content-Type": "application/json",
        },
      };
      const payload = {
        name: this.recipe_name,
        category: parseInt(this.recipe_cat),
        ingredients: this.ingredients.map((item) => {
          return {
            id: item.id,
            amount: parseInt(item.amount),
          };
        }),
      };

      axios
        .post(process.env.VUE_APP_API + `/api/v1/recipe`, payload, options)
        .then((resp) => {
          this.$moshaToast(resp?.data, {
            type: "success",
            position: "top-center",
            timeout: 3000,
          });
          this.$emit("widget_close_new_recipe");
        })
        .catch((err) => {
          this.$moshaToast(err, {
            type: "danger",
            position: "top-center",
            timeout: 3000,
          });
        });
    },
  },
};
</script>

<style scoped>
.ingredients .list_header {
  text-align: left;
  height: 100%;
  max-height: 470px;
  overflow-y: scroll;
}

.list_ingredients {
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
}

/* .list_ingredients .item_ingredient {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 3px 5px;
  width: max-content;
  border-radius: 24px;
  background: #1fcf8025;
  color: #1fcf80;
  border: 1px solid #1fcf80;
} */

/* .item_ingredient .bi {
  padding-top: 2px;
} */

.list-query-food {
  margin: 10px 0;
  padding: 0px;

  display: flex;
  flex-wrap: wrap;
  gap: 5px;

  text-decoration: none;
  list-style: none;
}

.list-query-food li {
  padding: 3px 5px;
  border-radius: 25px;
  background: #1fcf80;
  color: #ffffff;
}

.recipe_info {
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
  justify-content: center;
}
</style>