import { Theme } from "@/types/theme.type";
import { defineStore } from "pinia";
import { ref } from "vue";
import { SetTheme } from "~/wailsjs/go/main/App";

export const useThemeStore = defineStore(
  "theme",
  () => {
    const theme = ref<Theme>("system");

    function applyTheme(): void {
      const html = document.documentElement;

      if (theme.value === "dark") {
        html.classList.add("dark");
      } else {
        html.classList.remove("dark");
      }

      SetTheme(theme.value === "dark");
    }

    function toggleTheme(): void {
      if (theme.value === "dark") {
        theme.value = "light";
      } else if (theme.value === "light") {
        theme.value = "dark";
      } else {
        const prefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;

        if (prefersDark) {
          theme.value = "dark";
        } else {
          theme.value = "light";
        }
      }
      applyTheme();
    }

    const prefersDark: boolean = window.matchMedia("(prefers-color-scheme: dark)").matches;

    if (prefersDark) {
      theme.value = "light";
    } else {
      theme.value = "dark";
    }
    toggleTheme();

    window.matchMedia("(prefers-color-scheme: dark)").addEventListener("change", () => {
      const prefersDark: boolean = window.matchMedia("(prefers-color-scheme: dark)").matches;

      if (prefersDark) {
        theme.value = "light";
      } else {
        theme.value = "dark";
      }
      toggleTheme();
    });

    return { theme, applyTheme, toggleTheme };
  },
  {
    persist: {
      key: "theme",
      storage: localStorage,
      pick: ["theme"],
      afterHydrate: (ctx) => {
        ctx.store.applyTheme();
      },
    },
  }
);
