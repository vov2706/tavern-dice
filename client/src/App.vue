<script setup lang="ts">
import ToastsHost from "./components/ToastsHost.vue";
</script>

<template>
  <div class="app-root">
    <RouterView />
    <ToastsHost />
  </div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
.app-root {
  min-height: 100dvh;
}
</style>
