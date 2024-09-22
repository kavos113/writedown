<script setup lang="ts">
const props = defineProps<{
  path: string;
}>();

const paths = ["TOP"].concat(
  props.path.split("/").filter((path) => path !== "")
);
const breadcrumbs = paths.map((path, index) => {
  if (path === "TOP") {
    return {
      text: path,
      href: "/",
    };
  }
  return {
    text: path,
    href: "/" + paths.slice(1, index + 1).join("/"),
  };
});
console.log(breadcrumbs);
</script>

<template>
  <div class="breadcrumbs">
    <ul class="itemsWrapper">
      <li
        v-for="breadcrumb in breadcrumbs"
        :key="breadcrumb.text"
        class="items"
      >
        <NuxtLink :to="breadcrumb.href">{{ breadcrumb.text }}</NuxtLink>
        /
      </li>
    </ul>
  </div>
</template>

<style scoped>
.items {
  display: inline;
}

.itemsWrapper {
  list-style-type: none;
  padding: 0;
  margin: 0;
}
</style>
