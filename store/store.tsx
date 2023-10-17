import {configureStore, createSlice, PayloadAction} from '@reduxjs/toolkit';

import {ToolkitStore} from "@reduxjs/toolkit/dist/configureStore";

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


const slice = createSlice({
    name: "world",
    initialState: new Map<string, Map<number, any>>(),
    reducers: {
        addUpdate: function (state, action: PayloadAction<TableUpdate>) {
            const payload = action.payload;
            const tableName = payload.table;
            if (!state.has(tableName)) {
                state.set(tableName, new Map<any, null>())
            }

            const table = state.get(tableName)!
            const op = payload.op;
            let id = payload.value.id;

            if (op === TableOperationType.Update) {
                table.set(id, payload)
            }
            else if (op === TableOperationType.Remove) {
                table.delete(id)
            }
        },
    }
})


export type TableType<T> = Map<string, Map<number, T>>;
export type StoreType = ToolkitStore<TableType<any>>;
export const store: StoreType = configureStore({
    reducer: slice.reducer
})

export const {addUpdate} = slice.actions;