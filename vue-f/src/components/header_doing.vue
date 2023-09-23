<template>
    <div>
        <el-button type="primary" :icon="Back" @click="OnBack" style="float: left;margin-top: 10px;" circle />
        <el-button type="success" :icon="Edit" @click="onsubmit" style="float: left;margin-top: 10px;" circle />
        <el-dropdown :hide-on-click="false" @command="handleCommand">
            <img src="../assets/img/head.jpeg" style="width: 40px;border-radius: 50px; margin-top: 10px;">
            <template #dropdown>
                <el-dropdown-menu>
                    <el-dropdown-item command="user">
                        <el-icon>
                            <Tools></Tools>
                        </el-icon>
                        个人信息
                    </el-dropdown-item>
                    <el-dropdown-item command="logout">
                        <el-icon>
                            <House></House>
                        </el-icon>
                        退出登录
                    </el-dropdown-item>
                </el-dropdown-menu>
            </template>
        </el-dropdown>
    </div>
</template>

<script lang="ts" setup>
import { Tools, House } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router';
import { Back, Edit } from '@element-plus/icons-vue'
import { useUsersStore } from '../store/user';

const router = useRouter()
const user = useUsersStore()
const handleCommand = (command: string | number | object) => {
    if (command === 'logout') {
        window.sessionStorage.clear()
        window.location.href = '/'
    } else if (command === 'user') {
        router.push('/user')
    }
}

function OnBack(){
    router.push('/information')
}

function onsubmit(){
    router.push('/addmsg')
}

</script>

