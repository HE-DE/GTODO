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
                        <ThingG></ThingG>
                    </template>
                </Suspense>
            </el-container>
        </el-container>
    </div>
</template>

<script lang="ts" setup>
import { defineAsyncComponent } from 'vue'
import Aside from '../components/aside.vue';
import Header from '../components/header_doing.vue';
const ThingG = defineAsyncComponent(() => import('../components/thing_graph.vue'));
</script>