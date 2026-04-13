import {
  ConnectionWithStatus,
  CreateDBConnection,
  DBConnection,
  SaveDBConnectionResponse,
} from "@/types/connection.type";
import { DatabaseDriver } from "@/types/databaseDriver.types";
import {
  ConnectToDatabase,
  CreateConnection,
  DeleteConnection,
  DisconnectFromDatabase,
  GetActiveConnections,
  GetConnections,
  TestConnection,
  UpdateConnection,
} from "~/wailsjs/go/main/App";
import { models } from "~/wailsjs/go/models";

function toWailsConnection(conn: CreateDBConnection | DBConnection): models.DBConnection {
  return models.DBConnection.createFrom(conn);
}

function toWailsDatabaseDriver(driver: DatabaseDriver): models.DatabaseDriver {
  return models.DatabaseDriver.createFrom(driver);
}

export function testConnection(conn: CreateDBConnection, driver: DatabaseDriver): Promise<string> {
  return TestConnection(toWailsConnection(conn), toWailsDatabaseDriver(driver));
}

export function createConnection(conn: CreateDBConnection): Promise<SaveDBConnectionResponse> {
  return CreateConnection(toWailsConnection(conn));
}

export async function getConnections(): Promise<ConnectionWithStatus[] | null> {
  return GetConnections();
}

export function updateConnection(conn: DBConnection): Promise<SaveDBConnectionResponse> {
  return UpdateConnection(toWailsConnection(conn));
}

export function deleteConnection(id: number): Promise<void> {
  return DeleteConnection(id);
}

export function connectToDatabase(conn: DBConnection): Promise<string> {
  return ConnectToDatabase(toWailsConnection(conn));
}

export function disconnectFromDatabase(id: number): Promise<string> {
  return DisconnectFromDatabase(id);
}

export function getActiveConnections(): Promise<ConnectionWithStatus[] | null> {
  return GetActiveConnections();
}

export default {
  testConnection,
  createConnection,
  getConnections,
  updateConnection,
  deleteConnection,
  connectToDatabase,
  disconnectFromDatabase,
  getActiveConnections,
};
