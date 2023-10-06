<template>
    <el-main>
        <el-scrollbar>
            <div id="container" ref="chart">
            </div>
        </el-scrollbar>
    </el-main>
</template>
  
<script setup>
import { onMounted, reactive, ref } from "@vue/runtime-core"
import * as echarts from 'echarts'
import API from '../plugins/axiosInterfaces'
import { useUsersStore } from "../store/user";
const user = useUsersStore()

var msgall = await API({
    url: '/api/getmsg',
    method: 'get',
    params: {
        id: user.Id
    }
}).then(res => {
    return res.data.data
})

var countDone = 0

for (let i = 0; i < msgall.length; i++) {
    if (msgall[i].Status === 3) {
        countDone++
    }
}

var grade = countDone / msgall.length

//1.通过vue3.x中的refs标签用法，获取到container容器实例
const chart = ref(null);
//2.定义在本vue实例中的echarts实例
let myEcharts = reactive({});
//4.定义好echarts的配置数据
let option = {
    series: [
    {
      type: 'gauge',
      startAngle: 180,
      endAngle: 0,
      center: ['50%', '75%'],
      radius: '90%',
      min: 0,
      max: 1,
      splitNumber: 8,
      axisLine: {
        lineStyle: {
          width: 6,
          color: [
            [0.25, '#FF6E76'],
            [0.5, '#FDDD60'],
            [0.75, '#58D9F9'],
            [1, '#7CFFB2']
          ]
        }
      },
      pointer: {
        icon: 'path://M12.8,0.7l12,40.1H0.7L12.8,0.7z',
        length: '12%',
        width: 20,
        offsetCenter: [0, '-60%'],
        itemStyle: {
          color: 'auto'
        }
      },
      axisTick: {
        length: 12,
        lineStyle: {
          color: 'auto',
          width: 2
        }
      },
      splitLine: {
        length: 20,
        lineStyle: {
          color: 'auto',
          width: 5
        }
      },
      axisLabel: {
        color: '#464646',
        fontSize: 20,
        distance: -60,
        rotate: 'tangential',
        formatter: function (value) {
          if (value === 0.875) {
            return 'Grade A';
          } else if (value === 0.625) {
            return 'Grade B';
          } else if (value === 0.375) {
            return 'Grade C';
          } else if (value === 0.125) {
            return 'Grade D';
          }
          return '';
        }
      },
      title: {
        offsetCenter: [0, '-10%'],
        fontSize: 20
      },
      detail: {
        fontSize: 30,
        offsetCenter: [0, '-35%'],
        valueAnimation: true,
        formatter: function (value) {
          return Math.round(value * 100) + '';
        },
        color: 'inherit'
      },
      data: [
        {
          value: grade,
          name: '综合得分'
        }
      ]
    }
  ]
};

//onMounted钩子函数
onMounted(() => {
    //初始化echarts
    init();
})

//初始化echarts实例方法
const init = () => {
    //3.初始化container容器
    myEcharts = echarts.init(chart.value);
    //5.传入数据
    myEcharts.setOption(option);
    //additional：图表大小自适应窗口大小变化
    window.onresize = () => {
        myEcharts.resize();
    }
}

</script>
<style scoped>
#container {
    box-sizing: border-box;
    height: 600px;
    width: 100%;
    margin: 0 auto;
}
</style>