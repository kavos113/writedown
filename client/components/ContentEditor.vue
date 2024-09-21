<script setup lang="ts">
import { marked } from "marked";
import type { ResponseBody } from "~/lib/oapi-types";

const props = defineProps<{
    page: ResponseBody<'/pages/{pageID}', 'get', 200>;
}>();

const content = ref(props.page.body);
const output = computed(() => marked.parse(content.value));
</script>

<template>
    <div class="editor">
        <textarea v-model="content"/>
        <div v-html="output"/>
    </div>
</template>

<style scoped>
.editor {
    display: grid;
    grid-template-columns: 1fr 1fr;
}
</style>