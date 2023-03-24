<template>
  <div class="home">
<!--    <img alt="Vue logo" src="../assets/logo.png">-->
<!--    <HelloWorld msg="Welcome to Your Vue.js App"/>-->

<!--    导航组件-->
<!--    绑定触发子组件的事件 让绑定的事件等于传递过来的getCity-->
    <Nav @getCity="getCity"></Nav>
<!--    天气预报组件,发送Ajax请求-->

<!--    父组件给Forecast子组件传值-->
    <Forecast :choose_city="city"/>

    <p>Nav组件选中的城市:{{city}}</p>
  </div>
</template>

<script>
// @ is an alias to /src
//import HelloWorld from '@/components/HelloWorld.vue'
//引入Nav组件
import Nav from '@/components/Nav.vue'
//引入天气预报组件
import Forecast from "@/components/Forecast.vue";

export default {
  name: 'HomeView',
  //让Nav组件中选择的城市传递到Home组件
  data(){
    return {
      city:"",
    }
  },
  components: {
    //HelloWorld
    //引入Nav子组件
    Nav,
    //引入天气预报子组件
    Forecast,
  },
  //在父组件home组件中定义methods方法来接收子组件传来的数据
  methods:{
    getCity(choose_city){
      //让home的city等于子组件传过来的city
      this.city = choose_city
    }
  }
}
</script>
