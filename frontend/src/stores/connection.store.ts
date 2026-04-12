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
  const activeConnections = ref<DBConnection[]>([]);
  const selectedConnection = ref<DBConnection | null>(null);

  async function getConnections(): Promise<void> {
    const res = await connectionService.getConnections();
    if (res !== null) {
      connections.value = res;
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
    selectedConnection.value = res.connection;
    return res.message;
  }

  async function deleteConnection(id: number): Promise<void> {
    await connectionService.deleteConnection(id);
    connections.value = connections.value.filter((c) => c.id !== id);
  }

  async function connectToDatabase(conn: DBConnection): Promise<string> {
    const res = await connectionService.connectToDatabase(conn);
    activeConnections.value.push(conn);
    for (let i = 0; i < connections.value.length; i++) {
      if (conn.id === connections.value[i].id) {
        connections.value[i].isConnected = true;
      }
    }
    return res;
  }

  async function disconnectFromDatabase(id: number): Promise<string> {
    const res = await connectionService.disconnectFromDatabase(id);
    activeConnections.value = activeConnections.value.filter((c) => c.id !== id);
    for (let i = 0; i < connections.value.length; i++) {
      if (id === connections.value[i].id) {
        connections.value[i].isConnected = false;
      }
    }
    return res;
  }

  async function getActiveConnections(): Promise<void> {
    const res = await connectionService.getActiveConnections();
    if (res !== null) {
      activeConnections.value = res;
    }
  }

  return {
    connections,
    activeConnections,
    selectedConnection,
    getConnections,
    createConnection,
    updateConnection,
    deleteConnection,
    connectToDatabase,
    disconnectFromDatabase,
    getActiveConnections,
  };
});
