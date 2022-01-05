<template>
    <div class="main-frame">
        <div class="headline">
            <h1><span>We</span>Eat</h1>
        </div>
        <!-- <div class="action-row">
            <button class="action-btn medium" @click="isAddFood=!isAddFood">Add Food</button>
            <button class="action-btn medium" @click="isAddRecipe=!isAddRecipe">Add Recipe</button>
        </div> -->
        <div class="dashboard">
            <div class="dashboard-options">
                <div class="dashboard-option" @click="isAddFood=!isAddFood">
                    <span>Add Food</span>
                    <span class="option-icon">🫐</span>
                    <!-- <i class="bi bi-dice-6"></i> -->
                </div>
                <div class="dashboard-option" @click="isAddRecipe=!isAddRecipe">
                    <span>Add Recipe</span>
                    <span class="option-icon">🌮</span>
                    <!-- <i class="bi bi-bullseye"></i> -->
                </div>
                <div class="dashboard-option" @click="isGenerateMeals=!isGenerateMeals">
                    <span>Feeling lucky?</span>
                    <span class="option-icon">😬</span>
                    <!-- <i class="bi bi-dice-6"></i> -->
                </div>
                <div class="dashboard-option">
                    <span>Track intake</span>
                    <span class="option-icon">🍽</span>
                    <!-- <i class="bi bi-bullseye"></i> -->
                </div>
            </div>
            <div class="intake-today mt-2">
                <hr>
                <h4>{{info_text()}}<br> you have <span>1600</span> kcal today</h4>
                <IntakeDoughnut v-bind:chartData="state.chartData" v-bind:chartOptions="state.chartOptions" />
            </div>
            <!-- <div class="nothing-there-yet-text">
                <span>
                    There seems to be no data..go and do something
                </span>
            </div>
            <div class="nothing-there-yet-img">
                <img class="svg-img" src="../assets/svg/cooking.svg" alt="">
            </div> -->
        </div>
        <WidgetAddFood :class="{'widget-active':isAddFood}" @widget_close_new_food="isAddFood=!isAddFood"/>
        <WidgetAddRecipe :class="{'widget-active':isAddRecipe}" @widget_close_new_recipe="isAddRecipe=!isAddRecipe"/>
        <WidgetGenerateMeals :class="{'widget-active':isGenerateMeals}" @widget_close_generate_meals="isGenerateMeals=!isGenerateMeals" />
    </div>
</template>

<script lang="js">
import { defineComponent } from 'vue'

import WidgetAddFood from "./WidgetAddFood.vue"
import WidgetAddRecipe from "./WidgetAddRecipe.vue"
import WidgetGenerateMeals from "./WidgetGenerateMeals.vue"
import IntakeDoughnut from "./charts/IntakeDoughnut.vue"

export default defineComponent({
    name: "Dashboard",
    components: {
        WidgetAddFood,
        WidgetAddRecipe,
        WidgetGenerateMeals,
        IntakeDoughnut,
    },
    data() {
        return {
            kcal_today: 1600,
            isAddFood: false,
            isAddRecipe: false,
            isGenerateMeals: false,
            state: {
                chartData: {
                    datasets: [
                        {
                            data: [54, 160, 79],
                            backgroundColor: ['#ef233c50', '#ffd6a550', '#1fcf8050']
                        }
                    ],
                    // These labels appear in the legend and in the tooltips when hovering different arcs
                    labels: ['Carbohydrates', 'Proteins', 'Fats']
                },
                chartOptions: {
                    responsive: true,
                    
                }
            }
        };
    },
    methods: {
        info_text() {
            return "Good start,"
        }
    },
})

</script>

<style scoped>

hr {
    margin-top: 0;
}

.dashboard-options {
    display: flex;
    flex-wrap: wrap;
    gap: 15px;

    justify-content: center;
    align-content: flex-start;
}

.dashboard-options .dashboard-option {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    height: 100px;
    width: 155px;

    border-radius: 10px;
    background: #1fcf8025;
}

.dashboard-option span {
    font-size: 18px;
}

.dashboard-option .option-icon {
    font-size: 35px;
}

/* .dashboard-option .bi {
    font-size: 30px;
} */

/* .dashboard-option .bi-dice-6 {
    transform: rotate(45deg);
} */

.intake-today {
    display: grid;
    justify-content: center;
}

.intake-today h4 > span {
    color: #1fcf80;
    font-weight: bold;
}

.nothing-there-yet-text {
  background: #fff4f4;
  border: 1px solid #cccccc;
  border-radius: 14px;
  height: 60px;
  padding: 15px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 15px;
}

.nothing-there-yet-img {
  display: flex;
  justify-content: center;
  align-items: center;
}

.nothing-there-yet-img .svg-img {
  width: 80%;
  height: auto;
}

</style>