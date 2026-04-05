import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

import NewConnectionLayout from "@/layouts/NewConnectionLayout.vue";
import ConnectionPage from "@/pages/connection.page.vue";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: NewConnectionLayout,
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
  history: createWebHashHistory(),
  routes,
});
