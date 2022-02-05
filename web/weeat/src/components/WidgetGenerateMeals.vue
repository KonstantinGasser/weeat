<template>
  <div class="widget">
    <div class="widget-header">
        <span>We make your meals!</span>
        <i class="bi bi-x icon-medium" @click="closeWidget()"></i>
    </div>
    <div class="widget-body">
      Mhm sorry.. this feature is still under development 😅 - stay up to date and have a look at the <a href="/#/changelog">changelog</a>
    </div>
  </div>
</template>


<script>

export default {
  name: "WidgetGenerateMeals",
  setup() {
  },
  data() {
    return {
        emit_widget_name: 'widget_close_generate_meals',
        food_name: null,
        food_cat: 1,
        food_kcal: null,
        food_carbs: null,
        foot_protein: null,
        food_fats: null,
        isGenerating: false,
    };
  },
  beforeUnmounted() {
    console.log("destroying: generate.Meals")
    this.isGenerating = !this.isGenerating
  },
  methods: {
      closeWidget() {
        this.$emit(this.emit_widget_name)
        this.isGenerating=false
      },
  },
};
</script>


<style scoped>
.generate-meals {
  height: 200px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.generate-meals .bi {
  font-size: 75px;
  color: #1fcf80;
  padding: 25px;
  width: 150px;
  height: 150px;
  border-radius: 50%;
  /* background: radial-gradient(#1fcf8055,#1fcf8015, #fff); */
  transform: scale(1);
  animation: pulse 3s infinite;
}


@keyframes pulse {
  0% {
		transform: scale(0.95);
		background: #1fcf8055;
	}

	70% {
		transform: scale(1);
		box-shadow: 0 0 0 10px rgba(0, 0, 0, 0);
    background: #1fcf8015;
	}

	100% {
		transform: scale(0.95);
    background: #ffffff;
	}
}
</style>