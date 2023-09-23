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

                    <el-descriptions :column="3" size="large" class="mt-4">
                        <el-descriptions-item label="Username">{{ username }}</el-descriptions-item>
                        <el-descriptions-item label="Telephone">{{ Phone }}</el-descriptions-item>
                        <el-descriptions-item label="Email">{{ Email }}</el-descriptions-item>
                        <el-descriptions-item label="Sex">
                            <el-tag v-if="Sex === 0">man</el-tag>
                            <el-tag v-if="Sex === 1">girl</el-tag>
                        </el-descriptions-item>
                        <el-descriptions-item label="Msg">You have {{ Msging }} messages to do!
                        </el-descriptions-item>
                    </el-descriptions>
                    <p class="text-green-500" align="center">
                        <b>原此行终抵群星！</b>
                    </p>
                </el-page-header>
            </div>
            <div>
                <el-form :model="form" label-width="120px" ref="ruleFormRef" status-icon :rules="rules">
                    <el-form-item label="电话号码" prop="phone">
                        <el-input v-model="form.phone" placeholder="输入你的电话号码" />
                    </el-form-item>
                    <el-form-item label="E-mail" prop="email">
                        <el-input v-model="form.email" placeholder="输入你的Email" />
                    </el-form-item>
                    <el-form-item label="选择性别">
                        <el-radio-group v-model="form.sex">
                            <el-radio label="男" value=0 />
                            <el-radio label="女" value=1 />
                        </el-radio-group>
                    </el-form-item>
                    <el-form-item label="密码" prop="pwd">
                        <el-input v-model="form.pwd" placeholder="输入你的密码" type="password" autocomplete="off" />
                    </el-form-item>
                    <el-form-item label="确认密码" prop="checkpwd">
                        <el-input v-model="form.checkpwd" placeholder="再次输入你的密码" type="password" autocomplete="off" />
                    </el-form-item>
                    <el-form-item>
                        <el-button type="danger" @click="onSubmit(ruleFormRef)" round>确认修改</el-button>
                        <el-button type="primary" @click="onCancel" round>取消</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </el-scrollbar>
    </el-main>
</template>
    
<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { useUsersStore } from '../store/user'
import API from '../plugins/axiosInterfaces'
import { useRouter } from 'vue-router'
import type { FormInstance, FormRules } from 'element-plus'
import { ElMessage } from 'element-plus'

const router = useRouter()
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
const Msging = ref()
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

const form = reactive({
    name: username,
    phone: '',
    email: '',
    sex: '',
    pwd: '',
    checkpwd: '',
})
const ruleFormRef = ref<FormInstance>()
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

const validPwd = (rule: any, value: any, callback: any) => {
    if (value === '') {
        callback(new Error('密码不能为空'))
    } else {
        if (form.checkpwd !== '') {
            if (!ruleFormRef.value) return
            ruleFormRef.value.validateField('checkpwd', () => null)
        }
        callback()
    }
}

const validPwd2 = (rule: any, value: any, callback: any) => {
    if (value === '') {
        callback(new Error('请再次输入密码'))
    } else if (value !== form.pwd) {
        callback(new Error("两次输入的密码不一致"))
    } else {
        callback()
    }
}

const rules = reactive<FormRules<typeof form>>({
    phone: [{ validator: checkPhone, trigger: 'blur' }],
    email: [{ validator: checkEmail, trigger: 'blur' }],
    pwd: [{ validator: validPwd, trigger: 'blur' }],
    checkpwd: [{ validator: validPwd2, trigger: 'blur' }]
})


const onSubmit = (formEl: FormInstance | undefined) => {
    console.log('submit!')
    if (!formEl) return
    var sex = '0'
    if (form.sex === "女") {
        sex = '1'
    }
    formEl.validate((valid) => {
        if (valid) {
            API({
                url: '/api/updateuser',
                method: 'post',
                data: {
                    name: username.value,
                    phone: form.phone,
                    email: form.email,
                    sex: sex,
                    password: form.pwd,
                }
            }).then((res) => {
                if (res.data.status === "error") {
                    ElMessage.error(res.data.message)
                } else {
                    ElMessage.success(res.data.message)
                    // 跳转到登录页面
                    setTimeout(function () {
                        router.push('/login')
                    }, 1000)
                }
            })
        } else {
            console.log('error submit!')
            return false
        }
    })
}

const onCancel = () => {

    router.push('/information')
}

const onBack = () => {
    router.push("/user")
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