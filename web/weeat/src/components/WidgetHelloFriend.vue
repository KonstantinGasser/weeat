<template>
  <div class="widget">
    <div class="widget-header">
        <span>Hello there!</span>
        <i class="bi bi-x icon-medium" @click="closeWidget()"></i>
    </div>
    <div class="widget-body">
      <div class="cookie-agree-text">
        <span><span>we</span>Eat</span> would like to set a cookie 🍪 
        <br>If you chose to accept you can track your daily intake. 
        But you dont have to! Without a 🍪 you can still <i>add foods</i>, <i>share recipes</i> and <i>generate meals</i> 
        <br>...so no preasure the choice is yours 😉 
      </div>
      <div class="cookie-actions">
        <div class="btn-row">
          <button class="action-btn yes" @click="acceptCookie()">Accept</button>
          <button class="btn-choice no" @click="rejectCookie()">Reject</button>
        </div>
        <div class="link-row">
          <a class="small" href="">more information about privacy and cookies <i class="bi bi-info-circle"></i></a>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import { useCookies } from "vue3-cookies"

export default {
  name: "WidgetHelloFriend",
  data() {
    return {
        emit_widget_name: "widget_close_hello_friend",
    };
  },
  setup(){
      const { cookies } = useCookies()
      return { cookies }
  },
  methods: {
      closeWidget() {
        this.$emit(this.emit_widget_name, false)
      },
      rejectCookie(){
        this.cookies.remove("weeat_uid")
        this.$emit(this.emit_widget_name, false)
      },
      acceptCookie(){
        this.cookies.set("weeat_uid", "test")
        this.$emit(this.emit_widget_name, true)
      },
  },
};
</script>

<style scoped>

.cookie-agree-text {
  margin-bottom: 20px;
}
.cookie-agree-text span {
  font-weight: bold;
}
.cookie-agree-text span span {
  color: #1fcf80;
}

.cookie-actions .btn-row {
  display: flex;
  justify-content: space-between;
}
</style>