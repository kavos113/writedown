<script setup lang="ts">
import type { ResponseBody } from "~/lib/oapi-types";

const props = defineProps<{
    pageId: number;
}>();
const res = await useFetch(`/api/pages/${ props.pageId }`);
const page = res.data.value as ResponseBody<'/pages/{pageID}', 'get', 200>;

const isView = ref(true);

const toEdit = () => {
    isView.value = false;
}
const toView = () => {
    isView.value = true;
}
</script>

<template>
    <div>
        <button v-if="isView" @click="toEdit">編集</button>
        <button v-else @click="toView">保存</button>
        <ContentView class="contentView" :page="page"v-if="isView"/>
        <ContentEditor class="contentEditor" :page="page" v-else/>
    </div>
</template>

<style scoped>

</style>