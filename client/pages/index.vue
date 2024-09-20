<script
  setup
  lang="ts"
>

</script>

<template>
  <div>
    <h1>hellp</h1>
  </div>
  <main class="main">
    <SideBarLeft/>
    <ContentView :page-id="1"/>
    <SideBarRight/>
  </main>
</template>

<style scoped>
.main {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
}
</style>