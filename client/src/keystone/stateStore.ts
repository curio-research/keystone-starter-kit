import _ from 'lodash';
import { AllTableAccessors } from '../core/schemas';
import { TableOperationType, TableUpdate, IWorld, GetStateResponse } from './types';
import { TableAccessor } from 'keystone/tableAccessor';
import { KeystoneWebsocketUrl } from 'core/keystoneConfig';
import { KeystoneAPI } from 'index';
import { toast } from 'pages/Game';

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

    this.connectToKeystone();
  }

  private async connectToKeystone() {
    // initialize the websocket connection
    const ws = new WebSocket(`${KeystoneWebsocketUrl}/subscribeAllTableUpdates`);

    ws.onopen = () => {
      console.log('✅ Connected to keystone websocket ');
    };

    ws.onmessage = (event: MessageEvent) => {
      const jsonObj: any = JSON.parse(event.data);
      const updates = jsonObj as TableUpdate[];

      for (const update of updates) {
        if (this.isFetchingState) {
          this.addTableUpdateToPendingUpdates(update);
        } else {
          this.addUpdate(update);
        }
      }
    };

    ws.onerror = (event: Event) => {
      toast.toast({
        description: 'Error connecting to Keystone websocket',
        status: 'error',
        duration: 10_000,
        isClosable: true,
      });

      console.log(event);
    };

    this.setIsFetchingState(true);

    // call api
    const startTime = Date.now();

    const res = await KeystoneAPI.getAPI().post('/getState', {});

    console.log(`✅ Fetched initial state in ${Date.now() - startTime}ms`);

    const data = res.data as GetStateResponse;

    for (const table of data.tables) {
      for (const value of table.values) {
        const date = new Date();

        const tableUpdate: TableUpdate = {
          op: TableOperationType.Update,
          entity: value.entity,
          table: table.name,
          value: value.value,
          time: date,
        };

        this.addUpdate(tableUpdate);
      }
    }

    console.log('✅ Initial state synced ');

    this.applyAllPendingUpdates();
    this.setIsFetchingState(false);
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
