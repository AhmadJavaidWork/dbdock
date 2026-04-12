import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

import WelcomeLayout from "@/layouts/WelcomeLayout.vue";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: WelcomeLayout,
    children: [
      {
        path: "/",
        name: "Welcome",
        component: import("@/pages/welcome.page.vue"),
      },
    ],
  },
  {
    path: "/connections/:id",
    name: "Connections",
    component: import("@/pages/connected.page.vue"),
  },
];

export const router = createRouter({
  history: createWebHashHistory(),
  routes,
});
