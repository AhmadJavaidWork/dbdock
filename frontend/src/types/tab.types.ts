import { DisplayColumn } from "@/types/column.types";
import { Table } from "@/types/table.types";

export type Tab = {
  table: Table;
  columns: DisplayColumn[];
  total: number;
  limit: number;
  offset: number;
  appliedLimit: number;
  appliedOffset: number;
  attached: boolean;
};
