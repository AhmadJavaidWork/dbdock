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
  <main class="flex flex-col items-center justify-center mt-[10px] gap-[10px]">
    <div id="result" class="result">{{ greeting }}</div>
    <div id="input" class="flex">
      <input
        id="name"
        v-model="name"
        autocomplete="off"
        class="input"
        type="text"
      />
      <button class="btn" @click="greet">Greet</button>
    </div>
  </main>
</template>
