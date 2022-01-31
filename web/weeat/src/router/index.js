import { createRouter, createWebHashHistory } from 'vue-router'
import Dashboard from "../components/Dashboard.vue"


const routes = [
  {
    path: '/',
    name: 'Dashboard',
    component: Dashboard
  },
  {
    path: '/questions',
    name: 'QandA',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../components/QuestionsAndAnswers.vue')
  },
  {
    path: '/changelog',
    name: 'Changelog',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../components/Changelog.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
