<template>
  <el-main>
    <el-scrollbar>
      <el-table :data="pageData" :row-class-name="tableRowClassName">
        <el-table-column prop="InfoID" label="ID" width="auto" min-width="25%" />
        <el-table-column prop="CreateTime" label="Time" width="auto" min-width="25%" />
        <el-table-column prop="Content" label="Content" width="auto" min-width="50%" />
        <!-- <el-table-column prop="Status" label="Status" width="auto" min-width="50%">
          <template #default="scope">
              <span v-if="scope.row.Status===0">doing</span>
            </template>
        </el-table-column> -->
        <el-table-column fixed="right" label="Operations" width="120">
          <template #default>
            <el-button link type="primary" size="small" @click="handleClick">Detail</el-button>
            <el-button link type="primary" size="small">Edit</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination style="justify-content: center" background layout="prev, pager, next" :total="changePage.total - 1"
        @current-change="handleCurrentChange" v-model:current-page="changePage.currentPage" />
    </el-scrollbar>
  </el-main>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue'
import { useUsersStore } from '../store/user'
import API from '../plugins/axiosInterfaces'

const user = useUsersStore()
const tableData = ref([] as any[])
const pageData = ref([] as any[])
var changePage = reactive({
  currentPage: 1,
  total: tableData.value.length + 1 / 16,
})

var handleClick
//获取当前页数
var handleCurrentChange

var tableRowClassName


API({
  url: '/api/getmsg',
  method: 'get',
  params: {
    id: user.Id
  }
}).then(res => {
  console.log(res.data.data)
  tableData.value = res.data.data
  for (let i = 0; i < tableData.value.length; i++) {
    tableData.value[i].CreateTime = res.data.CTime[i]
    tableData.value[i].DoneTime = res.data.DTime[i]
  }
  pageData.value = tableData.value.slice(0, 16);
  changePage.total = tableData.value.length + 1 / 15
  handleCurrentChange = (value: number) => {
    //获取当前页码
    changePage.currentPage = value;
    //判断当前页是否为首页 页码从1开始，是则直接调用后端数据，否则要进行计算
    if (value > 1) {
      var i = (value - 1) * 16;  //计算当前页第一条数据的下标，
      var arry = [] as any[];  //建立一个临时数组
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
  tableRowClassName = ({
    rowIndex,
  }: {
    rowIndex: number
  }) => {
    if (pageData.value[rowIndex].Status === 0) {
      return 'doing-row'
    } else if (pageData.value[rowIndex].Status === 1) {
      return 'buzying-row'
    } else if (pageData.value[rowIndex].Status === 2){
      return 'finishing-row'
    }
      return ''
  }
})
</script>

<style>
.el-table .buzying-row {
  --el-table-tr-bg-color: #79bbff;
}

.el-table .doing-row {
  --el-table-tr-bg-color: #b3e19d;
}

.el-table .finishing-row {
  --el-table-tr-bg-color: #fab6b6;
}
</style>