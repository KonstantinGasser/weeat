<template>
  <div class="widget">
    <div class="widget-header">
        <span>New Recipe</span>
        <i class="bi bi-x icon-medium" @click="$emit('closeWidget')"></i>
    </div>
    <div class="">
      
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