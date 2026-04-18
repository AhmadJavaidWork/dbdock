import {
  createRouter,
  createWebHashHistory,
  RouteLocationNormalizedGeneric,
  RouteRecordRaw,
} from "vue-router";

import ConnectionLayout from "@/layouts/ConnectionLayout.vue";
import WelcomeLayout from "@/layouts/WelcomeLayout.vue";
import { isConnected } from "@/services/connection-manager.service";
import { useConnectionStore } from "@/stores/connection.store";
import { useLoaderStore } from "@/stores/loader.store";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: WelcomeLayout,
    children: [
      {
        path: "/",
        name: "Welcome",
        component: () => import("@/pages/welcome.page.vue"),
      },
    ],
  },
  {
    path: "/connections/:id",
    component: ConnectionLayout,
    children: [
      {
        path: "",
        name: "Connections",
        component: () => import("@/pages/connection.page.vue"),
      },
    ],
    beforeEnter: async function (to: RouteLocationNormalizedGeneric) {
      const connectionId = Number(to.params["id"]);
      if (!connectionId) return { name: "Welcome" };
      const isConnectionConnected = await isConnected(connectionId);
      if (!isConnectionConnected) return { name: "Welcome" };
      for (let i = 0; i < useConnectionStore().activeConnections.length; i++) {
        if (connectionId === useConnectionStore().activeConnections[i].id) {
          useConnectionStore().selectedConnection = useConnectionStore().activeConnections[i];
          break;
        }
      }
      return true;
    },
  },
];

export const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

router.beforeEach(async function () {
  const promises = [];
  useLoaderStore().show();
  try {
    if (useConnectionStore().connections.length === 0) {
      promises.push(useConnectionStore().getConnections());
    }
    if (useConnectionStore().activeConnections.length === 0) {
      promises.push(useConnectionStore().getActiveConnections());
    }
    await Promise.all(promises);
  } catch (error) {}
  useLoaderStore().hide();
});
