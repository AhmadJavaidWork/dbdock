import App from "@/App.vue";
import "@/assets/global.scss";
import "@/assets/tailwind.css";
import clickOutside from "@/directives/clickOutside.directive";
import { router } from "@/router";
import { useThemeStore } from "@/stores/theme.store";
import { createPinia } from "pinia";
import persistedState from "pinia-plugin-persistedstate";
import { createApp } from "vue";

const app = createApp(App);
const pinia = createPinia();

pinia.use(persistedState);

app.use(pinia);
app.use(router);

await useThemeStore().loadThemeConfig();

app.directive("click-outside", clickOutside);

app.mount("#app");
