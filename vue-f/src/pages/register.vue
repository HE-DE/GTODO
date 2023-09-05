<template>
    <el-row class="min-h-screen">
        <el-col :lg="8" :md="12" class="bg-indigo-500 flex items-center justify-center flex-col">
            <div>
                <div class="font-bold text-7xl text-light-50 mb-4">GTODO</div>
                <div class="text-gray-200 text-xl">一个使用Go+Vue+Vite开发的备忘录</div>
            </div>
        </el-col>
        <el-col :lg="16" :md="12" class="bg-indigo-50 flex items-center justify-center flex-col">
            <h2 class="font-bold text-3xl text-gray-800 ">用户注册</h2>
            <div class="flex items-center justify-center my-6 text-gray-400 space-x-3">
                <el-form :model="form" label-width="120px" ref="ruleFormRef" status-icon :rules="rules">
                    <el-form-item label="用户名">
                        <el-input v-model="form.name" placeholder="输入你的用户名" class="w-[650px]" />
                    </el-form-item>
                    <el-form-item label="身份信息">
                        <el-select v-model="form.identity" placeholder="选择你的身份">
                            <el-option label="管理员" value=0 />
                            <el-option label="普通用户" value=1 />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="电话号码" prop="phone">
                        <el-input v-model="form.phone" placeholder="输入你的电话号码" />
                    </el-form-item>
                    <el-form-item label="E-mail" prop="email">
                        <el-input v-model="form.email" placeholder="输入你的Email" />
                    </el-form-item>
                    <el-form-item label="选择性别">
                        <el-radio-group v-model="form.sex">
                            <el-radio label="男" value="0" />
                            <el-radio label="女" value="1" />
                        </el-radio-group>
                    </el-form-item>
                    <el-form-item label="密码" prop="pwd">
                        <el-input v-model="form.pwd" placeholder="输入你的密码" type="password" autocomplete="off" />
                    </el-form-item>
                    <el-form-item label="确认密码" prop="checkpwd">
                        <el-input v-model="form.checkpwd" placeholder="再次输入你的密码" type="password" autocomplete="off" />
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="onSubmit(ruleFormRef)">Create</el-button>
                        <el-button @click="onCancel">Cancel</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </el-col>
    </el-row>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'

const ruleFormRef = ref<FormInstance>()
// do not use same name with ref
const form = reactive({
    name: '',
    identity: '',
    phone: '',
    email: '',
    sex: '',
    pwd: '',
    checkpwd: '',
})
// 设置校验规则
const checkPhone = (rule: any, value: any, callback: any) => {
    if (!value) {
        return callback(new Error("请输入电话号码"))
    }
    console.log("checkPhone")
    setTimeout(() => {
        if (value === null) {
            callback(new Error('输入电话号码不可为空'))
        } else if (value === '') {
            callback(new Error('输入电话号码不可为空'))
        } else if (!/^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/.test(value)) {
            callback(new Error('请输入正确的电话号码格式'))
        } else {
            callback()
        }
    }, 1000)
}

const checkEmail = (rule: any, value: any, callback: any) => {
    if (!value) {
        callback(new Error('请输入E-mail'))
    }
    const regEmail = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(\.[a-zA-Z0-9_-])+/
    setTimeout(() => {
        if (value === '') {
            callback(new Error('E-mail不能为空'))
        } else if (!regEmail.test(value)) {
            callback(new Error('E-mail的格式有误'))
        } else {
            callback()
        }
    }, 1000)
}


const rules = reactive<FormRules<typeof form>>({
    phone: [{ validator: checkPhone, trigger: 'blur' }],
    email: [{validator:checkEmail,trigger:'blur'}]
})


const onSubmit = (formEl: FormInstance | undefined) => {
    console.log('submit!')
    if (!formEl) return
    formEl.validate((valid) => {
        if (valid) {
            console.log('submit!')
        } else {
            console.log('error submit!')
            return false
        }
    })
}

const onCancel = () => {
    console.log('Cancel!')
}

</script>