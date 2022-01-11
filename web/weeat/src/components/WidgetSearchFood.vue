<template>
  <div class="widget">
    <div class="widget-header">
        <span>Seach Food</span>
        <i class="bi bi-x icon-medium" @click="closeWidget()"></i>
    </div>
    <div class="content">
      <div class="search">
        <div class="input-group">
          <input v-model="query_query" type="text" class="form-control" placeholder="Search for Food" aria-label="Search for Food" aria-describedby="basic-addon2">
          <div class="input-group-append">
            <button class="action-btn append" id="basic-addon2" @click="searchFood()">go..</button>
          </div>
        </div>
      </div>
      <div class="content-items">
        <div v-for="item in query_food" :key="item.id">
          {{item}}
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import axios from 'axios'

export default {
  name: "WidgetSeachFood",
  components: {
  },
  data() {
    return {
        emit_widget_name: "widget_close_search_food",
        query_food: [],
        query_query: "",
    };
  },
  methods: {
      closeWidget() {
        this.$emit(this.emit_widget_name)
        this.query_food = []
        this.query_query = ""
      },
      searchFood() {
        if (this.query_query.length === 0 ) return

        axios.get(
          process.env.VUE_APP_API + `/api/v1/food/search?q=${this.query_query}&l=${process.env.VUE_APP_FOOD_SEARCH_LIMIT}`
        ).then(resp => {
          this.query_food = resp?.data?.data
        })
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