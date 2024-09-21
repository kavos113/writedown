<script setup lang="ts">
import { marked } from "marked";
import type { ResponseBody } from "~/lib/oapi-types";

const page = defineModel<ResponseBody<"/pages/{pageID}", "get", 200>>({
    required: true,
});

const output = computed(() => marked.parse(page.value.body));
</script>

<template>
    <div class="editor">
        <textarea v-model="page.body"/>
        <div v-html="output"/>
    </div>
</template>

<style scoped>
.editor {
    display: grid;
    grid-template-columns: 1fr 1fr;
}
</style>