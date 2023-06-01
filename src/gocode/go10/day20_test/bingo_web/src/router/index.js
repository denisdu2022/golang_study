import {createRouter, createWebHistory} from 'vue-router'
import Login from '../views/Login.vue'
import Base from '../views/Base'
import ShowCenter from '../views/ShowCenter'
import Host from '../views/Host'
import storage from "@/utils/storage";

const routes = [
    {
        meta: {
            title: 'bingo自动化运维平台'
        },
        path: '/bingo',
        alias: '/', // 给当前路径起一个别名
        name: 'Base',
        component: Base, // 快捷键：Alt+Enter快速导包,
        children: [
            {
                path: 'show_center',
                alias: '', // 给当前路径起一个别名
                name: 'ShowCenter',
                component: ShowCenter,
                meta: {
                    title: "展示中心",
                    authorization: true
                },
            },
            {
                path: 'host',
                name: 'Host',
                component: Host,
                meta: {
                    title: "主机管理",
                    authorization: true
                },
            },
        ],
    },
    {
        meta: {
            title: '账户登陆',
            authorization: false
        },
        path: '/login',
        name: 'Login',
        component: Login // 快捷键：Alt+Enter快速导包，
    },
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})


// 导航守卫
router.beforeEach((to, from, next) => {
    document.title = to.meta.title
    // 登录状态验证
    console.log("to meta", to.meta)
    if (to.meta.authorization && !storage.getUserInfo()) {
        next({"name": "Login"})
    } else {
        next()
    }
})

// 暴露路由对象
export default router