<template>
  <div class="multi_exec">
    <div>
      <h3>执行主机：</h3>
      <div>
        <a-tag closable @close="delete_select_host(info_index)" v-for="(info,info_index) in show_host_info.data"
               :key="info.id">
          {{ `${info.name}(${info.ip_addr}:${info.port})` }}
        </a-tag>
      </div>
    </div>
    <div style="margin-top: 10px;">
      <a-button @click="showModal">
        <plus-outlined/>
        从主机列表中选择
      </a-button>
      <a-modal :visible="visible" title="选择执行主机" @ok="handleOk" @cancel="hideModel" width="960px">
        <div>
          <a-row>
            <a-col :span="8">
              <a-form-item label="主机类别：" :label-col="formItemLayout.labelCol" :wrapper-col="formItemLayout.wrapperCol">
                <a-select style="width: 160px;" placeholder="请选择" v-model:value="host_form.form.category"
                          @change="has_change_category">
                  <a-select-option :value="value.id" v-for="(value, index) in category_list.data" :key="value.id">
                    {{ value.name }}
                  </a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :span="8">
              <a-form-item :label-col="formItemLayout.labelCol" :wrapper-col="formItemLayout.wrapperCol" label="主机别名：">
                <a-input placeholder="请输入"/>
              </a-form-item>
            </a-col>
            <a-col :span="4">
              <a-form-item :label-col="formItemLayout.labelCol" :wrapper-col="formItemLayout.wrapperCol" label="已选：">
                <span style="margin-left: 8px">
									{{ `${selectedRowKeys.length}` }}台主机
                </span>
              </a-form-item>
            </a-col>
            <a-col :span="4">
              <a-button type="primary" style="margin-top: 3px;" @click="refresh_data">
                <sync-outlined></sync-outlined>
                刷新
              </a-button>
            </a-col>
          </a-row>
        </div>
        <div>
          <a-table
              :columns="columns"
              :data-source="host_list.data"
              :pagination="false"
              :rowKey="record => record.id"
              :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
          ></a-table>
        </div>
      </a-modal>
    </div>


    <div class="cmd">
      <div>执行指令：</div>
      <v-ace-editor v-model:value="content" lang="sh" theme="chrome"
                    style="height: 300px; font-size: 20px;margin-bottom: 1rem;"></v-ace-editor>
    </div>
    <div style="margin-top: 20px; margin-bottom: 20px;">
      <a-button @click="showModal2">
        <plus-outlined/>
        从执行模板中选择
      </a-button>
      <a-modal :visible="visible2" title="选择执行模板" @ok="handleOk2" width="960px" @cancel="hideModel2">
        <div>
          <a-row>
            <a-col :span="10">
              <a-form-item label="模板类别：" :label-col="formItemLayout.labelCol" :wrapper-col="formItemLayout.wrapperCol">
                <a-select style="width: 160px;" placeholder="请选择" @change="handleSelectChange2"
                          v-model="template_form.form.category">
                  <a-select-option :value="value.id" v-for="(value) in template_category_list.data" :key="value.id">
                    {{ value.name }}
                  </a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :span="10">
              <a-form-item :label-col="formItemLayout.labelCol" :wrapper-col="formItemLayout.wrapperCol" label="模板名称：">
                <a-input placeholder="请输入" v-model:value="template_form.form.name"/>
              </a-form-item>
            </a-col>
            <a-col :span="4">
              <a-button type="primary" style="margin-top: 3px;" @click="refresh_data2">
                <sync-outlined></sync-outlined>
                刷新
              </a-button>
            </a-col>
          </a-row>
        </div>
        <div>
          <a-table
              :columns="tem_columns"
              :data-source="template_list.data"
              :rowKey="record => record.id"
              :row-selection="{ radioselectedRow: radioselectedRow, onChange: onSelectChange2,type: 'radio' }"
          >
          </a-table>
        </div>
      </a-modal>
    </div>
    <div>
      <a-button type="primary" @click="execute_cmd">
        <thunderbolt-outlined></thunderbolt-outlined>
        开始执行
      </a-button>
    </div>
  </div>
</template>

<script setup>
import {computed, reactive, ref} from "vue";
import {PlusOutlined, SyncOutlined,ThunderboltOutlined} from '@ant-design/icons-vue';
import http from "@/utils/http";


// vue3-ace-editor
import {VAceEditor} from 'vue3-ace-editor';
import 'ace-builds/src-noconflict/theme-chrome';
import 'ace-builds/src-noconflict/ext-language_tools';
import 'ace-builds/src-noconflict/mode-sh';
import {message} from "ant-design-vue";



const execute_cmd = () => {
  // 远程执行命令
  console.log(content.value)
  console.log(selected_host_ids)
  // 执行指令操作
  http.post(`/host/cmd`, {
    host_ids: selected_host_ids.value,
    cmd: content.value
  }).then((res) => {
    message.success("远程主机已经执行命令...")
  }).catch((error => {
    message.error("远程主机执行命令失败！")
  }))
}

// 表单分栏布局长度，弹窗的首行表单配置信息
const formItemLayout = reactive({
  labelCol: {span: 8},
  wrapperCol: {span: 14}
})
// 表单的字段设置，弹窗的表格的每一列数据的配置信息
const columns = ref([
  {
    slots: {title: 'customTitle'},
    scopedSlots: {customRender: 'action'}
  }, {
    title: '类别',
    dataIndex: 'category_name',
    key: 'category_name'
  }, {
    title: '主机名称',
    dataIndex: 'name',
    key: 'name'
  }, {
    title: '连接地址',
    dataIndex: 'ip_addr',
    key: 'ip_addr',
    width: 200
  }, {
    title: '端口',
    dataIndex: 'port',
    key: 'port'
  }, {
    title: '备注信息',
    dataIndex: 'description',
    key: 'description'
  }
])

// 显示选中的所有主机内容
const show_host_info = reactive({
  data: []
})
const visible = ref(false) // 是否显示主机列表的弹窗

// 主机表单信息
const host_form = reactive({
  form: {
    category: undefined // 当前选择的主机分类ID
  }
})

// 当前显示咋表格中的主机列表数据
const host_list = reactive({
  data: []
})

// 主机分类列表
const category_list = reactive({
  data: []
})

// 已经勾选的主机ID列表
const selectedRowKeys = ref([])
// 选中的主机id列表
const selected_host_ids = ref([])

// 计算属性
const hasSelected = computed(() => {
  return selectedRowKeys.length > 0
})

// 显示弹窗
const showModal = () => {
  visible.value = true
}

// 关闭弹窗
const hideModel = () => {
  visible.value = false;
}

// 选中主机时触发的，selectedRowKeys被选中的主机id列表
const onSelectChange = (selectedKeys) => {
  selectedRowKeys.value = selectedKeys
}

const get_host_category_list = () => {
  // 获取主机类别
  http.get(`/host/category`).then((response) => {
    category_list.data = response.data.data
  })
}
get_host_category_list()


const get_host_list = (category = null) => {
  // 获取主机列表
  let params = {}
  if (category !== null) {
    params.host_category_id = category
  }
  http.get(`/host`, {
    params: params,
  }).then((response) => {
    host_list.data = response.data.data.host_list
  })
}

get_host_list()

const has_change_category = (category) => {
  // 切换主机分类时，重新获取主机列表
  get_host_list(category)
}

const refresh_data = () => {
  // 刷新页面
  // location.reload()
  // 刷新数据
  get_host_list()
  host_form.form.category = undefined
  selectedRowKeys.value = []
}

const handleOk = () => {
  selected_host_ids.value = []
  show_host_info.data = []
  host_list.data.forEach((v, k) => {
    if (selectedRowKeys.value.includes(v.id)) { // 判断某元素是否在数组中用includes比较合适，不能用in
      show_host_info.data.push({
        id: v.id,
        name: v.name,
        ip_addr: v.ip_addr,
        port: v.port
      })
      selected_host_ids.value.push(v.id)
    }
  })
  // 关闭弹窗
  visible.value = false
}

const delete_select_host = (infoIndex) => {
  // 移除已经勾选的主机信息
  show_host_info.data.splice(infoIndex, 1)
  let idsList = selected_host_ids.value.splice(infoIndex, 1)
  let idIndex = selectedRowKeys.value.indexOf(idsList[0])
  selectedRowKeys.value.splice(idIndex, 1)
}

// 用户输入的模板命令
const content = ref("")



// 是否显示选择指令模板的窗口
const visible2 = ref(false);
// 搜索指令模板的数据的表单
const template_form = reactive({
  form: {
    category: undefined,
    name: "",
  }
})

// 指令模板分类列表
const template_category_list = reactive({
  data: [],
})

const tem_columns = ref([
  {
    title: '模板类型',
    dataIndex: 'category_name',
    key: 'category_name'
  },
  {
    title: '模板名称',
    dataIndex: 'name',
    key: 'name'

  },
  {
    title: '模板内容',
    dataIndex: 'content',
    key: 'content',
    width: 200
  }
])

// 指令模板列表
const template_list = reactive({
  data: [],
})

// 选项的指令模板的列表
const radioselectedRow = ref([]);

// 显示选择指令模板的弹窗
const showModal2 = () => {
  visible2.value = true
}

// 隐藏选择指令模板的弹窗
const hideModel2 = () => {
  visible2.value = false;
}

// 当选择指令模板时
const handleOk2 = (e) => {
  let tid = radioselectedRow.value[0] // 选中的模板id值
  // 通过模板id值，找到该模板记录中的cmd值，并赋值给content属性
  template_list.data.forEach((v, k) => {
    if (v.id === tid) {
      // 把用户选择的指令模板内容显示到编辑器中
      content.value = v.content
    }
  })
  visible2.value = false
}

const onSelectChange2 = (selectedRow) => {
  // [6, 7, 8, 9]
  console.log('>>>>> ', selectedRow)
  radioselectedRow.value = selectedRow
}

const handleSelectChange2 = (value) => {
  // 切换模板分类
  get_templates_list(value)
}

const refresh_data2 = () => {
  get_templates_list();
}

// 获取指令模板分类列表
const get_templates_category_list = () => {
  http.get(`/cmd/category`)
      .then(response => {
        template_category_list.data = response.data.data.cmd_category_list
      })
}

get_templates_category_list()

// 获取指令模板列表
const get_templates_list = (category = null) => {
  let params = {}
  if (category !== null) {
    params.command_category_id = category
  }
  http.get(`/cmd`, {
    params: params,
  }).then(response => {
    template_list.data = response.data.data.cmd_list
  })
}

get_templates_list()

</script>

<style scoped>
.multi_exec {
  text-align: left;
}
</style>