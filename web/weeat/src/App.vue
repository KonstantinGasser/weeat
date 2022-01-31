<template>
  <router-view/>
</template>

<style>
:root {
  --box-bg: linear-gradient(26deg, rgb(255, 255, 255) 0%, rgb(249, 250, 255) 100%);
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
body {
  background: #fff4f4;
  overflow: hidden;
  height: 100%;
}
@media only screen 
and (min-device-width : 375px) 
and (max-device-width : 812px)
and (-webkit-device-pixel-ratio : 3) { 
  body {
    position: fixed;
  }
}
a {
  color: #00a2ff;
  text-decoration: none;
}
a.small {
  font-size: 12px;
}

.text-underline {
  text-decoration: underline;
}

.text-emph {
  font-weight: bold;
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
  grid-template-rows: 55px max-content max-content;
  gap: 10px;
  grid-template: repeat(100px, 100px);
  
  height: 100vh;
  width: 95vw;
  margin: 0 auto;
  padding: 15px 10px;
}
.main-frame.blur {
  filter: brightness(0.8);
  transition: brightness 300ms linear;
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
  font-weight: bolder;
}
.headline h1 > span {
  color: #1fcf80;
  font-weight: bolder;
}

.headline h1 a > span {
  color: #1fcf80;
  font-weight: bolder;
}
.headline h1 a {
  color: black;
}

.inline-name {
  font-weight: bolder;
}

.inline-name span {
  color: #1fcf80;
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
.action-btn.append {
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
  height: 100%;
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
  opacity: 1;
  position: absolute;
  overflow: hidden;
  background: #ffffff;
  
  width: 95%;
  height: max-content;
  max-height: 95vh;
  bottom: 0;
  top: 0;
  left: 0;
  right: -200%;
  margin: auto;
  padding: 0 15px 5px 15px;
  border-radius: 14px 14px 14px 14px;
  box-shadow: 0 0 6px 2px rgba(0,0,0,0.1);
  transition: 250ms ease-out;
}
.widget.widget-active {
  transition: 500ms ease-out;
  /* bottom: 0%; */
  right: 0
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
.widget .widget-menu {
  display: flex;
  justify-content: flex-start;
  gap: 15px;
  padding: 15px 15px;
}
.widget .widget-menu div {
  padding: 3px 15px;
  border-radius: 14px;
  border: 1px solid #1fcf80;
  background-color: #1fcf8025;
  color: #1fcf80;
}
.widget .widget-menu div.selected {
  background-color: #1fcf80;
  color: #ffffff;
}
.bi:hover {
  cursor: pointer;
}
.icon-medium {
    font-size: 45px;
}
.nutrition_tag {
  width: max-content;
  padding: 3px 10px;
  font-weight: bolder;
  border-radius: 14px;
}
.nutrition_tag.tag_kcal {
  background: #ef233c25;
  color: #ef233c;
  border: 1px solid #ef233c;
}
.nutrition_tag.tag_carbs {
  background: #ffadad25;
  color: #ffadad;
  border: 1px solid #ffadad;
}
.nutrition_tag.tag_fats {
  background: #ffd6a525;
  color: #ffd6a5;
  border: 1px solid #ffd6a5;
}
.nutrition_tag.tag_protein {
  background: #1fcf8025;
  color: #1fcf80;
  border: 1px solid #1fcf80;
}
.nutrition_tag.tag_sugar {
  background: #f1c0e825;
  color: #f1c0e8;
  border: 1px solid #f1c0e8;
}
.list_ingredients {
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
}
.item_ingredient {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 3px 5px;
  width: max-content;
  border-radius: 24px;
  background: #1fcf8025;
  color: #1fcf80;
  border: 1px solid #1fcf80;
}
.item_ingredient .bi {
  padding-top: 2px;
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
