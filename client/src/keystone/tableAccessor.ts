import {IWorld} from 'store/types';

// typed table accessor
export class TableAccessor<T extends { Id: number }> {
  private tableName: string;
  constructor(tableName: string) {
    this.tableName = tableName;
  }

  // table name
  name(): string {
    return this.tableName;
  }

  // get struct from entity
  get(world: IWorld, entity: number): T | undefined {
    const table = world.get(this.tableName);
    if (!table) {
      return undefined;
    }

    return table.get(entity);
  }

  // get all entities
  getAll(world: IWorld): Array<T> {
    const table = world.get(this.tableName);
    if (!table) {
      return [];
    }

    return Array.from<T>(table.values());
  }

  getAny(world: IWorld): T | undefined {
    const table = world.get(this.tableName);
    if (!table) {
      return undefined;
    }

    const { value } = table.values().next();
    return value;
  }

  set(world: IWorld, id: number, val: T) {
    const table = world.get(this.tableName);
    if (!table) {
      return [];
    }

    table.set(id, val);
  }

  remove(world: IWorld, id: number) {
    const table = world.get(this.tableName);
    if (!table) {
      return [];
    }

    table.delete(id);
  }

  filter(world: IWorld): FilterArgs<T> {
    const table = world.get(this.tableName)!;
    return new FilterArgs(table);
  }

  allEntities(world: IWorld): Array<number> {
    const table = world.get(this.tableName);
    if (!table) {
      return [];
    }

    return Array.from<number>(table.keys());
  }
}

interface FilterFunction<T> {
  (t: T): boolean;
}

class FilterArgs<T> {
  private callbacks: FilterFunction<T>[];
  private table: Map<number, T>;

  constructor(table: Map<number, T>) {
    this.table = table;
    this.callbacks = Array<FilterFunction<T>>();
  }

  WithCondition(f: FilterFunction<T>): FilterArgs<T> {
    this.callbacks.push(f);
    return this;
  }

  Execute(): Array<T> {
    const callbacks = this.callbacks;
    const matchingValues = new Array<T>();

    this.table.forEach(function (value: T) {
      for (const cb of callbacks) {
        if (!cb(value)) {
          return;
        }
      }
      matchingValues.push(value);
    });

    return matchingValues;
  }
}

