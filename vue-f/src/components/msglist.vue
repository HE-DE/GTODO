<template>
  <el-main>
    <el-scrollbar>
      <el-table :data="pageData" :row-class-name="tableRowClassName" height="100%" :row-style="{ height: '50px' }">
        <el-table-column prop="InfoID" label="ID" width="auto" min-width="25%" />
        <el-table-column prop="CreateTime" label="Time" width="auto" min-width="25%" />
        <el-table-column prop="Content" label="Content" width="auto" min-width="50%" />
        <el-table-column fixed="right" label="Operations" width="240">
          <template v-slot="scope">
            <el-button color="#79bbff" size="small" @click="DoingUpdate(scope.row.InfoID)"><el-icon style="font-size: 15px;"><Loading /></el-icon></el-button>
            <el-button color="#b3e19d" size="small" ><el-icon style="font-size: 15px;"><Odometer /></el-icon></el-button>
            <el-button color="#fab6b6" size="small"><el-icon style="font-size: 15px;"><Lock /></el-icon></el-button>
            <el-button color="#b1b3b8" size="small"><el-icon style="font-size: 15px;"><Check /></el-icon></el-button>
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
import { Loading,Check,Odometer,Lock } from "@element-plus/icons-vue";

const user = useUsersStore()
const tableData = ref([] as any[])
const pageData = ref([] as any[])
var changePage = reactive({
  currentPage: 1,
  total: tableData.value.length + 1 / 12,
})

function DoingUpdate(InfoID:string){
  console.log(InfoID)
}

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
  pageData.value = tableData.value.slice(0, 12);
  changePage.total = tableData.value.length + 1 / 12
  handleCurrentChange = (value: number) => {
    //获取当前页码
    changePage.currentPage = value;
    //判断当前页是否为首页 页码从1开始，是则直接调用后端数据，否则要进行计算
    if (value > 1) {
      var i = (value - 1) * 12;  //计算当前页第一条数据的下标，
      var arry = [] as any[];  //建立一个临时数组
      //比如每页10条数据，第二页的第一条数据就是从 （2-1）*10 = 10 开始的 结束下标就是2*10=20 
      while (i < value * 12) {
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
      pageData.value = tableData.value.slice(0, 12);
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
    } else if (pageData.value[rowIndex].Status === 2) {
      return 'finishing-row'
    }
    return 'done-row'
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

.el-table .done-row {
  --el-table-tr-bg-color: #b1b3b8;
}
</style>