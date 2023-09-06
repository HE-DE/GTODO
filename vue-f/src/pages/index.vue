<template>
    <div>
        <el-container class="min-h-screen bg-gray-50">
            <el-aside>
                <el-header class="bg-yellow-300 font-bold flex items-center justify-center">
                    《用户信息》
                </el-header>
                <el-menu @open="handleOpen" @close="handleClose" class="bg-gray-50 min-h-screen-md">
                    <el-sub-menu index="1">
                        <template #title>
                            <el-icon>
                                <Histogram />
                            </el-icon>
                            <span>数据统计</span>
                        </template>
                        <el-menu-item index="1-1">
                            <el-icon><Select /></el-icon>
                            事项完成情况
                        </el-menu-item>
                        <el-menu-item index="1-2">
                            <el-icon>
                                <PieChart />
                            </el-icon>
                            任务用时统计
                        </el-menu-item>
                    </el-sub-menu>
                    <el-menu-item index="2">
                        <el-icon>
                            <Grid />
                        </el-icon>
                        <span>待办事项列表</span>
                    </el-menu-item>
                    <el-menu-item index="3" class="">
                        <el-icon>
                            <UserFilled />
                        </el-icon>
                        <span>群组成员</span>
                    </el-menu-item>
                </el-menu>
            </el-aside>
            <el-container>
                <el-header>Header</el-header>
                <el-main>Main</el-main>
            </el-container>
        </el-container>
    </div>
</template>

<script lang="ts" setup>
import {
    UserFilled,
    Histogram,
    Select,
    PieChart,
    Grid,
} from '@element-plus/icons-vue'
const handleOpen = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}


</script>