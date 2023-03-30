<template>
  <h1>ShowCenter</h1>
  <p>ant-design的按钮</p>
  <div>
  <p>
    <a-tag color="pink">pink</a-tag>
    <a-tag color="red">red</a-tag>
    <a-tag color="orange">orange</a-tag>
    <a-tag color="green">green</a-tag>
    <a-tag color="cyan">cyan</a-tag>
    <a-tag color="blue">blue</a-tag>
    <a-tag color="purple">purple</a-tag>
  </p>
  <p>
    <a-tag color="#f50">#f50</a-tag>
    <a-tag color="#2db7f5">#2db7f5</a-tag>
    <a-tag color="#87d068">#87d068</a-tag>
    <a-tag color="#108ee9">#108ee9</a-tag>
  </p>
  </div>
  <hr/>
  <p>使用Echarts的饼图</p>
  <div >
    <div  ref="chart"> </div>
  </div>


</template>

<script >
//导入
import {onMounted,ref} from "vue";
//导入echarts
import  * as echarts from 'echarts';
//echarts饼图
//定义响应式变量
const  chart = ref(); //const chart = reactive({value:""})
//设置饼图
let setBingtu = () => {
  //弹窗测试
  alert(123)
}

onMounted(() =>{
  setBingtu()
})

export default {
  name: 'ShowCenter',

}

</script>

<!--style 设置了scoped属性只针对当前组件生效-->
<style scoped>
table, tr, th, td {
  border: 1px solid red;
  border-collapse: collapse; /* 合并边框 */
}

th, td {
  width: 200px;
  text-align: center; /* 文本水平居中 */
  height: 30px;
  line-height: 30px;
}
</style>