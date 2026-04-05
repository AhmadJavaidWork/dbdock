import {
  CreateDBConnection,
  DBConnection,
  SaveDBConnectionResponse,
} from "@/types/connection.type";
import { DatabaseDriver } from "@/types/databaseDriver.types";
import {
  CreateConnection,
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

export async function getConnections(): Promise<DBConnection[] | null> {
  return GetConnections();
}

export function updateConnection(conn: DBConnection): Promise<SaveDBConnectionResponse> {
  return UpdateConnection(toWailsConnection(conn));
}

export default {
  testConnection,
  createConnection,
  getConnections,
  updateConnection,
};
