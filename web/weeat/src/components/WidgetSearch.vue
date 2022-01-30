<template>
  <div class="widget">
    <div class="widget-header">
      <span>Search...🔎</span>
      <i class="bi bi-x icon-medium" @click="closeWidget()"></i>
    </div>
    <div class="widget-menu">
      <div :class="{ selected: isSearchFood }" @click="selectSearch('food')">
        ...Food
      </div>
      <div
        :class="{ selected: isSearchRecipe }"
        @click="selectSearch('recipe')"
      >
        ...Recipe
      </div>
    </div>
    <div class="search">
      <div class="input-group">
        <input
          v-model="query_query"
          type="text"
          @input="search"
          class="form-control"
          :placeholder="searchPlaceholder"
          aria-label="Search for Food"
          aria-describedby="basic-addon2"
        />
        <div class="input-group-append">
          <button
            class="action-btn append"
            id="basic-addon2"
            @click="searchFood"
          >
            go..
          </button>
        </div>
      </div>
    </div>
    <div class="content">
      <div v-if="isSearchFood" class="content-items">
        <div v-for="item in query_food" :key="item.id" class="food-item">
          <h5>
            {{ item.name }}
            <span v-if="item.category === 1">🍒</span>
            <span v-if="item.category === 2">🥦</span>
            <span v-if="item.category === 3">🍗</span>
            <span v-if="item.category === 4">🍣</span>
            <span v-if="item.category === 5">🧀</span>
            <span v-if="item.category === 6">🍞</span>
            <span v-if="item.category === 7">🧃</span>
            <span v-if="item.category === 8">🍹</span>
          </h5>
          <div class="nutrition-labels">
            <div class="nutrition_tag tag_kcal">{{ item.kcal }} (kcal)</div>
            <div class="nutrition_tag tag_carbs">
              {{ item.carbs.toFixed(2) }}g (carbohydrates)
            </div>
            <div class="nutrition_tag tag_carbs">
              {{ item.sugar.toFixed(2) }}g (carbohydrates)
            </div>
            <div class="nutrition_tag tag_fats">
              {{ item.fats.toFixed(2) }}g (fat)
            </div>
            <div class="nutrition_tag tag_protein">
              {{ item.protein.toFixed(2) }}g (protein)
            </div>
          </div>
        </div>
      </div>
      <div v-if="isSearchRecipe" class="content-items">
        <div
          v-for="item in query_recipe"
          :key="item.id"
          class="food-item"
          @click="expand(item.id)"
        >
          <h5>
            {{ item.name }}
            <span v-if="item.category === 9">🥞</span>
            <span v-if="item.category === 10">🌯</span>
            <span v-if="item.category === 11">🥘</span>
            <span v-if="item.category === 12">🍿</span>
          </h5>
          <div class="list_ingredients">
            <div
              v-for="i in item.ingredients"
              :key="i.id"
              class="item_ingredient"
            >
              {{ i.name }} {{ i.amount }}
              <span v-if="i.category < 7">g</span>
              <span v-if="i.category >= 7">ml</span>
              <span v-if="i.category === 1">&nbsp;🍒</span>
              <span v-if="i.category === 2">&nbsp;🥦</span>
              <span v-if="i.category === 3">&nbsp;🍗</span>
              <span v-if="i.category === 4">&nbsp;🍣</span>
              <span v-if="i.category === 5">&nbsp;🧀</span>
              <span v-if="i.category === 6">&nbsp;🍞</span>
              <span v-if="i.category === 7">&nbsp;🧃</span>
              <span v-if="i.category === 8">&nbsp;🍹</span>
            </div>
          </div>
          <div class="recipe_info" :class="{ expand: item.id === expandBox }">
            <hr>
            <h6>Nutritions</h6>
            <div>
              <span class="nutrition_tag tag_kcal">{{
                item.nutritions.kcal.toFixed(2)
              }}</span>
              <span class="nutrition_tag tag_carbs">{{
                item.nutritions.carbs.toFixed(2)
              }}</span>
              <span class="nutrition_tag tag_sugar">{{
                item.nutritions.sugar.toFixed(2)
              }}</span>
              <span class="nutrition_tag tag_protein">{{
                item.nutritions.protein.toFixed(2)
              }}</span>
              <span class="nutrition_tag tag_fats">{{
                item.nutritions.fats.toFixed(2)
              }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- </div> -->
  </div>
</template>


<script>
import axios from "axios";

export default {
  name: "WidgetSeach",
  components: {},
  data() {
    return {
      expandBox: -1, // could be null to must not be any number >= 0, loop var start at 0
      isSearchFood: true,
      isSearchRecipe: false,
      emit_widget_name: "widget_close_search_food",
      query_food: [],
      query_recipe: [],
      query_query: "",
    };
  },
  computed: {
    searchPlaceholder() {
      if (this.isSearchFood) return "Search Food...";
      return "Search Recipe";
    },
  },
  methods: {
    closeWidget() {
      this.$emit(this.emit_widget_name);
      this.query_food = [];
      this.query_recipe = [];
      this.query_query = "";
      this.isSearchFood = true;
      this.isSearchRecipe = false;
    },
    expand(id) {
      if (this.expandBox === id) {
        this.expandBox = -1;
        return;
      }
      this.expandBox = id;
    },
    selectSearch(type) {
      this.query_query = "";
      switch (type) {
        case "food":
          this.isSearchFood = !this.isSearchFood;
          this.isSearchRecipe = false;
          break;
        case "recipe":
          this.isSearchRecipe = !this.isSearchRecipe;
          this.isSearchFood = false;
          break;
      }
    },
    search() {
      if (this.isSearchFood) {
        this.searchFood();
        return;
      }
      this.searchRecipe();
    },
    searchFood() {
      if (this.query_query.length === 0) {
        this.query_food = [];
      }
      axios
        .get(
          process.env.VUE_APP_API +
            `/api/v1/food/search?q=${this.query_query}&l=${process.env.VUE_APP_FOOD_SEARCH_LIMIT}`
        )
        .then((resp) => {
          console.log(resp?.data);
          this.query_food = resp?.data?.data;
        });
    },
    searchRecipe() {
      if (this.query_query.length === 0) {
        this.query_recipe = [];
      }

      axios
        .get(
          process.env.VUE_APP_API +
            `/api/v1/recipe/search?q=${this.query_query}&l=${process.env.VUE_APP_FOOD_SEARCH_LIMIT}`
        )
        .then((resp) => {
          console.log(resp?.data);
          this.query_recipe = resp?.data?.data;
        });
    },
  },
};
</script>

<style scoped>
.widget {
  padding: 0 0 5px 0 !important;
  height: 80vh !important;
  max-height: 80vh !important;
}
.widget-header {
  padding: 0 15px;
}
.search {
  padding: 0 15px;
}

.content-items {
  padding: 15px;
  display: flex;
  flex-direction: column;
  row-gap: 15px;
  height: 60vh;
  overflow-y: scroll;
}
.food-item {
  padding: 15px 15px;
  box-shadow: 0 0 10px 2px rgb(0 0 0 / 10%);
  background: var(--box-bg);
  border-radius: 14px;
  height: min-content;
}
.recipe_info {
  height: 0px;
  opacity: 0;
  transition: 100ms ease-out;
}

.recipe_info>div{
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
}
.recipe_info.expand {
  padding: 15px 0;
  height: max-content;
  opacity: 1;
}

.nutrition-labels {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}
</style>