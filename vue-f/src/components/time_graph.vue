<template>
    <el-main>
        <el-scrollbar>
            <div id="BarChart" ref="barchart"></div>
        </el-scrollbar>
    </el-main>
</template>
  
<script setup>
import { onMounted, reactive, ref } from "@vue/runtime-core"
import * as echarts from 'echarts'
import API from '../plugins/axiosInterfaces'
import { useUsersStore } from "../store/user";
const barchart = ref(null)
const user = useUsersStore()

var msgall = await API({
    url: '/api/getmsg',
    method: 'get',
    params: {
        id: user.Id
    }
}).then(res => {
   return res
})

console.log(msgall)

</script>