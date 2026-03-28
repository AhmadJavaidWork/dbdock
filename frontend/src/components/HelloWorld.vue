<script setup lang="ts">
import { ref } from 'vue'
import { Greet } from '~/wailsjs/go/main/App'

const name = ref<string>('')
const greeting = ref<string>('')

function greet() {
  Greet(name.value).then((result) => {
    greeting.value = result
  })
}
</script>

<template>
  <div class="p-6 sans text-lg">
    <h1 class="text-3xl">Welcome to DBDock</h1>
    <p class="mt-4 text-gray-700 font-sans">
      Offline-ready database client with bundled fonts.
    </p>
    <pre class="mt-6 p-4 bg-gray-100 rounded font-mono text-sm">
      SELECT * FROM users WHERE active = 1;
    </pre>
  </div>
</template>
