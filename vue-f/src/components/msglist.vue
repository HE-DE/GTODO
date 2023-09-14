<template>
  <el-main>
    <el-scrollbar>
      <el-table :data="tableData">
        <el-table-column prop="InfoID" label="ID" width="auto" min-width="25%" />
        <el-table-column prop="CreateTime" label="Time" width="auto" min-width="25%" />
        <el-table-column prop="Content" label="Content" width="auto" min-width="50%"/>
      </el-table>
    </el-scrollbar>
  </el-main>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useUsersStore } from '../store/user'
import API from '../plugins/axiosInterfaces'

const user = useUsersStore()
const tableData = ref([])
API({
  url: '/api/getmsg',
  method: 'get',
  params: {
    id: user.Id
  }
}).then(res => {
  console.log(res)
  tableData.value = res.data.data
})

const item = {
  date: '2016-05-02',
  name: 'Tom',
  address: 'No. 189, Grove St, Los Angeles',
}
</script>