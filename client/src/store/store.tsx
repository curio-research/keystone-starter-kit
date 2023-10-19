import { enableMapSet } from "immer";
import { configureStore, createSlice, getDefaultMiddleware, PayloadAction } from "@reduxjs/toolkit";
import { ToolkitStore } from "@reduxjs/toolkit/dist/configureStore";
import { AccessorsMap } from "../core/schemas";

enableMapSet();

enum TableOperationType {
  Remove = "removal",
  Update = "set",
  Add = "add",
}

export interface TableUpdate {
  op: TableOperationType;
  entity: number;
  table: string;
  value: any;
  time: Date;
}

const customizedMiddleware = getDefaultMiddleware({
  serializableCheck: false,
});

export interface StoreState {
  uiControls: {
    selectedTableDisplay: string | null;
  };
  tableState: Map<string, Map<number, any>>;
}

export const InitializeState = (): StoreState => {
  return {
    uiControls: {
      selectedTableDisplay: null,
    },
    tableState: new Map<string, Map<number, any>>(),
  };
};

const slice = createSlice({
  name: "world",
  initialState: InitializeState(),
  reducers: {
    addUpdate: function (state: StoreState, action: PayloadAction<TableUpdate>) {
      const payload = action.payload;
      let op = payload.op;
      if (op === TableOperationType.Add) {
        return undefined;
      }

      const tableName = payload.table;
      const accessor = AccessorsMap.get(tableName);
      if (accessor === undefined) {
        return undefined;
      }

      if (!state.tableState.has(tableName)) {
        state.tableState.set(tableName, new Map<number, any>());
      }

      const table = state.tableState.get(tableName)!;

      const id = payload.entity;
      if (op === TableOperationType.Update) {
        accessor.set(table, id, payload.value);
      } else if (op === TableOperationType.Remove) {
        accessor.remove(table, id);
      }
    },

    setSelectedTableDisplay: function (state: StoreState, action: PayloadAction<string>) {
      state.uiControls.selectedTableDisplay = action.payload;
    },
  },
});

export type TableType<T> = Map<number, T>;
export type WorldType = Map<string, TableType<any>>;
export type StoreType = ToolkitStore<StoreState>;

export const store: StoreType = configureStore({
  reducer: slice.reducer,
  middleware: customizedMiddleware,
});

export const { addUpdate, setSelectedTableDisplay } = slice.actions;
