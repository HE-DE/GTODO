<template>
    <div>
        <el-container>
            <Aside></Aside>
            <el-container>
                <el-header style="text-align: right; font-size: 24px" class="bg-blue-500">
                    <Header></Header>
                </el-header>
                <Suspense>
                    <template #default>
                        <Addmsga></Addmsga>
                    </template>
                </Suspense>
            </el-container>
        </el-container>
    </div>
</template>

<script lang="ts" setup>
import { defineAsyncComponent } from 'vue'
import Aside from '../components/aside.vue';
import Header from '../components/header_add.vue';
const Addmsga = defineAsyncComponent(() => import('../components/addmsga.vue'));

</script>