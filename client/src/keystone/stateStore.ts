import _ from 'lodash';
import { AllTableAccessors } from '../core/schemas';
import { TableOperationType, TableUpdate, IWorld } from './types';
import { TableAccessor } from 'keystone/tableAccessor';

// keystone's table state store
export class WorldState {
  public isFetchingState: boolean;
  public pendingTableUpdatesToInsert: TableUpdate[];
  public tableState: IWorld;
  public tableAccessors: Map<string, TableAccessor<any>>;

  constructor() {
    this.isFetchingState = false;
    this.pendingTableUpdatesToInsert = [];
    this.tableState = new Map<string, Map<number, any>>();
    this.tableAccessors = new Map<string, TableAccessor<any>>();

    // add table to worlds
    AllTableAccessors.forEach((accessor) => {
      this.tableAccessors.set(accessor.name(), accessor);
    });
  }

  // add update
  public addUpdate(update: TableUpdate) {
    const op = update.op;
    if (op === TableOperationType.Add) {
      return undefined;
    }

    const tableName = update.table;
    const accessor = this.tableAccessors.get(tableName);
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
    console.log('pending table updates to apply: ', this.pendingTableUpdatesToInsert.length);
    this.pendingTableUpdatesToInsert.forEach((update) => {
      this.addUpdate(update);
    });
    this.pendingTableUpdatesToInsert = [];
  }
}
