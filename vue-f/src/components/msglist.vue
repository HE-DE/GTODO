<template>
  <el-main>
    <el-scrollbar>
      <el-table :data="tableData">
        <el-table-column prop="date" label="Date" width="140" />
        <el-table-column prop="name" label="Name" width="120" />
        <el-table-column prop="address" label="Address" />
      </el-table>
    </el-scrollbar>
  </el-main>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useUsersStore } from '../store/user'
import API from '../plugins/axiosInterfaces'


const user = useUsersStore()
console.log(user.username)
API({
  url: '/api/getuser',
  method: 'get',
  params: {
    name: "t1"
  }
}).then(res => {
  console.log(res)
})

const item = {
  date: '2016-05-02',
  name: 'Tom',
  address: 'No. 189, Grove St, Los Angeles',
}
const tableData = ref(Array.from({ length: 20 }).fill(item))
</script>