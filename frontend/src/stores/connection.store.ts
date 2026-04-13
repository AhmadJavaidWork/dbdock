import connectionManagerService from "@/services/connection-manager.service";
import connectionService from "@/services/connection.service";
import {
  ConnectionWithStatus,
  CreateDBConnection,
  DBConnection,
  SaveDBConnectionResponse,
} from "@/types/connection.type";
import { Table } from "@/types/table.types";
import { defineStore } from "pinia";
import { computed, ref } from "vue";

export const useConnectionStore = defineStore("connection", () => {
  const connections = ref<ConnectionWithStatus[]>([]);
  const activeConnections = ref<ConnectionWithStatus[]>([]);
  const selectedConnection = ref<ConnectionWithStatus | null>(null);
  const connectionsTables = ref<
    Record<string, { tables: Table[]; schemas: Set<string | undefined> }>
  >({});

  const selectedConnectionTables = computed(function () {
    if (!selectedConnection.value) return [];
    if (!connectionsTables.value[selectedConnection.value.id]) return [];
    return connectionsTables.value[selectedConnection.value.id].tables;
  });

  const selectedConnectionSchemas = computed(function () {
    if (!selectedConnection.value) return [];
    if (!connectionsTables.value[selectedConnection.value.id]) return [];
    return connectionsTables.value[selectedConnection.value.id].schemas;
  });

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
    activeConnections.value = activeConnections.value.filter((c) => c.id !== conn.id);
    selectedConnection.value = res.connection;
    return res.message;
  }

  async function deleteConnection(id: number): Promise<void> {
    await connectionService.deleteConnection(id);
    connections.value = connections.value.filter((c) => c.id !== id);
    activeConnections.value = activeConnections.value.filter((c) => c.id !== id);
  }

  async function connectToDatabase(conn: ConnectionWithStatus): Promise<string> {
    const res = await connectionManagerService.connectToDatabase(conn);
    activeConnections.value.push(conn);
    for (let i = 0; i < connections.value.length; i++) {
      if (conn.id === connections.value[i].id) {
        connections.value[i].isConnected = true;
      }
    }
    return res;
  }

  async function disconnectFromDatabase(id: number): Promise<string> {
    const res = await connectionManagerService.disconnectFromDatabase(id);
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

  async function getTables(): Promise<void> {
    try {
      if (!selectedConnection.value) return;

      const res = await connectionManagerService.listTables(
        selectedConnection.value.id,
        selectedConnection.value.databaseDriver.name
      );
      connectionsTables.value[selectedConnection.value.id] = {
        tables: res,
        schemas: new Set<string | undefined>(),
      };
      for (let i = 0; i < res.length; i++) {
        connectionsTables.value[selectedConnection.value.id].schemas.add(res[i].schema);
      }
    } catch (error: unknown) {
      if (typeof error === "string" && error === "connection not found") {
      }
      console.log("error", error);
    }
  }

  return {
    connections,
    activeConnections,
    selectedConnection,
    connectionsTables,
    selectedConnectionTables,
    selectedConnectionSchemas,
    getConnections,
    createConnection,
    updateConnection,
    deleteConnection,
    connectToDatabase,
    disconnectFromDatabase,
    getActiveConnections,
    getTables,
  };
});
