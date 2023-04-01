//在组合API里声明一个基本数据类型的变量,需要先从vue引入ref,在声明
import {ref} from "vue";

let age = ref(0)
//声明函数 模拟Ajax请求 >从服务端获取的
let ajax_get_age= (num)=>{
    age.value = num
}
//让年龄做自加
let change_age = ()=>{
    age.value++
}

//对外导出
export default {
    age,
    ajax_get_age,
    change_age,
}