<template>
  <el-main>
    <el-scrollbar>
      <div>
        <el-page-header title="login" @back="onBack">
          <template #content>
            <div class="flex items-center">
              <el-avatar class="mr-3" :size="32"
                src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" />
              <span class="text-large font-600 mr-3"> {{ username }} </span>
              <span class="text-sm mr-2" style="color: var(--el-text-color-regular)">
                {{ ID }}
              </span>
              <el-tag v-if="isAdmin">admin</el-tag>
              <el-tag v-if="isUser">User</el-tag>
            </div>
          </template>
          <template #extra>
            <div class="flex items-center">
              <el-button type="primary" class="ml-2" @click="Edit">Edit</el-button>
            </div>
          </template>

          <el-descriptions :column="3" size="large" class="mt-4">
            <el-descriptions-item label="Username">{{ username }}</el-descriptions-item>
            <el-descriptions-item label="Telephone">18100000000</el-descriptions-item>
            <el-descriptions-item label="Email">123@qq.com</el-descriptions-item>
            <el-descriptions-item label="Remarks">
              <el-tag size="small" v-if="isAdmin">admin</el-tag>
              <el-tag size="small" v-if="isUser">User</el-tag>
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
        <el-carousel height="490px">
          <el-carousel-item v-for="item in imageUrl" :key="item">
            <div style=" display: flex;justify-content: center;align-items: center; ">
              <img :src="item.url" alt="" />
            </div>
          </el-carousel-item>
        </el-carousel>
      </div>
    </el-scrollbar>
  </el-main>
</template>
  
<script lang="ts" setup>
import { ref } from 'vue'
import { useUsersStore } from '../store/user'
import API from '../plugins/axiosInterfaces'
import { useRouter } from 'vue-router'
const router = useRouter()
function Edit() {
  router.push("/edituser")
}

const onBack = () => {
  router.push("/")
}
const user = useUsersStore()
const username = ref(user.username)
const ID = ref(user.Id)
const isAdmin = user.isAdmin
var isUser
if (isAdmin === false) {
  isUser = true
} else {
  isUser = false
}
const imageUrl = [
  { url: "http://www.gengdan.cn/wp-content/uploads/2023/03/2023030921323441.jpg" },
  { url: "http://www.gengdan.cn/wp-content/uploads/2023/03/2023030921333114.jpg" },
  { url: "http://www.gengdan.cn/wp-content/uploads/2023/03/2023030921323441.jpg" },
  { url: "http://www.gengdan.cn/wp-content/uploads/2023/03/2023030921325368.jpg" }];

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