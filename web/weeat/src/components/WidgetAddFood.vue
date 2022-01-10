<template>
  <div class="widget">
    <div class="widget-header">
        <span>New Food <br><small>(per 100{{unitText}})</small></span>
        <i class="bi bi-x icon-medium" @click="closeWidget()"></i>
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
              <option value="7">Soft-Drinks</option>
              <option value="8">Alcohol</option>
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
            <label for="floatingInputGrid">Carbohydrates (gramm) 🍞</label>
          </div>
        </div>
        <div class="form-group">
          <input v-model="food_sugar" type="text" inputmode="decimal" class="form-control" id="carbSugar" aria-describedby="sugarHelp" placeholder="Sugar 🍩">
          <small id="sugarHelp" class="form-text text-muted">Amount of sugar cannot be bigger then total carbs</small>
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
            <label for="floatingInputGrid">Fats (gramm) 🧈</label>
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
            <label for="floatingInputGrid">Protein (gramm) 🥜</label>
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

import axios from 'axios'

export default {
  name: "WidgetAddFood",
  data() {
    return {
        emit_widget_name: "widget_close_new_food",
        food_name: null,
        food_cat: 1,
        food_kcal: null,
        food_carbs: null,
        food_sugar: null,
        foot_protein: null,
        food_fats: null,
    };
  },
  unmounted() {
    console.log("destroying: new.Food")
  },
  computed: {
    unitText() {
      return this.food_cat < 7 ? "g" : "ml";
    }
  },
  methods: {
      closeWidget() {
        this.$emit(this.emit_widget_name)
        this.$options.unmounted()
      },
      addFood() {
        let options = {
            headers: {
                'Content-Type': 'application/json',
            }
        };

        const payload = {
            name: this.food_name,
            food_cat: parseInt(this.food_cat),
            kcal: parseFloat(this.food_kcal),
            carbs: parseFloat(this.food_carbs),
            sugar: parseFloat(this.food_sugar),
            fats: parseFloat(this.food_fats),
            protein: parseFloat(this.food_protein),
        }
  
        axios.post("http://localhost:30024/records/new/food", payload, options).then(resp =>{
          this.$moshaToast(resp?.data, {type: 'success',position: 'top-center', timeout: 3000})
          this.$emit('widget_close_new_food')

        }).catch(err => {
          this.$moshaToast(err?.response?.data, {type: 'danger',position: 'top-center', timeout: 3000})
        })
      }
  },
};
</script>