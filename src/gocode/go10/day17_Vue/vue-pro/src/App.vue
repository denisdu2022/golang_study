<template>
<!--  <nav>-->
<!--    <router-link to="/">Home</router-link> |-->
<!--    <router-link to="/about">About</router-link> |-->
<!--&lt;!&ndash;    下边普通的a标签会跳转&ndash;&gt;-->
<!--    <a href="https://baidu.com">百度</a>-->

<!--  </nav>-->

<!--&lt;!&ndash;  根据不同的路由显示不同的组件&ndash;&gt;-->
<!--  <router-view/>-->

<!--  <h1>hello vue</h1>-->
<!--  &lt;!&ndash;  根据不同的路由显示不同的组件&ndash;&gt;-->
  <router-view/>


</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
</style>
