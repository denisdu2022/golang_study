<template>
  <div>
    <ul>
<!--      循环menu_list,拿到键值-->
      <li v-for="menu in menu_list"><a :href="menu.link">{{menu.name}}</a></li>
      <li>
        &nbsp;&nbsp;&nbsp;<span>所在地:</span> <select v-model="city" name="" id="">
        <option value="北京">北京</option>
        <option value="上海">上海</option>
        <option value="深圳">深圳</option>
      </select>
      </li>
    </ul>
  </div>

</template>

<script>
export default {
  name: "Nav",
  //数据
  data(){
    return {
      //menu_list数组
      menu_list:[
        {name:"百度","link":"https://www.baidu.com"},
        {name: "腾讯", "link": "https://www.qq.com"},
        {name: "小米", "link": "https://www.xiaomi.com"},
      ],
      city:"北京",
    }
  }

}

</script>

<!--<style scoped></style>  定义样式作用域只在当前该组件下起作用,需要再标签上增加scoped-->
<style scoped>
ul, li {
  list-style: none;
  padding: 0;
  margin: 0;
}

ul::after {
  overflow: hidden;
  clear: both;
  display: block;
  content: "";
}

li {
  float: left;
  margin: 0 20px;
}

a {
  text-decoration: none;
  color: #666;
}
</style>