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
}

body {
  background: #fff4f4;
  position: relative;
  overflow: hidden;
  /* height: 100%; */
}

a {
  color: #00a2ff;
  text-decoration: none;
}
a.small {
  font-size: 12px;
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
  gap: 10px;
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
    align-items: center;
    gap: 15px;
  }

.action-row button {
  margin: 0 0;
}

.action-btn {
  padding: 3px 15px;

  width: 150px;

  font-size: 16px;
  font-weight: 600;

  border: 2px solid #1fcf80;
  border-radius: 10px;
  color: #1fcf80;
  background: #ffffff;
}

.action-btn.form-btn {
  height: 100%;
  width: min-content;
}

.action-btn.form-btn.btn-end {
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
}

.action-btn:hover {
  color: #ffffff;
  background: #1fcf80;
  transition: 200ms;
}

.btn-choice {
  padding: 3px 15px;

  width: 150px;

  font-size: 16px;
  font-weight: 600;

  background: #ffffff;
}

.btn-choice.no {
  border: 2px solid #ef233c;
  border-radius: 10px;
  color: #ef233c;
}

.btn-choice.no:hover {
  color: #ffffff;
  background: #ef233c;
  transition: 200ms;
}

.form-grid {
  padding: 0 15px;
}

.widget {
  position: absolute;
  overflow: hidden;
  background: #ffffff;
  
  width: 95%;
  height: max-content;

  bottom: -100%;
  left: 0;
  right: 0;
  margin-left: auto; 
  margin-right: auto;

  padding: 0 15px 25px 15px;
  border-radius: 14px 14px 0 0;
  box-shadow: 0 0 6px 2px rgba(0,0,0,0.1);
  transition: 250ms ease-out;
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
  line-height: 1;
}

.widget .widget-header span > small {
  font-size: 16px;
}

.bi:hover {
  cursor: pointer;
}

.icon-medium {
    font-size: 45px;
}


.loader,
.loader:before,
.loader:after {
  border-radius: 50%;
  width: 2.5em;
  height: 2.5em;
  -webkit-animation-fill-mode: both;
  animation-fill-mode: both;
  -webkit-animation: load7 1.8s infinite ease-in-out;
  animation: load7 1.8s infinite ease-in-out;
}
.loader {
  color: #1fcf80;
  font-size: 10px;
  margin: 80px auto;
  position: relative;
  text-indent: -9999em;
  -webkit-transform: translateZ(0);
  -ms-transform: translateZ(0);
  transform: translateZ(0);
  -webkit-animation-delay: -0.16s;
  animation-delay: -0.16s;
}
.loader:before,
.loader:after {
  content: '';
  position: absolute;
  top: 0;
}
.loader:before {
  left: -3.5em;
  -webkit-animation-delay: -0.32s;
  animation-delay: -0.32s;
}
.loader:after {
  left: 3.5em;
}
@-webkit-keyframes load7 {
  0%,
  80%,
  100% {
    box-shadow: 0 2.5em 0 -1.3em;
  }
  40% {
    box-shadow: 0 2.5em 0 0;
  }
}
@keyframes load7 {
  0%,
  80%,
  100% {
    box-shadow: 0 2.5em 0 -1.3em;
  }
  40% {
    box-shadow: 0 2.5em 0 0;
  }
}

</style>
