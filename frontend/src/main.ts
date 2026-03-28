import App from '@/App.vue';
import '@/assets/tailwind.css';
import { createPinia } from 'pinia';
import persistedState from 'pinia-plugin-persistedstate';
import { createApp } from 'vue';

const app = createApp(App);

const pinia = createPinia();
pinia.use(persistedState);

app.use(pinia);
app.mount('#app');
