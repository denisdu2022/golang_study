<template>
  <p>组合API的基本使用</p>
  <p>显示对象格式的响应数据</p>
  <p>{{user}}</p>

  <p @click="change_user_info">点击修改{{user.name}}的数据</p>
  <input type="text" v-model="user.age">
  <p>显示普通数据</p>
  <p>{{data}}</p>
  <input type="text" v-model="data.age">
  <hr>
  <p>显示非对象的基本类型数据</p>
  <p @click="msg_add">{{msg}}</p>
  <input type="text" v-model="msg">
  <p>{{new_msg}}</p>
</template>

<script>
import { reactive,ref, onMounted, watch, computed } from 'vue'
export default {
  name: "Group",
  // setup 会自动执行，并且会在befroreCreate钩子方法执行之前就自动执行了
  setup(){
    console.log(this); // undefined，setup执行的时候，连this都没有定义，所以，我们不可能在setup里面调用之前的选项API方法了。
    console.log("setup自动执行了");

    // 显示一个变量(对象)到页面中
    let data = {name:"xiaoming",age:17,sex:true,}
    // 一旦把数据转换成响应式数据，则vm对象就会自动的维护和保持数据的全局一致性, 也就是HTML页面中修改了数据，vm对象中的setup里面的数据也被发生改动
    let user = reactive(data); // // 把数据通过reactive函数进行处理的目的就是为了转换数据为响应式数据

    let change_user_info = ()=>{
      user.age = 30;
    }

    // 显示一个变量(基本类型：布尔值，数值，字符串)到页面中
    let message = 100
    let msg = ref(message);  // ref的内部本质就是右边代码的简写 let msg = reactive({value: message});
    // 如果希望数据提供给视图模板HTML中使用return通过对象进行返回

    console.log(msg.value);  // setup代码中如果要调用经过ref处理过的数据，需要value属性来操作
    msg.value=2000;          // 修改也是通过value属性来修改原来的数据

    // 操作数据的方法
    let msg_add = ()=>{
      msg.value++;
    }

    // 钩子方法的调用[组合API中不再提供created和beforeCreated，原来这2个钩子方法中代码直接写在setup中即可]
    onMounted(() => {
      console.log('Component is mounted!')
    });

    // 监听数据
    // watch("响应数据的变量名", (new_value, old_value)=>{
    //
    // });
    watch(msg, ()=>{
      if(msg.value == "python"){
        console.log("message的值不能是python了");
        msg.value = 2000;
      }
    });

    // 计算属性
    let new_msg = computed(()=>{
      return msg.value * 4;
    });

    return {
      user, // user:user的简写
      data, // 非响应数据
      msg,
      msg_add, // 对外提供操作数据的方法
      change_user_info,
      new_msg,
    }
  },
  // beforeCreate() {
  //   console.log("钩子beforeCreate执行了");
  // }
}
</script>

<style scoped>

</style>