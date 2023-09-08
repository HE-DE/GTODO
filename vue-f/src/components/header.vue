<template>
    <div class="toolbar">
        <el-dropdown>
            <img src="../assets/img/head.jpeg" style="width: 40px;border-radius: 50px; margin-top: 10px;">
            <template #dropdown>
                <el-dropdown-menu>
                    <el-dropdown-item>
                        <el-icon>
                            <Tools></Tools>
                        </el-icon>
                        个人信息
                    </el-dropdown-item>
                    <el-dropdown-item>
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

</script>

