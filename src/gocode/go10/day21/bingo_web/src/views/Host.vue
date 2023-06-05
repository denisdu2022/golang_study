<template>
    <div class="host">

        <div>
            <a-button type="primary" @click="showModal" style="margin: 8px 0">添加主机</a-button>
            <a-modal width="860px" v-model:visible="visible" title="Basic Modal" @ok="handleOk" @cancel="cancelForm">

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
                                        v-model:value="hostForm.form.category"
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
                                <a-input placeholder="端口号" addon-before="-p" v-model:value="hostForm.form.port" type="text"
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

        <a-table bordered :data-source="dataSource" :columns="columns">
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
                <template v-else-if="column.dataIndex === 'operation_edit'">

                    <a>edit</a>

                </template>
                <template v-else-if="column.dataIndex === 'operation_del'">
                    <a-popconfirm
                            v-if="dataSource.length"
                            title="是否要删除?"
                            @confirm="onDelete(record.key)"
                    >
                        <a>Delete</a>
                    </a-popconfirm>
                </template>
            </template>
        </a-table>
    </div>
</template>

<script setup>
import {computed, reactive, ref} from 'vue';
import {CheckOutlined, EditOutlined} from '@ant-design/icons-vue';
import {cloneDeep} from 'lodash-es';


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
}, {
    title: '编辑',
    dataIndex: 'operation_edit',
}, {
    title: '删除',
    dataIndex: 'operation_del',
}];
const dataSource = ref([
    {
        'id': 1,
        'category_name': '数据库服务器',
        'name': 'izbp13e05jqwodd605vm3gz',
        'ip_addr': '47.58.131.12',
        'port': 22,
        'remark': ''
    },
    {
        'id': 2,
        'category_name': '数据库服务器',
        'name': 'iZbp1a3jw4l12ho53ivhkkZ',
        'ip_addr': '12.18.125.22',
        'port': 22,
        'remark': ''
    },
    {
        'id': 3,
        'category_name': '缓存服务器',
        'name': 'iZbp1b1xqfqw257gs563k2iZ',
        'ip_addr': '12.19.135.130',
        'port': 22,
        'remark': ''
    },
    {
        'id': 4,
        'category_name': '缓存服务器',
        'name': 'iZbp1b1jw4l01ho53muhkkZ',
        'ip_addr': '47.98.101.89',
        'port': 22,
        'remark': ''
    }
]);
const count = computed(() => dataSource.value.length + 1);
const editableData = reactive({});
const edit = key => {
    editableData[key] = cloneDeep(dataSource.value.filter(item => key === item.key)[0]);
};
const save = key => {
    Object.assign(dataSource.value.filter(item => key === item.key)[0], editableData[key]);
    delete editableData[key];
};
const onDelete = key => {
    // dataSource.value = dataSource.value.filter(item => item.key !== key);
    //发送Ajax请求删除
};

// 添加主机
const visible = ref(false);
const showModal = () => {
    visible.value = true;
};
const handleOk = e => {
    resetForm()
    visible.value = false;
};

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
        category: "",
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
        category: [
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
const categoryList = reactive({
    data: []
})
</script>

<style scoped>

</style>