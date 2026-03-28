<script setup lang="ts">
import BaseBorderButton from '@/components/buttons/BaseBorderButton.vue'
import HelloWorld from '@/components/HelloWorld.vue'
import { useThemeStore } from '@/stores/theme.store'
</script>

<template>
  <BaseBorderButton @click="useThemeStore().toggleTheme"
    >Toggle Theme</BaseBorderButton
  >
  <!-- <img id="logo" alt="Wails logo" src="./assets/images/logo.svg" /> -->
  <HelloWorld />
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
