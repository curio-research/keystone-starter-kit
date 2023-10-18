import {Draft, enableMapSet} from "immer";
import {configureStore, createSlice, getDefaultMiddleware, PayloadAction} from '@reduxjs/toolkit';

import {ToolkitStore} from "@reduxjs/toolkit/dist/configureStore";
import {AccessorsMap, WithID} from "../core/schemas";

enableMapSet()

enum TableOperationType {
    Remove = "removal",
    Update = "set",
    Add = "add"
}

export interface TableUpdate {
    op: TableOperationType,
    entity: number,
    table: string,
    value: any,
    time: Date
}

const customizedMiddleware = getDefaultMiddleware({
    serializableCheck: false
})

const slice = createSlice({
    name: "world",
    initialState: new Map<string, Map<number, any>>(),
    reducers: {
        addUpdate: function (state: Map<string, Map<number, any>>, action: PayloadAction<TableUpdate>) {
            const payload = action.payload;
            let op = payload.op;
            if (op === TableOperationType.Add) {
                return undefined
            }

            const tableName = payload.table;
            const accessor = AccessorsMap.get(tableName);
            if (accessor === undefined) {
                return undefined
            }

            if (!state.has(tableName)) {
                state.set(tableName, new Map<number, any>())
            }

            const table = state.get(tableName)!

            const id = payload.entity;
            if (op === TableOperationType.Update) {
                accessor.set(table, id, payload.value)
            }
            else if (op === TableOperationType.Remove) {
                accessor.remove(table, id)
            }
        },
    }
})

export type TableType<T> = Map<number, T>
export type WorldType = Map<string, TableType<any>>;
export type StoreType = ToolkitStore<WorldType>;

export const store: StoreType = configureStore({
    reducer: slice.reducer,
    middleware: customizedMiddleware,
})

export const {addUpdate} = slice.actions;