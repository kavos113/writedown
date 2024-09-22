<script setup lang="ts">
const titleModel = defineModel<string>("title", {
  required: true,
});
const isView = defineModel<boolean>("isView", {
  required: true,
});

const title = ref(titleModel.value);
const isTitleView = ref(true);

const titleEdit = () => {
  if (!isView.value) {
    isTitleView.value = false;
  }
};

const update = () => {
  isTitleView.value = true;
  titleModel.value = title.value;
};

defineExpose({
  update,
});
</script>
<template>
  <div
    class="header"
    :class="{ editMode: !isView && isTitleView }"
    @click="titleEdit"
  >
    <h1
      v-if="isTitleView"
      class="title"
    >
      {{ title }}
      <span
        v-if="!isView"
        class="clickToEdit"
        >クリックして編集</span
      >
    </h1>
    <input
      v-else
      v-model="title"
      class="input"
    />
  </div>
</template>
<style scoped>
.header {
  border-bottom: 1px solid #aaa;
  vertical-align: middle;
  border-radius: 5px 5px 0 0;
  transition-property: background-color, border-bottom-color;
  transition-duration: 0.2s;
}

.title {
  margin: 0;
  padding: 5px;
  font-size: 32px;
  font-weight: normal;
}

.editMode:hover {
  cursor: pointer;
  background-color: #eee;
  border-bottom-color: #666;
}

.clickToEdit {
  font-size: 0.6em;
  margin-left: 10px;
  font-weight: normal;
  color: #666;
}

.input {
  box-sizing: border-box;
  width: 100%;
  height: 100%;
  margin: 0;
  border-radius: 5px 5px 0 0;
  border: none;
  font-size: 32px;
  font-family: monospace;
  transition-property: background-color;
  transition-duration: 0.2s;
  padding-left: 5px;
}

.input:focus {
  outline: none;
  background-color: #eee;
}
</style>
