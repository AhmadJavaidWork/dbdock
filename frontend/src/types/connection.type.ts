import { DatabaseDriver } from "@/types/databaseDriver.types";

export type DBConnectionBase = {
  name: string;
  databaseDriverId: number | null;
  host: string;
  port: number | null;
  username: string;
  password: string;
  databaseName: string;
};

export type CreateDBConnection = DBConnectionBase;

export type DBConnectionEntity = DBConnectionBase & {
  id: number;
  lastUsedAt?: string | null;
  createdAt: string;
  updatedAt: string;
};

export type DBConnection = DBConnectionEntity & {
  databaseDriver: DatabaseDriver;
};

export type SaveDBConnectionResponse = {
  connection: DBConnection;
  message: string;
};
