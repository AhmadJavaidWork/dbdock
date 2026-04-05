import { DatabaseDriver } from "@/types/databaseDriver.types";
import { GetSupportedDatabases } from "~/wailsjs/go/main/App";

export function fetchDatabaseDrivers(): Promise<DatabaseDriver[]> {
  return GetSupportedDatabases();
}
