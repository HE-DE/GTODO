<template>
    <el-row class="min-h-screen">
        <el-col :lg="16" :md="12" class="bg-indigo-500 flex items-center justify-center flex-col">
            <div>
                <div class="font-bold text-7xl text-light-50 mb-4">GTODO</div>
                <div class="text-gray-200 text-xl">一个使用Go+Vue+Vite开发的备忘录</div>
            </div>
        </el-col>
        <el-col :lg="8" :md="12" class="bg-indigo-50 flex items-center justify-center flex-col">
            <h2 class="font-bold text-5xl text-gray-800 ">欢迎回来</h2>
            <div class="flex items-center justify-center my-6 text-gray-400 space-x-3">
                <span class="h-[1px] w-16 bg-gray-400"></span>
                <span>账号密码登录</span>
                <span class="h-[1px] w-16 bg-gray-400"></span>
            </div>
            <el-form :model="form" class="w-[250px]">
                <el-form-item>
                    <el-input v-model="form.username" placeholder="请输入用户名" />
                </el-form-item>
                <el-form-item>
                    <el-input v-model="form.password" type="password" placeholder="请输入密码" />
                </el-form-item>
                <el-form-item>
                    <el-button round color="#626aef" class="w-[250px]" type="primary" @click="onSubmit">登录</el-button>
                </el-form-item>
            </el-form>
        </el-col>
    </el-row>
</template>

<script setup>
import { reactive } from 'vue'
import API from '../plugins/axiosInterfaces'
import { useRouter } from 'vue-router';
import { useUsersStore } from '../store/user'
import { ElMessage } from 'element-plus';
// do not use same name with ref
const router = useRouter()
const user = useUsersStore()
const form = reactive({
    username: "",
    password: ""
})

const onSubmit = function () {
    API({
        url: '/api/login',
        method: 'post',
        data: {
            name: form.username,
            password: form.password
        }
    }).then((res) => {
        console.log(res)
        var ID = res.data.data.UserID
        var IsAdmin = res.data.data.Indentity
        if (res.data.status === "error") {
            if (res.data.message === "用户名不存在") {
                ElMessage({
                    message: '用户名不存在,进入注册',
                    type: 'warning'
                })
                setTimeout(function () {
                    router.push('/register')
                }, 1000)
                return
            } else if (res.data.message === "密码错误") {
                ElMessage.error('密码错误')
                return
            }
        }
        if (IsAdmin === 1) {
            IsAdmin = true
        } else {
            IsAdmin = false
        }
        user.Login(form.username, IsAdmin, ID)
        console.log(user.username)
        console.log(user.Id)
        console.log(user.isLogin)
        ElMessage({
            message: '登录成功！',
            type: 'success'
        })
        setTimeout(function () {
            router.push('/')
        }, 1000)
    })
}
</script>