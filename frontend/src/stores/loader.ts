import { defineStore } from "pinia";
import { ref } from "vue";

const isLoading = ref(false);

export const useLoaderStore = defineStore("loader", () => {
  function show(): void {
    isLoading.value = true;
  }

  function hide(): void {
    isLoading.value = false;
  }

  return { isLoading, show, hide };
});
