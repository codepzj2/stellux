<template>
    <div>
        <a-card>
            <template #title>stellux文档</template>
            <template #extra>
                <a-button type="primary" @click="router.push({ path: '/document/content' })">
                    前往stellux文档
                </a-button>
            </template>
        </a-card>
    </div>
</template>

<script setup lang='ts'>
import { useRouter } from 'vue-router';

defineOptions({
    name: 'DocumentIndex'
})

const router = useRouter();

router.push({ name: 'DocumentOverview' });
</script>

<style lang='scss' scoped>

</style>