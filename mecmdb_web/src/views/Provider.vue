<template>
  <div class="provider">

    <div>
      <div>
        <a-button type="primary" @click="showModal" style="margin: 8px 0;">添加供应商</a-button>
      </div>
    
      <!-- <a-button type="primary" @click="showModal" style="margin: 8px 0">添加供应商</a-button> -->
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
          <!-- 展示列表 -->
          <a-form-item has-feedback label="供应商名称" name="provider_name">
            <a-input v-model:value="providerForm.form.name" type="text" autocomplete="off"/>
          </a-form-item>
          <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
<!-- 重置按钮 -->
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
          edit
        </template>
        <template v-else-if="column.dataIndex === 'operation_del'">
          <a-popconfirm
              v-if="dataSource.length"
              title="Sure to delete?"
              @confirm="onDelete(record.id)"
          >
            <a>Delete</a>
          </a-popconfirm>
        </template>
      </template>
    </a-table>
  </div>
</template>

<script setup>
import {computed, onMounted, reactive, ref} from 'vue';
import {CheckOutlined, EditOutlined} from '@ant-design/icons-vue';
import {cloneDeep} from 'lodash-es';
import http from "../utils/http"

// 供应商表格
const columns = [{
  title: '供应商名称',
  dataIndex: 'category_name',
}];


const dataSource = ref([]);
const count = computed(() => dataSource.value.length + 1);
const editableData = reactive({});
const edit = key => {
  editableData[key] = cloneDeep(dataSource.value.filter(item => key === item.key)[0]);
};
const save = key => {
  Object.assign(dataSource.value.filter(item => key === item.key)[0], editableData[key]);
  delete editableData[key];
};
const onDelete = id => {
  console.log("id", id)
  // axios请求
  http.delete("/assets/provider", {
    params: {
      "id": id
    }

  }).then((res) => {

    console.log("res", res);
    dataSource.value = dataSource.value.filter(function (item) {
      return item.id !== id
    });
    message.success('删除主机成功！')

  })


};


// 添加主机
const visible = ref(false);
const showModal = () => {
  visible.value = true;
};
const handleOk = e => {

  // 发送Ajax请求，添加主机
  // 获取主机类别

  // providerForm.form.name = parseInt(Form.form.name)

  console.log("form", providerForm.form.name)

  http.post("/assets/provider", providerForm.form).then((res) => {
    console.log("res", res);

    // 前端显示添加主机
    dataSource.value.unshift(res.data.data.name)


  })

  // 样式处理
  resetForm()
  visible.value = false;





};

const cancelForm = e => {
  resetForm()
  visible.value = false;
};
const providerForm = reactive({
  labelCol: {span: 6},
  wrapperCol: {span: 14},
  other: '',
  form: {
    name: '',
    // host_category_id: "",
    // ip_addr: '',
    // username: '',
    // port: '',
    // remark: '',
    // password: ''
  },
  rules: {
    name: [
      {required: true, message: '请输入供应商名称', trigger: 'blur'},
      {min: 3, max: 30, message: '长度在3-10位之间', trigger: 'blur'}
    ],
    // password: [
    //   {required: true, message: '请输入连接密码', trigger: 'blur'},
    //   {min: 3, max: 30, message: '长度在3-10位之间', trigger: 'blur'}
    // ],
    // host_category_id: [
    //   {required: true, message: '请选择类别', trigger: 'change'}
    // ],
    // username: [
    //   {required: true, message: '请输入用户名', trigger: 'blur'},
    //   {min: 3, max: 30, message: '长度在3-10位', trigger: 'blur'}
    // ],
    // ip_addr: [
    //   {required: true, message: '请输入连接地址', trigger: 'blur'},
    //   {max: 30, message: '长度最大15位', trigger: 'blur'}
    // ],
    // port: [
    //   {required: true, message: '请输入端口号', trigger: 'blur'},
    //   {max: 5, message: '长度最大5位', trigger: 'blur'}
    // ]
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
const providerList = reactive({
  data: []
})

onMounted(() => {
  // 获取供应商列表
  http.get("/assets/provider").then((res) => {
    console.log("res", res);

    dataSource.value = res.data.data.provider_list
  })

  // /获取主机类别
  // http.get("/host/category").then((res) => {
  //   console.log("res", res);

  //   categoryList.data = res.data.data.host_category_list
  // })


})


</script>

<style scoped>

</style>