<template>
  <h1>天气预报</h1>
  <h3>输入城市名称获取七日天气</h3>
  <p><input type="text" placeholder="请输入城市"  :value="choose_city">
    <button @click="sendAjax" >获取七日天气</button>
  </p>


  <table border="1" v-if="forecast.length >1">
    <tr>
      <td>日期</td>
      <td>最高气温</td>
      <td>最低气温</td>
      <td>天气类型</td>
      <td>天气详情</td>
    </tr>
    <tr v-for="day_forecast in forecast">
      <td>{{day_forecast.day}}</td>
      <td>{{day_forecast.tem1}}</td>
      <td>{{day_forecast.tem}}</td>
      <td>{{day_forecast.wea}}</td>
      <td>{{day_forecast.narrative}}</td>
    </tr>
  </table>


</template>

<script>
//对内导入axios组件
import axios from "axios";

//对外导出应用的默认设置
export default {
  name: "Forecast",
  data(){
    return {
      forecast:[],
      //choose_city: "北京",
    }
  },
  //方法
  methods:{
    //发送Ajax请求
    sendAjax(){
      //定义this对象
      var that = this;
      //获取天气信息 发起get请求
      //axios.get("url",{params:{}})
      axios.get("https://v0.yiketianqi.com/api?unescape=1&version=v91&appid=32631524&appsecret=6gMTwOdt",{
        params:{
          city: that.choose_city,
        }
      }).then(function (response){
        //then响应
        console.log("response>>>",response)
        that.forecast = response.data.data
      })
    },
  },
  //create方法
  created(){
    //在构建数据的时候自动触发
    //在页面构建的时候加载
    this.sendAjax()
  },
  //监听choose_city的变化发送Ajax请求
  watch:{
    choose_city(newVal,oldVal){
      //choose_city发生变化就发送Ajax请求
      this.sendAjax()
    },
  },
  //在子组件中接收父组件传来的数据
  //使用props
  props:{
    choose_city:{
      default:"北京", //默认值
      type:String,
    },
  },
}
</script>

<style scoped>

</style>