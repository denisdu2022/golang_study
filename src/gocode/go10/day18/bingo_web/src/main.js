import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
//引入ant-design-vue
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
//引入echarts
let echarts = require('echarts')

const app=createApp(App)

app.use(router).use(Antd).mount('#app')

//配置全局的echarts
app.config.globalProperties.$echarts = echarts;
