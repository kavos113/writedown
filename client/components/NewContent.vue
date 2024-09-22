<script setup lang="ts">
import type { ResponseBody, RequestBody } from "~/lib/oapi-types";
import ContentEditor from "./content/ContentEditor.vue";

const { $toast } = useNuxtApp();
const props = defineProps<{
  parentId: number;
  parentPath: string;
}>();
const router = useRouter();

const editor = ref<InstanceType<typeof ContentEditor> | null>(null);
const isView = ref(true);

const page: ResponseBody<"/pages/{pageID}", "get", 200> = {
  id: 0,
  path: props.parentPath,
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

const cancel = () => {
  router.back();
};
</script>

<template>
  <div>
    <div class="wrapper">
      <ContentTitle
        ref="contentTitle"
        v-model:title="page.title"
        v-model:is-view="isView"
        class="contentTitle"
      />
      <div class="buttonWrapper">
        <MenuButton @click="save">保存</MenuButton>
        <MenuButton @click="cancel">キャンセル</MenuButton>
      </div>
      <ContentBreadCrumbs
        :path="page.path"
        class="contentBreadCrumbs"
      />
      <ContentEditor
        ref="editor"
        v-model="page"
        class="contentEditor"
      />
    </div>
  </div>
  <SideBarRight class="sideBarRight" />
</template>

<style scoped>
.wrapper {
  display: grid;
  grid-template-rows: 55px 24px 36px 1fr;
  gap: 10px;
}

.contentTitle {
  grid-row: 1;
}

.buttonWrapper {
  grid-row: 3;
}

.contentBreadCrumbs {
  grid-row: 2;
}

.contentEditor {
  grid-row: 4;
}
</style>
