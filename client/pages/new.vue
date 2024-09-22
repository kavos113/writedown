<script setup lang="ts">
const state = window.history.state;

const parentId = state?.parentId as number;
const parentPath = state?.parentPath as string;
</script>

<template>
  <main class="main">
    <SideBarLeft class="sideBarLeft" />
    <NewContent :parent-id="parentId" :parent-path="parentPath" />
  </main>
</template>

<style scoped>
.main {
  display: grid;
  grid-template-columns: 1fr 4fr 1fr;
}
</style>
