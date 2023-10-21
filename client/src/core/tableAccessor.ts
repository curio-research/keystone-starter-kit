import { ITable, IWorld } from 'store/types';

export class TableAccessor<T extends { Id: number }> {
  private tableName: string;
  constructor(tableName: string) {
    this.tableName = tableName;
  }

  name(): string {
    return this.tableName;
  }

  get(table: ITable<T>, id: number): T | undefined {
    return table.get(id);
  }

  getAll(world: IWorld): Array<T> {
    const table = world.get(this.tableName);

    if (!table) {
      return [];
    }
    return Array.from<T>(table.values());
  }

  getAny(table: ITable<T>): T | undefined {
    const { value } = table.values().next();
    return value;
  }

  set(table: ITable<T>, id: number, val: T) {
    table.set(id, val);
  }

  remove(table: ITable<T>, id: number) {
    table.delete(id);
  }

  filter(table: ITable<T>): FilterArgs<T> {
    return new FilterArgs(table);
  }

  allEntities(table: ITable<T>): Array<number> {
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
