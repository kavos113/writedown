<script setup lang="ts"></script>

<template>
  <main class="main">
    <SideBarLeft class="sideBarLeft" />
    <Content :page-id="1" />
  </main>
</template>

<style scoped>
.main {
  display: grid;
  grid-template-columns: 1fr 4fr 1fr;
}
</style>
