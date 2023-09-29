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

var doingTime = await API({
    url: '/api/updatedoing',
    method: 'get',
    params: {
        id: user.Id
    }
}).then(res => {
    return res.data.data
})

var AdminTime = 0
var UserTime = 0

for (let i = 0; i < msgall.length; i++) {
    if (msgall[i].Status === 3) {
        if (msgall[i].AdminId === 0) {
            UserTime = UserTime + doingTime[i];
        } else {
            AdminTime = AdminTime + doingTime[i];
        }
    }
}

//1.通过vue3.x中的refs标签用法，获取到container容器实例
const chart = ref(null);
//2.定义在本vue实例中的echarts实例
let myEcharts = reactive({});
//4.定义好echarts的配置数据
let option = {
    tooltip: {
        trigger: 'item'
    },
    legend: {
        top: '5%',
        left: 'center'
    },
    series: [
        {
            name: 'Access From',
            type: 'pie',
            radius: ['40%', '70%'],
            center: ['30%', '50%'],
            avoidLabelOverlap: false,
            label: {
                show: false,
                position: 'center'
            },
            emphasis: {
                label: {
                    show: true,
                    fontSize: 40,
                    fontWeight: 'bold'
                }
            },
            labelLine: {
                show: false
            },
            data: [
                { value: UserTime, name: "用户本人完成事项时间" },
                { value: AdminTime, name: "管理员分配事项的时间" }
            ]
        }
    ]
}

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