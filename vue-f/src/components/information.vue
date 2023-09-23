<template>
  <el-main>
    <el-scrollbar>
      <div>
        <el-page-header title="Back" @back="onBack">
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
            <el-descriptions-item label="Telephone">{{ Phone }}</el-descriptions-item>
            <el-descriptions-item label="Email">{{ Email }}</el-descriptions-item>
            <el-descriptions-item label="Sex">
              <el-tag  v-if=" Sex === 0 ">man</el-tag>
              <el-tag  v-if=" Sex === 1 ">girl</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="Msg">You have {{ Msging }} messages to do!
            </el-descriptions-item>
          </el-descriptions>
          <p class="text-green-500" align="center">
            <b>愿此行终抵群星！</b>
          </p>
        </el-page-header>
      </div>
      <div class="block text-center">
        <el-carousel height="490px">
          <el-carousel-item v-for=" item  in  imageUrl " :key=" item ">
            <div style=" display: flex;justify-content: center;align-items: center; ">
              <img :src=" item.url " alt="" />
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
  router.push("/information")
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

const Email = ref("")
const Phone = ref("")
const Sex = ref()
const Msging=ref()
API({
  url: '/api/getuser',
  method: 'get',
  params: {
    name: username.value,
  }
}).then(res => {
  Email.value = res.data.data.Email
  Phone.value = res.data.data.Phone
  Sex.value = res.data.data.Sex
  Msging.value = res.data.data.Msging
})


//图片的滚动播放池
const imageUrl = [
  { url: "https://www.cioinsight.com.cn/uploads/images/20230223/20230223163502_18189.jpg" },
  { url: "https://img.yanlutong.com/uploadimg/image/20230217/20230217095912_63131.jpg" },
  { url: "https://pic2.zhimg.com/v2-090dee1d971cdcd3520f7f6bf4b2749c_1440w.jpg?source=172ae18b" },
  { url: "https://pic.kts.g.mi.com/3c921c7ee6b63ca648cc6798c6de39b25686095922872919379.png" }];

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