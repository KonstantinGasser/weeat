<template>
    <div class="main-frame">
        <div class="headline">
            <h1>Hello <span>Friend</span></h1>
        </div>
        <div class="action-row">
            <div class="btns">
                <button class="action-btn medium" @click="isAddFood=!isAddFood">Add Food</button>
                <button class="action-btn medium" @click="isAddRecipe=!isAddRecipe">Add Recipe</button>
            </div>
        </div>
        <div class="some-info-will-go-here">
            <div class="nothing-there-yet-text">
                <span>
                    There seems to be no data..go and do something
                </span>
            </div>
            <div class="nothing-there-yet-img">
                <img class="svg-img" src="../assets/svg/basketball.svg" alt="">
            </div>
        </div>
        <WidgetAddFood :class="{'widget-active':isAddFood}" @widget_close_new_food="isAddFood=!isAddFood"/>
        <WidgetAddRecipe :class="{'widget-active':isAddRecipe}" @widget_close_new_recipe="isAddRecipe=!isAddRecipe"/>
    </div>
</template>

<script lang="js">
import WidgetAddFood from "./WidgetAddFood.vue"
import WidgetAddRecipe from "./WidgetAddRecipe.vue"

export default {
    name: "Dashboard",
    components: {
        WidgetAddFood,
        WidgetAddRecipe,
    },
    data() {
        return {
            isAddFood: false,
            isAddRecipe: false,
        };
    },
    methods: {
        addFruit(){
            this.food_name = "hello";
        },
    }
}

</script>

<style scoped>

.nothing-there-yet-text {
  background: #fff4f4;
  border: 1px solid #cccccc;
  border-radius: 14px;
  height: 60px;
  padding: 15px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.nothing-there-yet-img {
  display: grid;
  align-items: center;
}

.nothing-there-yet-img .svg-img {
  width: 100%;
  height: auto;
}

</style>