<template>
<h1>ShowCenter</h1>
  <div :style="{ width: '300px', border: '1px solid #d9d9d9', borderRadius: '4px' }">
    <a-calendar v-model:value="value" :fullscreen="false" @panelChange="onPanelChange" />
  </div>

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
</template>

<script>
import { defineComponent, ref , toRaw , reactive} from 'vue';

export default defineComponent({
  setup() {
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

    return {
      value,
      onPanelChange,
      visible,
      showModal,
      handleOk,
      labelCol: {
        style: {
          width: '150px',
        },
      },
      wrapperCol: {
        span: 14,
      },
      formState,
      onSubmit,
    };
  },
});

</script>

<!--scoped只对当前组件生效样式-->
<style scoped>

</style>