import _ from "lodash";
import { AccessorsMap } from "../core/schemas";
import { TableOperationType, TableUpdate, WorldType } from "./types";

// keystone's table state store
export class TableStateStore {
  public isFetchingState: boolean;
  public pendingTableUpdatesToInsert: TableUpdate[];
  public tableState: WorldType;

  constructor() {
    this.isFetchingState = false;
    this.pendingTableUpdatesToInsert = [];
    this.tableState = new Map<string, Map<number, any>>();
  }

  // add update
  public addUpdate(update: TableUpdate) {
    const op = update.op;
    if (op === TableOperationType.Add) {
      return undefined;
    }

    const tableName = update.table;
    const accessor = AccessorsMap.get(tableName);
    if (accessor === undefined) {
      return undefined;
    }

    if (!this.tableState.has(tableName)) {
      this.tableState.set(tableName, new Map<number, any>());
    }

    const table = this.tableState.get(tableName)!;

    const id = update.entity;
    if (op === TableOperationType.Update) {
      accessor.set(table, id, update.value);
    } else if (op === TableOperationType.Remove) {
      accessor.remove(table, id);
    }
  }

  // set is fetching state
  public setIsFetchingState(isFetchingState: boolean) {
    this.isFetchingState = isFetchingState;
  }

  public addTableUpdateToPendingUpdates(update: TableUpdate) {
    const arr = _.clone(this.pendingTableUpdatesToInsert);
    arr.push(update);
    this.pendingTableUpdatesToInsert = arr;
  }

  // apply all pending updates
  public applyAllPendingUpdates() {
    this.pendingTableUpdatesToInsert.forEach((update) => {
      this.addUpdate(update);
    });
    this.pendingTableUpdatesToInsert = [];
  }
}
