<template>
    <el-form :model="form" label-width="120px">
        <el-form-item label="接受任务的用户" prop="userName">
            <el-select v-model="form.userName" class="m-2" placeholder="Select" size="large">
                <el-option v-for="item in ops" :key="item.value" :label="item.label" :value="item.value" />
            </el-select>
        </el-form-item>
        <el-form-item label="待办事项">
            <el-input v-model="form.content" type="textarea" />
        </el-form-item>
        <el-form-item>
            <el-button type="success" @click="onSubmit">添加</el-button>
            <el-button @click="onCancel">Cancel</el-button>
        </el-form-item>
    </el-form>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import API from '../plugins/axiosInterfaces'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router';
import { useUsersStore } from '../store/user';

// do not use same name with ref
const router = useRouter()
const user = useUsersStore()
const form = reactive({
    content: "",
    userName: "",
})


var ops = await API({
    url: '/api/getusername',
    method: 'get',
}).then(res => {
    var opsres = [] as any[]
    for (let i = 0; i < res.data.data.length; i++) {
        opsres.push({
            label: res.data.data[i],
            value: res.data.data[i]
        })
    }
    return opsres
})



const onSubmit = () => {
    if (form.content === "") {
        ElMessage.error("请输入内容")
        return
    } else {
        API({
            url: '/api/getuser',
            method: 'get',
            params: {
                name: form.userName
            }
        }).then(res => {
            console.log(res.data.data.UserID)
            API({
                url: '/api/addmsga',
                method: 'get',
                params: {
                    UserId: res.data.data.UserID,
                    Content: form.content,
                    AdminId: user.Id
                }
            }).then(res1 => {
                if (res1.data.msg === "success") {
                    ElMessage.success("添加成功")
                    setTimeout(function () {
                        router.push('/doing')
                    }, 1000)
                } else {
                    ElMessage.error(res1.data.msg)
                }
            })
        })
    }
}

const onCancel = () => {
    router.push('/doing')
}

</script>