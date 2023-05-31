<template>
<h1>主机管理组件</h1>

  <button @click="sexAjax">点击按钮发送Ajax请求</button>

</template>

<script>
//导入axios模块
import axios from "axios";

export default {
  name: "Host",
    methods:{
      "sexAjax":function (){
          axios.get("http://127.0.0.1:8080/api/host",
              {params:{
                  }}
          ).then(response => {
                  console.log(">>>",response)
              })
      }
    }
}
</script>

<style scoped>

</style>