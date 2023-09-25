<template>
    <el-aside>
        <el-menu @open="handleOpen" @close="handleClose" class="bg-gray-50" style="height: 100vh;">
            <el-sub-menu index="1">
                <template #title>
                    <el-icon>
                        <Histogram />
                    </el-icon>
                    <span>数据统计</span>
                </template>
                <el-menu-item index="1-1" @click="Finishing">
                    <el-icon><Select /></el-icon>
                    事项完成情况
                </el-menu-item>
                <el-menu-item index="1-2" @click="Timing">
                    <el-icon>
                        <PieChart />
                    </el-icon>
                    任务用时统计
                </el-menu-item>
            </el-sub-menu>
            <el-menu-item index="2" @click="TODOThing">
                <el-icon>
                    <Grid />
                </el-icon>
                <span>待办事项列表</span>
            </el-menu-item>
            <el-menu-item index="3" @click="UserList">
                <el-icon>
                    <UserFilled />
                </el-icon>
                <span>群组成员</span>
            </el-menu-item>
        </el-menu>
    </el-aside>
</template>

<script lang="ts" setup>
import {
    UserFilled,
    Histogram,
    Select,
    PieChart,
    Grid,
} from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
const router = useRouter()
const handleOpen = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
function TODOThing() {
    router.push('/doing')
}
function UserList(){
    router.push('/userlist')
}
function Timing() {
    router.push('/timing')
}
function Finishing(){
    router.push('/finishing')
}

</script>