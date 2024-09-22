<script setup lang="ts">
import type { ResponseBody, RequestBody } from "~/lib/oapi-types";
import ContentEditor from "./ContentEditor.vue";

const { $toast } = useNuxtApp();
const router = useRouter();

const props = defineProps<{
    pageId: number;
}>();
const res = await useFetch(`/api/pages/${ props.pageId }`);
const page = res.data.value as ResponseBody<'/pages/{pageID}', 'get', 200>;

const isView = ref(true);
const editor = ref<InstanceType<typeof ContentEditor> | null>(null);

const edit = () => {
    isView.value = false;
};

const update = () => {
    isView.value = true;
    if(editor.value) {
        editor.value.update();
    }
    const req: RequestBody<'/pages/{pageID}', 'patch'> = {
        id: page.id,
        parentID: 0, // TODO: openapi修正
        name: page.name,
        title: page.title,
        body: page.body,
        creator: page.creator,
    };
    useFetch(`/api/pages/${ page.id }`, {
        method: 'PATCH',
        body: JSON.stringify(req),
        onResponse: (res) => {
            if (res.response.ok) {
                $toast.success('保存しました');
            } else {
                $toast.error('保存に失敗しました');
            }
        },
    });
};

const newPage = () => {
    router.push({
        name: 'new',    
        params: { parentId: page.id },
    });
}
</script>

<template>
    <div>
        <div class="header">
        <h1 class="title">{{ page.title }}</h1>
        </div>
        <button v-if="isView" @click="edit">編集</button>
        <button v-else @click="update">保存</button>
        <button @click="newPage">新規</button>
        <ContentView v-if="isView" class="contentView" :page="page"/>
        <ContentEditor v-else ref="editor" v-model="page" class="contentEditor"/>
    </div>
    <SideBarRight class="sideBarRight"/>
</template>

<style scoped>
.header {
    border-bottom: 1px solid #aaa;
    margin-bottom: 10px;
}

.title {
    margin: 0;
}
</style>