import { DBConnection } from "@/types/connection.type";
import { CreateConnection, TestConnection } from "~/wailsjs/go/main/App";
import { models } from "~/wailsjs/go/models";

function toWailsConnection(conn: DBConnection): models.DBConnection {
  return models.DBConnection.createFrom(conn);
}

export function testConnection(conn: DBConnection): Promise<string> {
  return TestConnection(toWailsConnection(conn));
}

export function createConnection(conn: DBConnection): Promise<string> {
  return CreateConnection(toWailsConnection(conn));
}
