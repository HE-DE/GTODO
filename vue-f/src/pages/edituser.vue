<template>
    <div>
        <el-container>
            <Aside></Aside>
            <el-container>
                <el-header style="text-align: right; font-size: 24px" class="bg-blue-500">
                    <Header></Header>
                </el-header>
                    <UserEdit></UserEdit>
            </el-container>
        </el-container>
    </div>
</template>

<script lang="ts" setup>
import Aside from '../components/aside.vue';
import Header from '../components/header.vue';
import UserEdit from '../components/useredit.vue';
</script>

