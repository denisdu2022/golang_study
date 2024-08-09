import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
//引入ant-design-vue
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
//引入echarts
let echarts = require('echarts')
//引入xterm
import 'xterm/css/xterm.css'
import 'xterm/lib/xterm'


const app=createApp(App)

app.use(router).use(Antd).mount('#app')

//配置全局的echarts
app.config.globalProperties.$echarts = echarts;
