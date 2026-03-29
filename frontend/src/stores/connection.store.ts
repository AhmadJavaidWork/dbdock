import { DBConnection } from "@/types/connection.type";
import { defineStore } from "pinia";
import { ref } from "vue";

export const useConnectionStore = defineStore("connection", () => {
  const connections = ref<DBConnection[]>([]);
  const activeConnection = ref<DBConnection | null>(null);

  return { connections, activeConnection };
});
