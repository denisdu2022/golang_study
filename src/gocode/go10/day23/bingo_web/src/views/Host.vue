<template>
    <div class="host">

        <div>
            <a-button type="primary" @click="showModal" style="margin: 8px 0">添加主机</a-button>
            <a-modal width="860px" v-model:visible="visible" title="添加主机" @ok="onHostFormSubmit" @cancel="cancelForm">

                <a-form
                    ref="formRef"
                    name="custom-validation"
                    :model="hostForm.form"
                    :rules="hostForm.rules"
                    v-bind="layout"
                    @finish="handleFinish"
                    @validate="handleValidate"
                    @finishFailed="handleFinishFailed"
                >
                    <a-form-item label="主机类别" prop="zone" name="category">
                        <a-row>
                            <a-col :span="12">
                                <a-select
                                    ref="select"
                                    v-model:value="hostForm.form.host_category_id"
                                    @change="handleCategorySelectChange"
                                >
                                    <a-select-option :value="category.id" v-for="category in categoryList.data" :key="category.id">
                                        {{ category.name }}
                                    </a-select-option>
                                </a-select>
                            </a-col>

                        </a-row>

                    </a-form-item>

                    <a-form-item has-feedback label="主机名称" name="name">
                        <a-input v-model:value="hostForm.form.name" type="text" autocomplete="off"/>
                    </a-form-item>


                    <a-form-item label="连接地址" name="username">

                        <a-row>
                            <a-col :span="8">
                                <a-input placeholder="用户名" addon-before="ssh" v-model:value="hostForm.form.username" type="text"
                                         autocomplete="off"/>
                            </a-col>
                            <a-col :span="8">
                                <a-input placeholder="ip地址" addon-before="@" v-model:value="hostForm.form.ip_addr" type="text"
                                         autocomplete="off"/>
                            </a-col>
                            <a-col :span="8">
                                <a-input placeholder="端口号" addon-before="-p" v-model:value="hostForm.form.port" type="number"
                                         autocomplete="off"/>
                            </a-col>
                        </a-row>
                    </a-form-item>

                    <a-form-item has-feedback label="连接密码" name="password">
                        <a-input v-model:value="hostForm.form.password" type="password" autocomplete="off"/>
                    </a-form-item>

                    <a-form-item has-feedback label="备注信息" name="remark">
                        <a-textarea placeholder="请输入主机备注信息" v-model:value="hostForm.form.remark" type="text"
                                    :auto-size="{ minRows: 3, maxRows: 5 }" autocomplete="off"/>
                    </a-form-item>


                    <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
                        <a-button @click="resetForm">Reset</a-button>
                    </a-form-item>
                </a-form>

            </a-modal>
        </div>

        <a-table bordered :data-source="hostList.data" :columns="columns">
            <template #bodyCell="{ column, text, record }">
                <template v-if="column.dataIndex === 'name'">
                    <div class="editable-cell">
                        <div v-if="editableData[record.key]" class="editable-cell-input-wrapper">
                            <a-input v-model:value="editableData[record.key].name" @pressEnter="save(record.key)"/>
                            <check-outlined class="editable-cell-icon-check" @click="save(record.key)"/>
                        </div>
                        <div v-else class="editable-cell-text-wrapper">
                            {{ text || ' ' }}
                            <edit-outlined class="editable-cell-icon" @click="edit(record.key)"/>
                        </div>
                    </div>
                </template>
                <template v-if="column.dataIndex === 'console'">
                    <router-link :to="`/bingo/host/console/${record.id}`">远程连接</router-link>
                </template>
                <template v-else-if="column.dataIndex === 'operation_edit'">

                    <a>
                        <edit-outlined/>
                    </a>

                </template>
                <template v-else-if="column.dataIndex === 'operation_del'">
                    <a-popconfirm
                        v-if="hostList.data.length"
                        title="Sure to delete?"
                        @confirm="onDelete(record)"
                    >
                        <a>
                            <delete-outlined/>
                        </a>
                    </a-popconfirm>
                </template>
            </template>
        </a-table>
    </div>
</template>

<script setup>
import {computed, onMounted, reactive, ref} from 'vue';
import {CheckOutlined, DeleteOutlined, EditOutlined} from '@ant-design/icons-vue';
import {cloneDeep} from 'lodash-es';
import http from "@/utils/http";
import {message} from 'ant-design-vue';

// 主机表格
const columns = [{
    title: '类别',
    dataIndex: 'category_name',
}, {
    title: '主机名称',
    dataIndex: 'name',
}, {
    title: 'IP地址',
    dataIndex: 'ip_addr',
}, {
    title: '端口',
    dataIndex: 'port',
}, {
    title: '备注',
    dataIndex: 'remark',
},{
    title: 'console',
    dataIndex: 'console',
}, {
    title: '编辑',
    dataIndex: 'operation_edit',
}, {
    title: '删除',
    dataIndex: 'operation_del',
}];
const hostList = reactive({data: []})
;
const count = computed(() => hostList.data.length + 1);
const editableData = reactive({});
const edit = key => {
    editableData[key] = cloneDeep(hostList.data.filter(item => key === item.key)[0]);
};
const save = key => {
    Object.assign(hostList.data.filter(item => key === item.key)[0], editableData[key]);
    delete editableData[key];
};
const onDelete = record => {

    console.log("delete id", record)
    http.delete(`host`, {
        params: {
            "id": record.id
        }
    }).then((response) => {

        hostList.data = hostList.data.filter(item => item.id !== record.id);
        message.success('删除主机成功！')

    }).catch((err) => {
        message.error('删除主机失败！')
    });


};

// 添加主机
const visible = ref(false);
const showModal = () => {
    visible.value = true;
};

const onHostFormSubmit = () => {

    // 将数据提交到后台进行保存，但是先进行连接校验，验证没有问题，再保存
    hostForm.form.port = parseInt(hostForm.form.port)

    const formData = new FormData();
    for (let attr in hostForm.form) {
        formData.append(attr, hostForm.form[attr])
    }

    http.post(`host`, formData).then((response) => {
        console.log("response.data.data.host:::", response.data.data.host)
        hostList.data.unshift(response.data.data.host)
        message.success('成功添加主机信息！', 3)
        // 清空
        resetForm()
        visible.value = false;


    }).catch((err) => {
        // 清空
        resetForm()
        visible.value = false;
        console.log("err:::", err)
        //message.error('添加主机失败!', 1.5)
        message.error('添加主机失败! '+err.response.data.message, 3)
    });
}

const cancelForm = e => {
    resetForm()
    visible.value = false;
};
const hostForm = reactive({
    labelCol: {span: 6},
    wrapperCol: {span: 14},
    other: '',
    form: {
        name: '',
        host_category_id: "",
        ip_addr: '',
        username: '',
        port: '',
        remark: '',
        password: ''
    },
    rules: {
        name: [
            {required: true, message: '请输入主机名称', trigger: 'blur'},
            {min: 3, max: 30, message: '长度在3-10位之间', trigger: 'blur'}
        ],
        password: [
            {required: true, message: '请输入连接密码', trigger: 'blur'},
            {min: 3, max: 30, message: '长度在3-10位之间', trigger: 'blur'}
        ],
        host_category_id: [
            {required: true, message: '请选择类别', trigger: 'change'}
        ],
        username: [
            {required: true, message: '请输入用户名', trigger: 'blur'},
            {min: 3, max: 30, message: '长度在3-10位', trigger: 'blur'}
        ],
        ip_addr: [
            {required: true, message: '请输入连接地址', trigger: 'blur'},
            {max: 30, message: '长度最大15位', trigger: 'blur'}
        ],
        port: [
            {required: true, message: '请输入端口号', trigger: 'blur'},
            {max: 5, message: '长度最大5位', trigger: 'blur'}
        ]
    }
});
const formRef = ref();
const layout = {
    labelCol: {
        span: 4,
    },
    wrapperCol: {
        span: 24,
    },
};
const handleFinish = values => {
    console.log(values, hostForm);
};

const handleFinishFailed = errors => {
    console.log(errors);
};

const resetForm = () => {
    formRef.value.resetFields();
};

const handleValidate = (...args) => {
    console.log(args);
};

const handleCategorySelectChange = (value) => {
    // 切换主机类别的回调处理
    console.log(value)
}

const categoryList = reactive({
    data: []
})

// 查看主机列表
const getHostList = () => {
    http.get("host").then((res) => {
        console.log("host_list:::", res)
        hostList.data = res.data.data.host_list
    })
}
// 查看主机类别列表
const getHostCategory = () => {
    http.get("host/category").then((res) => {
        console.log("host_category_list:::", res.data.data)
        categoryList.data = res.data.data
    })
}

//Console



onMounted(() => {
    getHostList()
    getHostCategory()
})


</script>

<style scoped>


</style>