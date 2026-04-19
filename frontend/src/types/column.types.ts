export type Column = {
  name: string;
  rows: any[];
};

export type DisplayColumn = Column & {
  width: number;
  isResizing: boolean;
  startX: number;
  startWidth: number;
  frameId: number | null;
  pendingWidth: number;
};
