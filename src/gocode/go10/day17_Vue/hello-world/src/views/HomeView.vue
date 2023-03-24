<template>
  <div class="home">
<!--    <img alt="Vue logo" src="../assets/logo.png">-->
<!--    <HelloWorld msg="Welcome to Your Vue.js App"/>-->
<!--  使用Nav子组件-->
    <Nav></Nav>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
//引入Nav子组件
import Nav from "@/components/Nav.vue";

export default {
  name: 'HomeView',
  components: {
    //加载Nav子组件
    Nav
    //HelloWorld
  }
}
</script>
