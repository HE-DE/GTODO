<template>
  <el-main>
    <el-scrollbar>
      <el-table :data="pageData">
        <el-table-column prop="InfoID" label="ID" width="auto" min-width="25%" />
        <el-table-column prop="CreateTime" label="Time" width="auto" min-width="25%" />
        <el-table-column prop="Content" label="Content" width="auto" min-width="50%" />
      </el-table>
      <el-pagination style="justify-content: center" background layout="prev, pager, next" :total="changePage.total-1"
        @current-change="handleCurrentChange" v-model:current-page="changePage.currentPage" />
    </el-scrollbar>
  </el-main>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue'
import { useUsersStore } from '../store/user'
import API from '../plugins/axiosInterfaces'

const user = useUsersStore()
const tableData = ref([])
const pageData = ref([])
var changePage = reactive({
  currentPage: 1,
  total: tableData.value.length + 1 / 16,
})

//获取当前页数
var handleCurrentChange

API({
  url: '/api/getmsg',
  method: 'get',
  params: {
    id: user.Id
  }
}).then(res => {
  tableData.value = res.data.data
  pageData.value = tableData.value.slice(0, 16);
  changePage.total = tableData.value.length + 1 / 15
  handleCurrentChange = (value: number) => {


    //获取当前页码
    changePage.currentPage = value;


    //判断当前页是否为首页 页码从1开始，是则直接调用后端数据，否则要进行计算
    if (value > 1) {

      var i = (value - 1) * 16;  //计算当前页第一条数据的下标，

      var arry = [];  //建立一个临时数组

      //比如每页10条数据，第二页的第一条数据就是从 （2-1）*10 = 10 开始的 结束下标就是2*10=20 
      while (i < value * 16) {
        //解决最后一页出现null值
        if (tableData.value[i] != null) {
          arry.push(tableData.value[i]);
          i++;
          continue
        }
        break
      }
      pageData.value = arry

    } else {

      pageData.value = tableData.value.slice(0, 16);

    }
  };
})



</script>