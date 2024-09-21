<script setup lang="ts">
</script>

<template>
  <div>
    <h1>Page</h1>
  </div>
  <main class="main">
    <SideBarLeft class="sideBarLeft"/>
    <ContentView class="contentView" :page-id="1"/>
    <SideBarRight class="sideBarRight"/>
  </main>
</template>

<style scoped>
.main {
    display: grid;
    grid-template-columns: 1fr 4fr 1fr;
}

.sideBarLeft {
    background-color: aqua;
}

.contentView {
    background-color: bisque;
}

.sideBarRight {
    background-color: aquamarine;
}
</style>