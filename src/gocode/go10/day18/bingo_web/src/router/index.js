import { createRouter, createWebHistory } from 'vue-router'

import ShowCenter from "../views/ShowCenter.vue";
//导入GroupAPI
import GroupAPI from "../views/GroupAPI.vue";
//导入login组件
import Login from "../views/Login.vue";
//导入base组件
import Base from "../views/Base.vue";

const routes = [
  {
    path: '/',
    name: 'ShowCenter',
    component: ShowCenter
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/base',
    name: 'Base',
    component: Base
  },
  {
    path: '/groupAPI',
    name: 'GroupAPI',
    component: GroupAPI
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
