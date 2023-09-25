<template>
    <el-main>
        <el-scrollbar>
            <el-table :data="pageData" :row-class-name="tableRowClassName" height="100%" :row-style="{ height: '50px' }">
                <el-table-column prop="UserID" label="ID" width="auto" min-width="25%" />
                <el-table-column prop="Name" label="Name" width="auto" min-width="25%" />
                <el-table-column prop="Msging" label="Msging" width="auto" min-width="25%" />
                <el-table-column prop="Msged" label="Msged" width="auto" min-width="25%" />
            </el-table>
            <el-pagination style="justify-content: center" background layout="prev, pager, next"
                :total="changePage.total - 1" @current-change="handleCurrentChange"
                v-model:current-page="changePage.currentPage" />
        </el-scrollbar>
    </el-main>
</template>
  
<script lang="ts" setup>
import { ref, reactive } from 'vue'
import { useUsersStore } from '../store/user'
import API from '../plugins/axiosInterfaces'
import { ElMessage } from 'element-plus'

const user = useUsersStore()
const tableData = ref([] as any[])
const pageData = ref([] as any[])
var changePage = reactive({
    currentPage: 1,
    total: tableData.value.length + 1 / 12,
})

//获取当前页数
var handleCurrentChange

var tableRowClassName

API({
    url: '/api/getuserlist',
    method: 'get',
}).then(res => {
    console.log(res.data.data)
    tableData.value = res.data.data
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
        if (pageData.value[rowIndex].Name === user.username) {
            return 'doing-row'
        } else {
            return ''
        }
    }
})
</script>
  
<style>
.el-table .doing-row {
    --el-table-tr-bg-color: #b3e19d;
}
</style>