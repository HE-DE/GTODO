<template>
    <el-form :model="form" label-width="120px">
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
})


const onSubmit = () => {
    if (form.content === "") {
        ElMessage.error("请输入内容")
        return
    } else {
        API({
            url: '/api/addmsgu',
            method: 'get',
            params: {
                UserId: user.Id,
                Content: form.content
            }
        }).then(res => {
            if (res.data.msg === "success") {
                ElMessage.success("添加成功")
                setTimeout(function () {
                    router.push('/information')
                }, 1000)
            } else {
                ElMessage.error(res.data.msg)
            }
        })
    }
}

const onCancel = () => {
    router.push('/doing')
}

</script>