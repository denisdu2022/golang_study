import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
//引入ant-design-vue
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
//引入settings
import './settings';
import settings from "@/settings";
//设置app常量
const app = createApp(App)

//app.use加载route和ant-design-vue
app.use(Antd).use(router).mount('#app')
//配置全局settings
app.config.globalProperties.$settings = settings;