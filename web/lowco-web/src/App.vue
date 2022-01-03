<template>
  <Dashboard />
</template>

<script>
import Dashboard from './components/Dashboard.vue'

export default {
  name: 'App',
  components: {
   Dashboard 
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

body {
  background: #fff4f4;
  overflow: hidden;
}
.justify-center {
  justify-content: center;
}

.justify-end {
  justify-content: flex-end;
}

.main-frame {
  position: relative;

  background: #fff4f4;
  display: grid;
  grid-template-rows: 50px 50px auto;
  gap: 25px;
  grid-template: repeat(100px, 100px);
  
  height: 100vh;
  width: 65vw;

  margin: 0 auto;
  padding: 15px 10px;
}

@media only screen 
and (min-device-width : 375px) 
and (max-device-width : 812px)
and (-webkit-device-pixel-ratio : 3) { 
  .main-frame {
    width: 100vw;
  }
}

.action-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.headline {
  display: flex;
  justify-content: center;
}

.headline h1 {
  font-size: 2.5rem;
  text-align: left;
}

.headline h1 > span {
  color: #1fcf80;
  font-weight: bolder;
}

.action-row {
  display: flex;
  justify-content: center;
}

.action-row button {
  margin: 0 10px;
}

.action-btn {
  padding: 3px 15px;

  width: max-content;

  font-size: 16px;
  font-weight: 600;

  border: 2px solid #1fcf80;
  border-radius: 10px;
  color: #1fcf80;
  background: #ffffff;
}

.action-btn:hover {
  color: #ffffff;
  background: #1fcf80;
  transition: 200ms;
}

.form-grid {
  padding: 0 15px;
}

.widget {
  transition: 250ms ease-out;
  position: absolute;
  background: #ffffff;
  
  width: 90%;
  height: 75%;

  bottom: -75%;
  left: 0;
  right: 0;
  margin-left: auto; 
  margin-right: auto;

  padding: 0 15px;
  border-radius: 14px 14px 0 0;
  box-shadow: 0 0 6px 2px rgba(0,0,0,0.1);
}

.widget.widget-active {
  transition: 500ms ease-out;
  bottom: 0;
}

.widget .widget-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.widget .widget-header span {
  font-size: 24px;
}

.icon-medium {
  font-size: 45px;
}

.nothing-there-yet {
  display: grid;
  align-items: center;
}

.nothing-there-yet .svg-img {
  width: 100%;
  height: auto;
}


</style>
