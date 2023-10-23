export type ITable<T> = Map<number, T>;
export type IWorld = Map<string, ITable<any>>;

export interface StoreState {
  uiControls: {
    selectedTableDisplay: string | null;
  };
  isFetchingState: boolean;
  pendingTableUpdatesToInsert: TableUpdate[];
  tableState: IWorld;
}

export interface TableUpdate {
  op: TableOperationType;
  entity: number;
  table: string;
  value: any;
  time: Date;
}

export enum TableOperationType {
  Remove = 'removal',
  Update = 'set',
  Add = 'add',
}

// get state types
export interface GetStateResponse {
  tick: number;
  tables: TableResponse[];
}

export interface TableResponse {
  name: string;
  values: ValueResponse[];
}

export interface ValueResponse {
  entity: number;
  value: any;
}
