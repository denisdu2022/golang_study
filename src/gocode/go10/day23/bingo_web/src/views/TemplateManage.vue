<template>
  <div class="template_manager">
    <div class="template-top">
      <a-row>
        <a-col :span="5">
          <a-form-item label="模板类别：" :label-col="formItemLayout.labelCol" :wrapper-col="formItemLayout.wrapperCol">
            <a-select style="width: 160px;" placeholder="请选择" @change="handleSelectChange"
                      v-model:value="search.category_id">
              <a-select-option :value="value.id" v-for="(value) in template_category_list.data" :key="value.id">
                {{ value.name }}
              </a-select-option>
            </a-select>
          </a-form-item>
        </a-col>
        <a-col :span="10">
          <a-form-item :label-col="formItemLayout.labelCol" :wrapper-col="formItemLayout.wrapperCol" label="模板名称：">
            <a-input v-model:value="search.name" placeholder="请输入" @change="handleInputChange"/>
          </a-form-item>
        </a-col>
        <a-col :span="2">
          <a-button type="primary" style="margin-top: 3px;" @click="refresh_data">
            <sync-outlined></sync-outlined>
            刷新
          </a-button>
        </a-col>
      </a-row>
    </div>
    <div style="margin-top: 20px;">
      <a-button type="primary" @click="showModal">
        <plus-outlined></plus-outlined>
        新建
      </a-button>
      <!-- 新建指令模板的弹窗 -->
      <a-modal :visible="visible" title="新建模板" ok-text="添加" cancel-text="取消" @ok="handleSubmit" @cancel="hideModel"
               width="960px">
        <a-form-item label="模板类别">
          <a-row>
            <a-col :span="6">
              <a-select placeholder="" v-model:value="template_form.category_id">
                <a-select-option :value="value.id" v-for="(value) in template_category_list.data" :key="value.id">
                  {{ value.name }}
                </a-select-option>
              </a-select>
            </a-col>
            <a-col :span="6">
              <a-button type="link" @click="showTemplateCategoryFormModal">添加类别</a-button>
            </a-col>
          </a-row>
        </a-form-item>
        <a-form-item label="模板名称">
          <a-input v-model:value="template_form.name"/>
        </a-form-item>
        <a-form-item label="模板内容">
          <v-ace-editor v-model:value="template_form.content" lang="sh" theme="chrome"
                        style="height: 300px; font-size: 20px;margin-bottom: 1rem;"></v-ace-editor>
        </a-form-item>
      </a-modal>
    </div>
    <!-- 新建模板分类弹窗 -->
    <a-modal
        :width="600"
        title="新建模板类别"
        :visible="templateCategoryFromVisible"
        @cancel="templateCategoryFormCancel"
        ok-text="提交"
        cancel-text="取消"
        @ok="onTemplateCategoryFromSubmit"
    >
      <a-form-model ref="hostCategoryRuleForm" :model="templateCategoryForm.form" :rules="templateCategoryForm.rules"
                    :label-col="templateCategoryForm.labelCol" :wrapper-col="templateCategoryForm.wrapperCol">
        <a-form-model-item ref="name" label="类别名称" prop="name">
          <a-row>
            <a-col :span="24">
              <a-input placeholder="请输入模板类别名称" v-model:value="templateCategoryForm.form.name"/>
            </a-col>
          </a-row>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <div style="margin-top: 10px;">
      <a-table :columns="columns" :data-source="templates_list.data" :rowKey="record => record.id">
        <template #bodyCell="{ column, text, record }">
          <template v-if="column.dataIndex === 'action'">
            <a-popconfirm v-if="templates_list.data.length" title="您确认要删除当前主机信息吗?">
              <a href="javascript:;">删除</a> |
            </a-popconfirm>
            <a href="javascript:;">编辑</a>
          </template>
        </template>
      </a-table>
    </div>
  </div>
</template>

<script setup>
import {reactive, ref} from "vue";
import axios from "@/utils/http";
import {message} from "ant-design-vue";
import {PlusOutlined, SyncOutlined} from '@ant-design/icons-vue';
import {VAceEditor} from 'vue3-ace-editor';
import 'ace-builds/src-noconflict/theme-chrome';
import 'ace-builds/src-noconflict/ext-language_tools';
import 'ace-builds/src-noconflict/mode-sh';

const formItemLayout = reactive({
  labelCol: {span: 6},
  wrapperCol: {span: 14}
})

const columns = ref([
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
    sorter: {
      compare: (a, b) => a.id - b.id,
      multiple: 3,
    },
  },
  {
    title: '模板类型',
    dataIndex: 'category_name',
    key: 'category_name'
  },
  {
    title: '模板名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '模板内容',
    dataIndex: 'content',
    key: 'content',
    width: 200
  },
  {
    title: '操作',
    key: 'action',
    dataIndex: 'action',
    width: 200,
    scopedSlots: {customRender: 'action'}
  }
])

// 是否显示弹窗
const visible = ref(false)

// 指令模板分类列表
const template_category_list = reactive({
  data: []
})

// 指令模板列表
const templates_list = reactive({
  data: []
})

// 添加指令模板时的cmd指令
const content = ref('')

const template_form = reactive({
  category_id: undefined,
  name: '',
  content: ''
})

const search = reactive({
  category: null,
  name: "",
})


const showModal = () => {
  visible.value = true
}

const hideModel = () => {
  visible.value = false;
}

const refresh_data = () => {
  search.category = null
  search.name = "";
  get_templates_list(search.category)
}

const handleSelectChange = (value) => {
  // 切换模板分类
  search.name = "";
  get_templates_list(value)
}

const handleInputChange = () => {
  search.category_id = undefined
  console.log(search.name)
  get_templates_list(null, search.name)
}

// 获取指令模板列表
const get_templates_list = (category = null, template_name = null) => {
  // 查询参数
  let params = {}
  if (category !== null) {
    // 如果有指定模板分类ID，则根据模板分类ID查询模板列表
    params.command_category_id = category
  }
  if (template_name !== null) {
    // 如果有执行模板名称，则根据模板名称查询模板列表
    params.name = template_name
  }
  axios.get(`/cmd`, {
    params: params,
  }).then(response => {
    templates_list.data = response.data.data.cmd_list
  })
}

get_templates_list()


// 获取模板指令分类列表
const get_templates_category_list = () => {
  axios.get(`/cmd/category`).then(response => {
    template_category_list.data = response.data.data.cmd_category_list
  })
}
get_templates_category_list()

const handleSubmit = () => { // handleSubmit必须声明并使用这个方法，名字不能改
  let params = {}
  params.name = template_form.name
  params.command_category_id = template_form.category_id
  params.content = template_form.content
  // 拿到验证成功之后的数据
  // 发送添加模板的请求
  axios.post(`/cmd`, params)
      .then(response => {
        console.log(response)
        visible.value = false
        get_templates_list()
      })
}

// 是否显示添加主机类别的弹窗
const templateCategoryFromVisible = ref(false)

// 添加模板类别需要的数据属性
const templateCategoryForm = reactive({
  labelCol: {span: 6},
  wrapperCol: {span: 14},
  other: '',
  form: {
    name: ''
  },
  rules: {
    name: [
      {required: true, message: '请输入类别名称', trigger: 'blur'},
      {min: 3, max: 10, message: '长度在3-10位之间', trigger: 'blur'}
    ]
  }
})

const showTemplateCategoryFormModal = () => {
  // 显示新建主机类别的表单窗口
  templateCategoryFromVisible.value = true
}

const hostCategoryFormCancel = () => {
  // 取消并关闭添加主机类别的表单窗口
  resetTemplateCategoryForm() // 清空表单内容
  templateCategoryFromVisible.value = false // 关闭对话框
}

const resetTemplateCategoryForm = () => {
  // 重置添加主机类别的表单信息
  templateCategoryForm.form.name = "";
}

const templateCategoryFormCancel = () => {
  // 取消并关闭添加主机类别的表单窗口
  resetTemplateCategoryForm() // 清空表单内容
  templateCategoryFromVisible.value = false // 关闭对话框
}

const onTemplateCategoryFromSubmit = () => {
  // 添加主机类别的表单提交处理
  axios.post("cmd/category", templateCategoryForm.form).then(response => {
    if (response.data.code === 0) {
      message.success({
        content: "创建模板类别成功!",
        duration: 2,
      }).then(() => {
        template_category_list.data.push(response.data.data.cmdCategory)
        // 清空表单数据，并关闭添加主机分类表单窗口
        hostCategoryFormCancel()
      })
    }
  })
}

</script>

<style scoped>
.template-top {
  margin-top: 15px;
  margin-bottom: 15px;
}

.template_manager {
  text-align: left;
}
</style>
