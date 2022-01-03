<template>
  <div class="widget">
    <div class="widget-header">
        <span>New Food</span>
        <i class="bi bi-x icon-medium" @click="$emit('widget_close_new_food')"></i>
    </div>
    <div class="">
      <div class="row g-2 my-1">
        <div class="col-md">
          <div class="form-floating">
            <input
              v-model="food_name"
              type="text"
              class="form-control"
              id="floatingInputGrid1"
              placeholder="Pumkin"
            />
            <label for="floatingInputGrid">Name of the Food</label>
          </div>
        </div>
        <div class="col-md">
          <div class="form-floating">
            <select
              v-model="food_cat"
              class="form-select"
              id="floatingSelectGrid"
              aria-label="Floating label select example"
            >
              <option value="1" selected>Fruit</option>
              <option value="2">Vegetable</option>
              <option value="3">Meat</option>
              <option value="4">Fish</option>
              <option value="5">Dairy</option>
              <option value="6">Grains</option>
            </select>
            <label for="floatingSelectGrid">Food Category</label>
          </div>
        </div>
      </div>
      <div class="row g-2 my-1">
        <div class="col-md">
          <div class="form-floating">
            <input
              v-model="food_kcal"
              type="text"
              inputmode="decimal"
              class="form-control"
              id="floatingInputGrid2"
              placeholder="24"
            />
            <label for="floatingInputGrid">Calorines (kcal)</label>
          </div>
        </div>
        <div class="col-md">
          <div class="form-floating">
            <input
              v-model="food_carbs"
              type="text"
              inputmode="decimal"
              class="form-control"
              id="floatingInputGrid3"
              placeholder="Pumkin"
            />
            <label for="floatingInputGrid">Carbohydrates (gramm)</label>
          </div>
        </div>
      </div>
      <div class="row g-2 my-1">
        <div class="col-md">
          <div class="form-floating">
            <input
              v-model="food_fats"
              type="text"
              inputmode="decimal"
              class="form-control"
              id="floatingInputGrid2"
              placeholder="24"
            />
            <label for="floatingInputGrid">Fats (gramm)</label>
          </div>
        </div>
        <div class="col-md">
          <div class="form-floating">
            <input
              v-model="food_protein"
              type="text"
              inputmode="decimal"
              class="form-control"
              id="floatingInputGrid3"
              placeholder="Pumkin"
            />
            <label for="floatingInputGrid">Protein (gramm)</label>
          </div>
        </div>
      </div>
      <div class="row g-2 d-flex justify-end my-2">
        <button class="action-btn" @click="addFood()">Add</button>
      </div>
    </div>
  </div>
</template>


<script>

export default {
  name: "WidgetAddFood",
  setup() {
  },
  data() {
    return {
        food_name: null,
        food_cat: 1,
        food_kcal: null,
        food_carbs: null,
        foot_protein: null,
        food_fats: null,
    };
  },
  methods: {
      addFood() {
        this.$moshaToast("Hello World", {type: 'success', position: 'top-center', timeout: 3000})
      }
  },
};
</script>