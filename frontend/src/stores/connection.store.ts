import { useToast } from "@/composables/useToast";
import connectionService from "@/services/connection.service";
import {
  CreateDBConnection,
  DBConnection,
  SaveDBConnectionResponse,
} from "@/types/connection.type";
import { defineStore } from "pinia";
import { ref } from "vue";

export const useConnectionStore = defineStore("connection", () => {
  const connections = ref<DBConnection[]>([]);
  const activeConnection = ref<DBConnection | null>(null);
  const selectedConnection = ref<DBConnection | null>(null);

  async function getConnections(): Promise<void> {
    try {
      const res = await connectionService.getConnections();
      if (res !== null) {
        connections.value = res;
      }
    } catch (error) {
      useToast(error as string, "error");
    }
  }

  async function createConnection(conn: CreateDBConnection): Promise<string> {
    const res: SaveDBConnectionResponse = await connectionService.createConnection(conn);
    connections.value.unshift(res.connection);
    return res.message;
  }

  async function updateConnection(conn: DBConnection): Promise<string> {
    const res: SaveDBConnectionResponse = await connectionService.updateConnection(conn);
    for (let i = 0; i < connections.value.length; i++) {
      if (connections.value[i].id === res.connection.id) {
        connections.value.splice(i, 1);
        connections.value.unshift(res.connection);
        break;
      }
    }
    return res.message;
  }

  return {
    connections,
    activeConnection,
    selectedConnection,
    getConnections,
    createConnection,
    updateConnection,
  };
});
