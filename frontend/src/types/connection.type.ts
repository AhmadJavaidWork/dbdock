import { DatabaseDriver } from "@/types/database.types";

export type DBConnection = {
  name: string;
  type: DatabaseDriver;
  host: string;
  port: number;
  username: string;
  password: string;
  database: string;
};
