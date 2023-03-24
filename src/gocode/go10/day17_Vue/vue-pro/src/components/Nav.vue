<template>
  <div>
    <ul>
      <li v-for="menu in menu_list"><a :href="menu.link">{{ menu.name }}</a></li>
      <li>
        &nbsp;&nbsp;&nbsp; <span>所在地：</span><select v-model="city" name="" id="">
        <option value="北京">北京</option>
        <option value="上海">上海</option>
        <option value="深圳">深圳</option>
      </select>
      </li>
    </ul>

    <p>下拉菜单选中的城市:{{city}}</p>

  </div>
</template>

<script>
export default {
  name: "Nav",
  data() {
    return {
      menu_list: [
        {name: "百度", "link": "http://www.baidu.com"},
        {name: "腾讯", "link": "http://www.qq.com"},
        {name: "小米", "link": "http://www.xiaomi.com"},
      ],
      city:"北京",
    }
  },
  //create()自动触发的方法
  // created() {
  //   //子组件传递给父组件数据使用this.$emit
  //   this.$emit("getCity",this.city)
  // },
  //侦听属性,侦听city的改变
  watch:{
    //监听city事件
      city(newVal,oldVal){
        //在浏览器控制台打印新值和旧值
        console.log(newVal,oldVal),
        //把新值传递给父组件
        this.$emit("getCity",newVal)
      }
  },

}
</script>

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