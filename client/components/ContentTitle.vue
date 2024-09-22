<script setup lang="ts">
const titleModel = defineModel<string>("title", {
  required: true,
});
const isView = defineModel<boolean>("isView", {
  required: true,
});

const title = ref(titleModel.value);

const update = () => {
  titleModel.value = title.value;
};

defineExpose({
  update,
});
</script>
<template>
  <div class="header">
    <h1
      v-if="isView"
      class="title"
    >
      {{ title }}
    </h1>
    <input
      v-else
      v-model="title"
    />
  </div>
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
