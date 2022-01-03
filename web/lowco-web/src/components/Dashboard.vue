<template>
    <div class="main-frame">
        <div class="headline">
            <h1>Hello <span>Friend</span></h1>
        </div>
        <div class="action-row">
            <div class="btns">
                <button class="action-btn medium" @click="isAddFood=!isAddFood">Add Food</button>
                <button class="action-btn medium" @click="isAddFood=!isAddFood">Add Recipe</button>
            </div>
        </div>
        <div class="some-info-will-go-here">
            <div class="nothing-there-yet">
                <img class="svg-img" src="../assets/svg/basketball.svg" alt="">
            </div>
        </div>
        <WidgetAddFood :class="{'widget-active':isAddFood}" @closeWidget="isAddFood=!isAddFood"/>
        <WidgetAddRecipe :class="{'widget-active':isAddFood}" @closeWidget="isAddFood=!isAddFood"/>
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

</style>