import { DBConnection } from "@/types/connection.type";
import { TestConnection } from "~/wailsjs/go/main/App";

export function testConnection(conn: DBConnection): Promise<string> {
  return TestConnection(conn);
}
