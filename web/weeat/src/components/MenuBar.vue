<template>
  <div class="menu-bar">
    <div class="menu-item" @click="$emit('clicked_item','WidgetAddFood')">
      <!-- <i class="bi bi-plus-circle"></i> -->
      <span class="item-icon">🫐</span>
    </div>
    <div class="menu-item" @click="$emit('clicked_item','WidgetAddRecipe')">
      <!-- <i class="bi bi-plus-circle"></i> -->
      <span class="item-icon">🌮</span>
    </div>
    <div class="menu-item" @click="$emit('clicked_item','WidgetGenerateMeals')">
      <!-- <span>Feeling lucky?</span> -->
      <span class="item-icon">😬</span>
    </div>
    <div class="menu-item" @click="$emit('clicked_item','WidgetSearchFood')">
      <!-- <span>Search</span> -->
      <span class="item-icon">🔎</span>
    </div>
  </div>
</template>

<script lang="js">

export default {
    name: "MenuBar",
    components: {
  
    },
    data() {
        return {
            
        };
    },
    computed: {
        
    },
    methods: {
    },
};

</script>

<style scoped>
.menu-bar {
    position: absolute;
    bottom: 25px;
    left: 25px;
    right: 25px;
    margin: 0 auto;

    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 35px;
    padding: 3px 15px;
    width: max-content;
    
    background: #ffffff;
    border-radius: 25px;
    height: 50px;
    box-shadow: 0 0 6px 2px rgba(0, 0, 0, 0.1);
}

.menu-item {
    display: grid;
}
.menu-item span {
    font-size: 30px;
}
</style>