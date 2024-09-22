<script setup lang="ts">
import { marked } from "marked";
import type { ResponseBody } from "~/lib/oapi-types";

const page = defineModel<ResponseBody<"/pages/{pageID}", "get", 200>>({
  required: true,
});

const content = ref(page.value.body);
const output = computed(() => marked.parse(content.value));

const update = () => {
  page.value.body = content.value;
};

defineExpose({
  update,
});
</script>

<template>
  <div class="editor">
    <textarea
      v-model="content"
      class="textarea"
    />
    <div
      v-html="output"
      class="output"
    />
  </div>
</template>

<style scoped>
.editor {
  display: grid;
  grid-template-columns: 1fr 1fr;
  border: 1px solid #aaa;
}

.textarea {
  width: inherit;
  height: inherit;
  border: none;
  resize: none;
  border-right: 1px solid #aaa;
  background-color: #eee;
  padding: 10px;
}

.output {
  padding: 10px;
}
</style>
