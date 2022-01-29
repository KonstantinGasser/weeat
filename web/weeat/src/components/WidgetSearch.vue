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
          @input="searchFood"
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
    <div v-if="isSearchFood">
      <div class="content">
        <div class="content-items">
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
      </div>
    </div>
    <div v-if="isSearchRecipe">
      <div class="content">
        <div class="content-items">
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
      </div>
    </div>
  </div>
</template>


<script>
import axios from "axios";

export default {
  name: "WidgetSeach",
  components: {},
  data() {
    return {
      isSearchFood: true,
      isSearchRecipe: false,
      emit_widget_name: "widget_close_search_food",
      query_food: [],
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
      this.query_query = "";
    },
    selectSearch(type) {
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
          this.query_food = resp?.data?.data;
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
  display: grid;
  row-gap: 15px;
  height: 65vh;
  overflow-y: scroll;
}
.food-item {
  padding: 15px 15px;
  box-shadow: 0 0 10px 2px rgb(0 0 0 / 10%);
  border-radius: 14px;
}

.nutrition-labels {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}
</style>