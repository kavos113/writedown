<template>
  <div class="root-app-content">
    <h1>WriteDown</h1>
    <NuxtPage class="root-page-view" />
  </div>
</template>
<style>
html, body, #__nuxt {
  margin: 0;
  padding: 0;
}

html, body, #__nuxt {
  height: 100%;
}

.root-app-content {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.root-page-view {
  flex-grow: 1;
}
</style>
