<template>
    <a-layout style="min-height: 100vh">
        <a-layout-sider v-model:collapsed="collapsed" collapsible>
            <div class="logo"
                 style="font-style: italic;text-align: center;font-size: 20px;color:#fff;margin: 10px 0;line-height: 50px;font-family: 'Times New Roman'">
                <span><a-switch v-model:checked="checked"/> DevOps</span>
            </div>
            <a-menu v-for="menu in menu_list" v-model:selectedKeys="selectedKeys" theme="dark" mode="inline">


                <a-menu-item v-if="menu.children.length===0" key="{{menu.id}}">
                    <MyIcon :icon="menu.icon"></MyIcon>
                    <span>
            <router-link :to="menu.menu_url">
             {{ menu.title }}
            </router-link>
          </span>
                </a-menu-item>

                <a-sub-menu v-else :key="menu.id">
                    <template #title>
            <span>
               <MyIcon :icon="menu.icon"></MyIcon>
              <span>{{ menu.title }}</span>
            </span>
                    </template>

                    <a-menu-item v-for="child_menu in menu.children" :key="child_menu.id">
                        <router-link :to="child_menu.menu_url">{{ child_menu.title }}</router-link>
                    </a-menu-item>

                </a-sub-menu>

            </a-menu>
        </a-layout-sider>


        <a-layout>
            <a-layout-header style="background: #fff; padding: 20px">

                <a-row type="flex" justify="start">

                    <a-col :span="6">
                        <a-breadcrumb>
                            <a-breadcrumb-item href="">
                                <home-outlined/>
                            </a-breadcrumb-item>
                            <a-breadcrumb-item href="">
                                <user-outlined/>
                                <span>Application List</span>
                            </a-breadcrumb-item>
                            <a-breadcrumb-item>Application</a-breadcrumb-item>
                        </a-breadcrumb>
                    </a-col>

                    <a-col :span="1" :offset="17">
                        <a-breadcrumb>

                            <a-popconfirm
                                    title="您确认要注销登陆吗?"
                                    ok-text="Yes"
                                    cancel-text="No"
                                    @confirm="logout"
                                    @cancel="cancel"
                            >
                                <a-button type="primary" class="logout">
                                    注销
                                </a-button>
                            </a-popconfirm>

                        </a-breadcrumb>
                    </a-col>

                </a-row>

            </a-layout-header>
            <a-layout-content style="margin: 0 16px">
                <a-breadcrumb style="margin: 16px 0">
                    <a-breadcrumb-item>User</a-breadcrumb-item>
                    <a-breadcrumb-item>Bill</a-breadcrumb-item>
                </a-breadcrumb>
                <div :style="{ padding: '24px', background: '#fff', minHeight: '860px' }">
                    <router-view/>
                </div>
            </a-layout-content>
            <a-layout-footer style="text-align: center">
                Ant Design ©2018 Created by Ant UED
            </a-layout-footer>
        </a-layout>
    </a-layout>
</template>


<script setup>
import {reactive, ref} from 'vue';
import storage from "@/utils/storage";
import MyIcon from "@/components/MyIcon";
import router from '@/router/index.js'

const collapsed = ref(false)
const selectedKeys = ref(['1'])
const checked = ref(true)

const menu_list = reactive([
    {
        id: 1, icon: 'settings', title: '系统管理', tube: '', 'menu_url': '/bingo/show_center', children: []
    },
    {
        id: 2, icon: 'desktop', title: '资产管理', 'menu_url': '/bingo/host', children: []
    },
    {
        "id": 3, icon: 'cloud', title: '批量任务', tube: '', menu_url: '/bingo/workbench', children: [
            {id: 10, icon: 'mail', title: '执行任务', 'menu_url': '/bingo/multi_exec'},
            {id: 11, icon: 'mail', title: '命令管理', 'menu_url': '/bingo/template_manage'},
        ]
    },
    {
        id: 4, icon: 'git', title: '代码发布', tube: '', menu_url: '/bingo/workbench', children: [
            {id: 12, title: '应用管理', menu_url: '/bingo/release'},
            {id: 13, title: '发布申请', menu_url: '/bingo/release'}
        ]
    },
    {id: 5, icon: 'history', title: '定时计划', tube: '', menu_url: '/bingo/workbench', children: []},

    {id: 7, icon: 'monitor', title: '监控预警', tube: '', 'menu_url': '/bingo/workbench', children: []},

    {
        id: 9, icon: 'user', title: '用户中心', tube: '', menu_url: '/bingo/workbench', children: [
            {id: 20, title: '账户管理', tube: '', menu_url: '/bingo/workbench'},
            {id: 21, title: '角色管理', tube: '', menu_url: '/bingo/workbench'},
            {id: 22, title: '系统设置', tube: '', menu_url: '/bingo/workbench'}
        ]
    }
])


const logout = () => {
    storage.clearStorage()
    router.push("/login")
}

</script>
<style>
#components-layout-demo-side .logo {
    height: 32px;
    margin: 16px;
    background: rgba(255, 255, 255, 0.3);
}

.site-layout .site-layout-background {
    background: #fff;
}

[data-theme='dark'] .site-layout .site-layout-background {
    background: #141414;
}

.ant-menu .ant-menu-item a {
    color: rgba(255, 255, 255, 0.65);
}

.ant-menu .ant-menu-item-selected a, .ant-menu .ant-menu-item-selected a:hover {
    color: white;
}


</style>