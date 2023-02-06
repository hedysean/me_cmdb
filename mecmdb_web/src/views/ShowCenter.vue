<template>
  <h1>ShowCenter</h1>


  <div>
    <div class="c1">
      <div :style="{ width: '400px', border: '1px solid #d9d9d9', borderRadius: '4px' }">
        <a-calendar v-model:value="value" :fullscreen="false" @panelChange="onPanelChange"/>
      </div>
    </div>    <div class="c2">
      <!--  饼图-->
      <div class="mychart" ref="chart"></div>
    </div>

    <div class="c3">
      <div>
        <a-button type="primary" @click="showModal">Open Modal</a-button>
        <a-modal v-model:visible="visible" title="Basic Modal" @ok="handleOk">
          <a-form
              :model="formState"
              name="basic"
              :label-col="{ span: 8 }"
              :wrapper-col="{ span: 16 }"
              autocomplete="off"
              @finish="onFinish"
              @finishFailed="onFinishFailed"
          >
            <a-form-item
                label="Username"
                name="username"
                :rules="[{ required: true, message: 'Please input your username!' }]"
            >
              <a-input v-model:value="formState.username"/>
            </a-form-item>

            <a-form-item
                label="Password"
                name="password"
                :rules="[{ required: true, message: 'Please input your password!' }]"
            >
              <a-input-password v-model:value="formState.password"/>
            </a-form-item>

            <a-form-item name="remember" :wrapper-col="{ offset: 8, span: 16 }">
              <a-checkbox v-model:checked="formState.remember">Remember me</a-checkbox>
            </a-form-item>

            <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
              <a-button type="primary" html-type="submit">Submit</a-button>
            </a-form-item>
          </a-form>

        </a-modal>
      </div>
    </div>
  </div>


</template>

<script setup>
import {onMounted, reactive, ref} from 'vue';
import * as echarts from 'echarts';

// 日历组件
const value = ref();
const onPanelChange = (value, mode) => {
  console.log(value, mode);
};


// 模态对话框组件
const visible = ref(false);
const showModal = () => {
  visible.value = true;
};
const handleOk = e => {
  console.log(e);
  visible.value = false;
};

// form表单
const formState = reactive({
  username: '',
  password: '',
  remember: true,
});
const onFinish = values => {
  console.log('Success:', values);
};
const onFinishFailed = errorInfo => {
  console.log('Failed:', errorInfo);
};

// echarts饼图
const chart = ref();  // const chart = reactive({value:""})
// setBing
let setBingTu = () => {
  var myChart = echarts.init(chart.value);
  var option;
  option = {
    legend: {
      top: 'bottom'
    },
    toolbox: {
      show: true,
      feature: {
        mark: {show: true},
        dataView: {show: true, readOnly: false},
        restore: {show: true},
        saveAsImage: {show: true}
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
          {value: 40, name: '河南'},
          {value: 38, name: '湖南'},
          {value: 32, name: '湖北'},
          {value: 30, name: 'rose 4'},
          {value: 28, name: 'rose 5'},
          {value: 26, name: 'rose 6'},
          {value: 22, name: 'rose 7'},
          {value: 18, name: 'rose 8'}
        ]
      }
    ]
  };
  option && myChart.setOption(option);
}

onMounted(() => {
  setBingTu()
});


</script>

<style scoped>

.mychart {
  width: 600px;
  height: 600px;
}

.c1, .c2, .c3 {
  float: left;
  margin-left: 50px;
}



</style>