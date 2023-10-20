import { enableMapSet } from "immer";
import { configureStore, createSlice, getDefaultMiddleware, PayloadAction } from "@reduxjs/toolkit";
import { AccessorsMap } from "../core/schemas";
import _ from "lodash";

enableMapSet();

export enum TableOperationType {
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

const customizedMiddleware = getDefaultMiddleware({
  serializableCheck: false,
});

// ---------------------------
// store state structure
// ---------------------------

export interface StoreState {
  uiControls: {
    selectedTableDisplay: string | null;
  };
  isFetchingState: boolean;
  pendingTableUpdatesToInsert: TableUpdate[];
  tableState: WorldType;
}

export const InitializeState = (): StoreState => {
  return {
    uiControls: {
      selectedTableDisplay: null,
    },
    isFetchingState: false,
    pendingTableUpdatesToInsert: [],
    tableState: new Map<string, Map<number, any>>(),
  };
};

const slice = createSlice({
  name: "world",
  initialState: InitializeState(),
  reducers: {
    // apply table update to state
    addUpdate: function (state: StoreState, action: PayloadAction<TableUpdate>) {
      const payload = action.payload;
      const op = payload.op;
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

    // is it syncing backend state
    setIsFetchingState: function (state: StoreState, action: PayloadAction<boolean>) {
      state.isFetchingState = action.payload;
    },

    // adds a table update to array of pending updates
    addTableUpdateToPendingUpdates: function (state: StoreState, action: PayloadAction<TableUpdate>) {
      const arr = _.clone(state.pendingTableUpdatesToInsert);
      arr.push(action.payload);

      state.pendingTableUpdatesToInsert = arr;
    },

    // apply all pending state updates
    // TODO: filter by tick
    applyAllPendingUpdates: function (state: StoreState, action: PayloadAction<number>) {
      for (const update of state.pendingTableUpdatesToInsert) {
        slice.caseReducers.addUpdate(state, { payload: update, type: "addUpdate" });
      }
      state.pendingTableUpdatesToInsert = [];
    },
  },
});

export type TableType<T> = Map<number, T>;
export type WorldType = Map<string, TableType<any>>;

export const store = configureStore({
  reducer: slice.reducer,
  middleware: customizedMiddleware,
});

export const { addUpdate, setSelectedTableDisplay, setIsFetchingState, addTableUpdateToPendingUpdates, applyAllPendingUpdates } = slice.actions;
