<template>
<h1>ShowCenter</h1>
  <h2>Echarts饼图</h2>
  <div>
    <div class="echarts">
      <div class="myChart" ref="chart"></div>
    </div>
  </div>
  <hr/>
  <h2>日历</h2>
  <div :style="{ width: '300px', border: '1px solid #d9d9d9', borderRadius: '4px' }">
    <a-calendar v-model:value="value" :fullscreen="false" @panelChange="onPanelChange" />
  </div>
  <hr/>
  <h2>模态对话框</h2>
  <div>
    <a-button type="primary" @click="showModal">Open Modal</a-button>
    <a-modal v-model:visible="visible" title="Basic Modal" @ok="handleOk">
      <a-form :model="formState" :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-item label="Activity name">
          <a-input v-model:value="formState.name" />
        </a-form-item>
        <a-form-item label="Instant delivery">
          <a-switch v-model:checked="formState.delivery" />
        </a-form-item>
        <a-form-item label="Activity type">
          <a-checkbox-group v-model:value="formState.type">
            <a-checkbox value="1" name="type">Online</a-checkbox>
            <a-checkbox value="2" name="type">Promotion</a-checkbox>
            <a-checkbox value="3" name="type">Offline</a-checkbox>
          </a-checkbox-group>
        </a-form-item>
        <a-form-item label="Resources">
          <a-radio-group v-model:value="formState.resource">
            <a-radio value="1">Sponsor</a-radio>
            <a-radio value="2">Venue</a-radio>
          </a-radio-group>
        </a-form-item>
        <a-form-item label="Activity form">
          <a-input v-model:value="formState.desc" type="textarea" />
        </a-form-item>
        <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
          <a-button type="primary" @click="onSubmit">Create</a-button>
          <a-button style="margin-left: 10px">Cancel</a-button>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
  <p></p>
  <hr/>




</template>

<!--setup简写-->
<script setup>
import {  ref , toRaw , reactive , onMounted} from 'vue';
//导入echarts
import * as echarts from 'echarts';

    //日历组件
    const value = ref();
    const onPanelChange = (value, mode) => {
      console.log(value, mode);
    };


    //对话框
    const visible = ref(false);
    const showModal = () => {
      visible.value = true;
    };
    const handleOk = e => {
      console.log(e);
      visible.value = false;
    };

    //form表单
    const formState = reactive({
      name: '',
      delivery: false,
      type: [],
      resource: '',
      desc: '',
    });
    const onSubmit = () => {
      console.log('submit!', toRaw(formState));
    };

    //echarts饼图
    const chart = ref() //相当于 const chart = reactive({value:""})
    //设置饼图的方法
    let setEcharts = ()=> {
      //alert(123)
      let myChart = echarts.init(chart.value);
      let option;

      option = {
        legend: {
          top: 'bottom'
        },
        toolbox: {
          show: true,
          feature: {
            mark: { show: true },
            dataView: { show: true, readOnly: false },
            restore: { show: true },
            saveAsImage: { show: true }
          }
        },
        series: [
          {
            name: 'Nightingale Chart',
            type: 'pie',
            radius: [50, 250],
            center: ['50%', '50%'],
            roseType: 'area',
            itemStyle: {
              borderRadius: 8
            },
            data: [
              { value: 40, name: '北京' },
              { value: 38, name: '上海' },
              { value: 32, name: '广州' },
              { value: 30, name: '深圳' },
              { value: 28, name: '珠海' },
              { value: 26, name: '大理' },
              { value: 22, name: '广东' },
              { value: 18, name: '陕西' }
            ]
          }
        ]
      };

      option && myChart.setOption(option);

    }
    //onMounted
    onMounted(()=>{
      //页面构建时执行setEcharts方法
      setEcharts()
    })




</script>

<!--scoped只对当前组件生效样式-->
<style scoped>
  .myChart{
    height: 550px;
  }
</style>