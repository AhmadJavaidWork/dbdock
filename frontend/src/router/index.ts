import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

import EmptyLayout from "@/layouts/EmptyLayout.vue";
import ConnectionPage from "@/pages/connection.page.vue";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: EmptyLayout,
    children: [
      {
        path: "",
        name: "connection",
        component: ConnectionPage,
      },
    ],
  },
];

export const router = createRouter({
  history: createWebHashHistory(), // ✅ best for Wails
  routes,
});
