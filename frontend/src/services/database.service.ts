import { DatabaseDriver } from "@/types/database.types";
import { GetSupportedDatabases } from "~/wailsjs/go/main/App";

export function fetchDatabases(): Promise<DatabaseDriver[]> {
  return GetSupportedDatabases();
}
