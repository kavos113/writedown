<script setup lang="ts">
import { marked } from "marked";
import type { ResponseBody } from "~/lib/oapi-types";

const props = defineProps<{
    page: ResponseBody<'/pages/{pageID}', 'get', 200>;
}>();

const markedContent = marked.parse(props.page.body);
</script>

<template>
  <div class="wrapper">
    <h2>コンテント</h2>
    <div v-html="markedContent"></div>
  </div>
</template>

<style scoped>

</style>