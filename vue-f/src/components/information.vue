<template>
  <el-main>
    <el-scrollbar>
      <div>
        <el-page-header @back="onBack">
          <template #content>
            <div class="flex items-center">
              <el-avatar class="mr-3" :size="32" src="../assets/img/head.jpeg" />
              <span class="text-large font-600 mr-3"> UserName </span>
              <span class="text-sm mr-2" style="color: var(--el-text-color-regular)">
                ID
              </span>
              <el-tag>Default/Admin</el-tag>
            </div>
          </template>
          <template #extra>
            <div class="flex items-center">
              <el-button type="primary" class="ml-2">Edit</el-button>
            </div>
          </template>

          <el-descriptions :column="3" size="large" class="mt-4">
            <el-descriptions-item label="Username">kooriookami</el-descriptions-item>
            <el-descriptions-item label="Telephone">18100000000</el-descriptions-item>
            <el-descriptions-item label="Email">123@qq.com</el-descriptions-item>
            <el-descriptions-item label="Remarks">
              <el-tag size="small">admin</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="Msg">No.1188, Wuzhong Avenue, Wuzhong District, Suzhou, Jiangsu Province
            </el-descriptions-item>
          </el-descriptions>
          <p class="mt-4 text-sm">
            Element Plus team uses <b>weekly</b> release strategy under normal
            circumstance, but critical bug fixes would require hotfix so the actual
            release number <b>could be</b> more than 1 per week.
          </p>
        </el-page-header>
      </div>
      <div class="block text-center">
        <span class="demonstration">Switch when indicator is hovered (default)</span>
        <el-carousel height="430px">
          <el-carousel-item v-for="item in 4" :key="item">
            <img src="../assets/img/head.jpeg"/>
          </el-carousel-item>
        </el-carousel>
      </div>
    </el-scrollbar>
  </el-main>
</template>
  
<script lang="ts" setup>
import { useUsersStore } from '../store/user'
import API from '../plugins/axiosInterfaces'
import { ElNotification as notify } from 'element-plus'
const user = useUsersStore()
const onBack = () => {
  notify('Back')
}

</script>

<style scoped>
.demonstration {
  color: var(--el-text-color-secondary);
}

.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 150px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>