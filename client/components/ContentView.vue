<script
  setup
  lang="ts"
>
import { marked } from "marked";

const props = defineProps<{
    pageId: number;
}>();
const response = await useFetch(`/api/pages/${ props.pageId }`);
const markedContent = marked.parse(response.data.value.body);
</script>

<template>
  <h2>コンテント</h2>
</template>

<style scoped>

</style>