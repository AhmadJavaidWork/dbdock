import App from "@/App.vue";
import "@/assets/tailwind.css";
import clickOutside from "@/directives/clickOutside.directive";
import { router } from "@/router";
import { createPinia } from "pinia";
import persistedState from "pinia-plugin-persistedstate";
import { createApp } from "vue";

const app = createApp(App);

const pinia = createPinia();
pinia.use(persistedState);

app.use(pinia);
app.use(router);

app.directive("click-outside", clickOutside);

app.mount("#app");
