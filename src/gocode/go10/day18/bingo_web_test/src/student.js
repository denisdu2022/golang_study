import {reactive} from "vue";

export default reactive(
    {
        info:{

        },
        get_info(){
            this.info={
                name:"xiaoming",
                age:18,

            }
        },
        //自增
        sum_info(){
            this.info.age++
        }
    }
)