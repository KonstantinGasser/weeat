<template>
    <div class="main-frame" :class="{'blur': !(!isAddFood&&!isAddRecipe&&!isGenerateMeals&&!isSearchFood)}">
        <div class="headline">
            <h1><span>we</span>Eat</h1>
        </div>
        <div v-if="has_cookie" class="intake-today">
            <h3>{{info_text()}}...<br> you have <span>1600</span> kcal today</h3>
            <IntakeDoughnut v-bind:chartData="state.chartData" v-bind:chartOptions="state.chartOptions" />
        </div>
        <div v-if="!has_cookie">
            <div class="nothing-there-yet">
                <span class="text-center">
                    You have rejected the cookie 🍪. Thereby you will be not able to track your intake for this day 🙁
                </span>
                <div>
                    <button class="action-btn" @click="enableCookie()">enable 🍪</button>
                </div> 
            </div>
        </div>
        <div class="footer">
            <div>
                <a href="/#/questions">Questions <i class="bi bi-box-arrow-up-right"></i></a>
            </div>
            <div class="github">
                <a href="https://github.com/KonstantinGasser/weeat-app" target="_blank">
                    <img src="../assets/GitHub-Mark/PNG/GitHub-Mark-32px.png" alt="">
                </a>
                <a href="https://github.com/KonstantinGasser/weeat-app" target="_blank">
                    <span>Made with ❤️</span>
                </a>
            </div>
            <div>
                <a href="/#/changelog">Changelog <i class="bi bi-box-arrow-up-right"></i></a>
            </div>
        </div> 
    </div>
    <MenuBar @clicked_item="openWidget" :clickable="!isAddFood&&!isAddRecipe&&!isGenerateMeals&&!isSearchFood"/>
    <WidgetHelloFriend :class="{'widget-active': !isHelloWorld}" @widget_close_hello_friend="cookie_check_resp"/>
    <WidgetAddFood :class="{'widget-active':isAddFood}" @widget_close_new_food="isAddFood=!isAddFood"/>
    <WidgetAddRecipe :class="{'widget-active':isAddRecipe}" @widget_close_new_recipe="isAddRecipe=!isAddRecipe"/>
    <WidgetGenerateMeals :class="{'widget-active':isGenerateMeals}" @widget_close_generate_meals="isGenerateMeals=!isGenerateMeals" />
    <WidgetSearch :class="{'widget-active':isSearchFood}" @widget_close_search_food="isSearchFood=!isSearchFood" />
</template>

<script lang="js">
import { defineComponent } from 'vue'
import { useCookies } from "vue3-cookies"

import MenuBar from "./MenuBar.vue"

import WidgetHelloFriend from "./WidgetHelloFriend.vue"
import WidgetAddFood from "./WidgetAddFood.vue"
import WidgetAddRecipe from "./WidgetAddRecipe.vue"
import WidgetGenerateMeals from "./WidgetGenerateMeals.vue"
import WidgetSearch from "./WidgetSearch.vue"

import IntakeDoughnut from "./charts/IntakeDoughnut.vue"

export default defineComponent({
    name: "Dashboard",
    components: {
        MenuBar,
        WidgetHelloFriend,
        WidgetAddFood,
        WidgetAddRecipe,
        WidgetGenerateMeals,
        WidgetSearch,
        IntakeDoughnut,
    },
    setup(){
        const { cookies } = useCookies()
        return { cookies }
    },
    mounted() {
        const is_set = this.has_cookie_set()
        this.has_cookie = is_set
        this.isHelloWorld = is_set
    },
    data() {
        return {
            has_cookie: false,
            kcal_today: 1600,
            isHelloWorld: false,
            isAddFood: false,
            isAddRecipe: false,
            isGenerateMeals: false,
            isSearchFood: false,
            panel_overview: true,
            panel_browse: false,
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
        
    },
    methods: {
        openWidget(name) {
            switch (name) {
                case "WidgetAddFood":
                    this.isAddFood=!this.isAddFood
                    break;
                case "WidgetAddRecipe":
                    this.isAddRecipe=!this.isAddRecipe
                    break;
                case "WidgetGenerateMeals":
                    this.isGenerateMeals=!this.isGenerateMeals
                    break;
                case "WidgetSearchFood":
                    this.isSearchFood=!this.isSearchFood
                    break;
            
                default:
                    break;
            }
        },
        cookie_check_resp() {
            this.has_cookie = this.has_cookie_set()
            this.isHelloWorld = !this.isHelloWorld
        },
        enableCookie() {
            this.isHelloWorld = !this.isHelloWorld
        },
        has_cookie_set(){
            return this.cookies.get("weeat_uid") !== null ? true : false
        },
        info_text() {
            return "Good start"
        }
    },
})

</script>

<style scoped>

hr {
    margin-top: 0;
}

.dashboard {
    height: max-content;
    display: grid;
    gap: 15px;
    margin-top: 15px;
}

.dashboard-options {
    display: flex;
    flex-wrap: wrap;
    row-gap: 10px;
    column-gap: 20px;

    justify-content: center;
    align-content: flex-start;
}

.dashboard-options .dashboard-option {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    height: 75px;
    width: 145px;

    border-radius: 10px;
    background: #1fcf8025;
    box-shadow: 0 0 10px 2px rgb(0 0 0 / 10%);
}

.dashboard-option span {
    font-size: 16px;
}

.dashboard-option .option-icon {
    font-size: 30px;
    height: 35px;
}

.dash-content-links {
    display: flex;
    justify-content: space-evenly;
    /* border-bottom: 1px solid #000000; */
}

.dash-content-links span {
    font-size: 16px;
    font-weight: bold;
    padding: 5px 15px;
    border-radius: 10px;
    border: 1px solid black
}
.intake-today {
    display: grid;
    justify-content: center;
}

.intake-today h3 > span {
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

.footer {
    position: absolute;
    bottom: 0;
    margin: 0 auto;
    left: 0;
    right: 0;

    padding: 0px 25px 10px 25px;
    display: flex;
    justify-content: space-evenly;
    align-items: flex-end;
}

.footer div {
    width: 30%;
    display: flex;
    justify-content: center;
}

.footer .github {
    display: grid;
}

.footer .github span {
    font-size: 12px;
    text-align: center;
}

.footer .github a {
    margin: 0 auto;
}

.footer a {
    color: #918f8f;
    font-size: 14px;
}

</style>