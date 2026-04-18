import { Column } from "@/types/column.types";
import { CreateDBConnection, DBConnection } from "@/types/connection.type";
import { Order } from "@/types/query.types";
import { Table } from "@/types/table.types";
import {
  ConnectToDatabase,
  DisconnectFromDatabase,
  GetTableData,
  IsConnected,
  ListTables,
  RunQuery,
} from "~/wailsjs/go/main/App";
import { models } from "~/wailsjs/go/models";

function toWailsConnection(conn: CreateDBConnection | DBConnection): models.DBConnection {
  return models.DBConnection.createFrom(conn);
}

export function connectToDatabase(conn: DBConnection): Promise<string> {
  return ConnectToDatabase(toWailsConnection(conn));
}

export function disconnectFromDatabase(id: number): Promise<string> {
  return DisconnectFromDatabase(id);
}

export function runQuery(id: number, query: string): Promise<Record<string, any>> {
  return RunQuery(id, query);
}

export function listTables(id: number, driver: string): Promise<Table[]> {
  return ListTables(id, driver);
}

export function isConnected(id: number): Promise<boolean> {
  return IsConnected(id);
}

export function getTableData(
  id: number,
  driver: string,
  tableName: string,
  limit: number,
  offset: number,
  orderBy: string,
  order: Order
): Promise<Column[]> {
  return GetTableData(id, driver, tableName, limit, offset, orderBy, order);
}

export default {
  connectToDatabase,
  disconnectFromDatabase,
  runQuery,
  listTables,
  isConnected,
  getTableData,
};
