<script setup lang="ts">
import type { ResponseBody, RequestBody } from "~/lib/oapi-types";
import ContentEditor from "./ContentEditor.vue";

const { $toast } = useNuxtApp();
const props = defineProps<{
  parentId: number;
}>();

const editor = ref<InstanceType<typeof ContentEditor> | null>(null);

const page: ResponseBody<"/pages/{pageID}", "get", 200> = {
  id: 0,
  path: "",
  name: "",
  title: "",
  body: "",
  creator: "",
  created_at: "",
  updated_at: "",
};

const save = () => {
  if (editor.value) {
    editor.value.update();
  }
  console.log(page);

  const req: RequestBody<"/pages", "post"> = {
    parentID: props.parentId,
    name: page.name,
    title: page.title,
    body: page.body,
    creator: page.creator,
  };
  useFetch(`/api/pages`, {
    method: "POST",
    body: JSON.stringify(req),
    onResponse: (res) => {
      if (res.response.ok) {
        $toast.success("保存しました");
      } else {
        $toast.error("保存に失敗しました");
      }
    },
  });
};
</script>

<template>
  <div>
    <div>
      <div class="header">
        <h1 class="title">新規ページ</h1>
      </div>
      <button @click="save">保存</button>
      <button>キャンセル</button>
      <ContentEditor ref="editor" v-model="page" />
    </div>
  </div>
  <SideBarRight class="sideBarRight" />
</template>

<style scoped></style>
