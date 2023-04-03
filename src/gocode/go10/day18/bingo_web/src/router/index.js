import { createRouter, createWebHistory } from 'vue-router'

import ShowCenter from "../views/ShowCenter.vue";

//导入login组件
import Login from "../views/Login.vue";
//导入base组件
import Base from "../views/Base.vue";
//导入host组件
import Host from "@/views/Host.vue";

const routes = [
  // {
  //   path: '/bingo',
  //   name: 'bingo',
  //   component: Base
  // },
  // {
  //   path: '/',
  //   name: 'ShowCenter',
  //   component: ShowCenter
  // },
  // {
  //   path: '/host',
  //   name: 'Host',
  //   component: Host
  // },

    //bingo的功能路由
  {
      path:'/bingo',
      alias:'/', //给当前路径设置一个别名
      name:'Base',
      component: Base, //可以使用快捷键:Atl_Enter快速导包
      //子路由
      children:[
        {
          path:'show_center',
          alias:'', //给当前路径设置一个别名 alias:'',  等于alias:'/',
          name:'ShowCenter',
          component: ShowCenter,
        },
        {
          path:'host',
          //alias:'',
          name:'Host',
          component: Host,
        },
      ]
  },
    //bingo的login
    {
        meta:{
            title:'账户登录'
        },
        //alias: '',
        path: '/login',
        name: 'Login',
        component: Login
    }



]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
