export type Order = "ASC" | "DESC";

export type Result<T> = {
  result: T;
  total: number;
};
