<template>
    <div class="main-frame">
        <div class="headline">
            <h1><span>we</span>Eat</h1>
        </div>
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
            <div v-if="cookie_set" class="intake-today mt-2">
                <hr>
                <h4>{{info_text()}}<br> you have <span>1600</span> kcal today</h4>
                <IntakeDoughnut v-bind:chartData="state.chartData" v-bind:chartOptions="state.chartOptions" />
            </div>
            <div v-if="!cookie_set">
                <div class="nothing-there-yet">
                    <span>
                        You have rejected the cookie 🍪. Thereby you will be able to track your intake for this day 🙁.
                    </span>
                    <div>
                        <button class="action-btn" @click="enableCookie()">enable 🍪</button>
                    </div> 
                </div>
            </div>
            
        </div>
        <WidgetHelloFriend :class="{'widget-active': isHelloWorld}" @widget_close_hello_friend="cookie_check_resp"/>
        <WidgetAddFood :class="{'widget-active':isAddFood}" @widget_close_new_food="isAddFood=!isAddFood"/>
        <WidgetAddRecipe :class="{'widget-active':isAddRecipe}" @widget_close_new_recipe="isAddRecipe=!isAddRecipe"/>
        <WidgetGenerateMeals :class="{'widget-active':isGenerateMeals}" @widget_close_generate_meals="isGenerateMeals=!isGenerateMeals" />
    </div>
</template>

<script lang="js">
import { defineComponent } from 'vue'
import { useCookies } from "vue3-cookies"

import WidgetHelloFriend from "./WidgetHelloFriend.vue"
import WidgetAddFood from "./WidgetAddFood.vue"
import WidgetAddRecipe from "./WidgetAddRecipe.vue"
import WidgetGenerateMeals from "./WidgetGenerateMeals.vue"

import IntakeDoughnut from "./charts/IntakeDoughnut.vue"

export default defineComponent({
    name: "Dashboard",
    components: {
        WidgetHelloFriend,
        WidgetAddFood,
        WidgetAddRecipe,
        WidgetGenerateMeals,
        IntakeDoughnut,
    },
    setup(){
        const { cookies } = useCookies()
        return { cookies }
    },
    mounted() {
        this.isHelloWorld = this.cookies.get("weeat_uid") === null ? true : false
    },
    data() {
        return {
            isHelloWorld: false,
            cookie_set: false,
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
    computed: {
        has_cookie(){
            return this.cookies.get("weeat_uid") !== null ? true : false
        },
    },
    methods: {
        cookie_check_resp(resp) {
            this.cookie_set = resp
            this.isHelloWorld = !this.isHelloWorld
        },
        enableCookie() {
            this.isHelloWorld = !this.isHelloWorld
        },
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
    gap: 20px;

    justify-content: center;
    align-content: flex-start;
}

.dashboard-options .dashboard-option {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    height: 80px;
    width: 140px;

    border-radius: 10px;
    background: #1fcf8025;
    box-shadow: 0 0 10px 2px rgb(0 0 0 / 10%);
}

.dashboard-option span {
    font-size: 18px;
}

.dashboard-option .option-icon {
    font-size: 35px;
}

.intake-today {
    display: grid;
    justify-content: center;
}

.intake-today h4 > span {
    color: #1fcf80;
    font-weight: bold;
}

.nothing-there-yet {
    margin: 25px 0;
    background: #fff4f4;
    border: 1px solid #cccccc;
    border-radius: 14px;
    height: max-content;
    padding: 15px;
    display: grid;
    gap: 15px;
    align-items: center;
    justify-content: center;
}

.nothing-there-yet div {
  width: max-content;
}

</style>