<template>
    <div class="main-frame">
        <div class="action-row">
            <div class="headline">
                <h1>Hello <span>Friend</span></h1>
            </div>
            <div class="btns">
                <button class="action-btn medium" @click="isAddFood=!isAddFood">Add Food</button>
            </div>
        </div>
        <div v-if="isAddFood" class="content">
           <div class="subheade">
               <h4>Add a Food</h4>
               <div class="form-grid">
                   <div class="row g-2 my-1">
                        <div class="col-md">
                            <div class="form-floating">
                            <input v-model="food_name" type="text" class="form-control" id="floatingInputGrid1" placeholder="Pumkin" >
                            <label for="floatingInputGrid">Name of the Food</label>
                            </div>
                        </div>
                        <div class="col-md">
                            <div class="form-floating">
                            <select v-model="food_cat" class="form-select" id="floatingSelectGrid" aria-label="Floating label select example">
                                <option value="1" selected >Fruit</option>
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
                                <input v-model="food_kcal" type="decimal" class="form-control" id="floatingInputGrid2" placeholder="24">
                                <label for="floatingInputGrid">Calorines (kcal)</label>
                            </div>
                        </div>
                        <div class="col-md">
                            <div class="form-floating">
                                <input v-model="food_carbs" type="decimal" class="form-control" id="floatingInputGrid3" placeholder="Pumkin" >
                                <label for="floatingInputGrid">Carbohydrates (gramm)</label>
                            </div>
                        </div>
                    </div>
                    <div class="row g-2 my-1">
                        <div class="col-md">
                            <div class="form-floating">
                                <input v-model="food_fats" type="decimal" class="form-control" id="floatingInputGrid2" placeholder="24" >
                                <label for="floatingInputGrid">Fats (gramm)</label>
                            </div>
                        </div>
                        <div class="col-md">
                            <div class="form-floating">
                                <input v-model="food_protein" type="decimal" class="form-control" id="floatingInputGrid3" placeholder="Pumkin" >
                                <label for="floatingInputGrid">Protein (gramm)</label>
                            </div>
                        </div>
                    </div>
                    <div class="row g-2 d-flex justify-center my-1">
                        <button class="action-btn" @click="addFruit()">Add</button>
                    </div>
               </div>
           </div>
        </div>
    </div>
</template>

<script lang="ts">

// import axios from 'axios';

export default {
    name: "Dashboard",
    data() {
        return {
            isAddFood: true,
            food_name: "",
            food_cat: 1,
            food_kcal: null,
            food_carbs: null,
            food_fats: null,
            food_protein: null,
        };
    },
    methods: {
        addFruit(): void {
            this.food_name = "hello";
        },
    }
}

</script>

<style scoped>

</style>