import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
//引入ant-design-vue
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
//引入settings
import './settings';
import settings from "@/settings";

//import echarts form 'echarts'
let echarts = require('echarts')

//设置app常量
const app = createApp(App)

//app.use加载route和ant-design-vue
app.use(router).use(Antd).mount('#app')
//配置全局settings
app.config.globalProperties.$settings = settings;
//配置全局echarts
app.config.globalProperties.$echarts = echarts;