<script setup lang="ts">
import type { ResponseBody } from '~/lib/oapi-types';

const props = defineProps<{
    parentId: number;
}>();

const page: ResponseBody<'/pages/{pageID}', 'get', 200> = {
    id: 0,
    path: '',
    name: '',
    title: '',
    body: '',
    creator: '',
    created_at: '',
    updated_at: '',
};
</script>

<template>
    <div>
        <div>
            <div class="header">
            <h1 class="title">新規ページ</h1>
            </div>
            <button>保存</button>
            <button>キャンセル</button>
            <ContentEditor v-model="page" />
        </div>
    </div>
    <SideBarRight class="sideBarRight"/>
</template>

<style scoped>
</style>