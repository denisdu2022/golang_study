<template>
<h1>主机管理组件</h1>

  <button @click="sexAjax">点击按钮发送Ajax请求</button>

</template>

<script>
//导入axios模块
import axios from "axios";
import http from "@/utils/http";

export default {
  name: "Host",
    methods:{
      "sexAjax":function (){
          http.get("host",{
              "headers":{
                 "Authorization":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImY2NGUyZTJjLTk0NWUtNDQwZi04NGE3LTM1MmQ0YmNjOTZjOSIsInVzZXJuYW1lIjoiand0Iiwibmlja25hbWUiOiIiLCJhdmF0YXIiOiIiLCJleHAiOjE2ODU1MjQ4MTIsImlhdCI6MTY4NTUyMzAxMiwibmJmIjoxNjg1NTIzMDEyfQ.wcRXSrf4GJk753WBiDi9ywRCmLTrdjiHv3hQl58dBbE",
              }
              },).then(response => {
                  console.log(">>>",response)
              })
      }
    }
}
</script>

<style scoped>

</style>