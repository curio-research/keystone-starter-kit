import {store, TableType} from "../store/store";


export class TableAccessor<T extends {Id: number}> {
    private tableName: string
    constructor(tableName: string) {
        this.tableName = tableName;
    }

    name(): string {
        return this.tableName;
    }
    get(table: TableType<T>, id: number): T | undefined {
        return table.get(id)
    }

    getAny(table: TableType<T>): T | undefined {
        table.forEach((value) => {
            return value
        })
        return undefined;
    }
    set(table: TableType<T>, id: number, val: T) {
        table.set(id, val)
    }

    filter(table: TableType<T>, val: T): FilterArgs<T> {
        return new FilterArgs(table);
    }

    allEntities(table: TableType<T>): Array<number> {
        return Array.from<number>(table.keys())
    }
}

interface FilterFunction<T> {
    (t: T): boolean
}

class FilterArgs<T> {
    private callbacks: FilterFunction<T>[];
    private table: Map<number, T>;

    constructor(table: Map<number, T>) {
        this.table = table;
        this.callbacks = Array<FilterFunction<T>>();
    }

    WithCondition(f: FilterFunction<T>): FilterArgs<T> {
        this.callbacks.push(f)
        return this
    }

    Execute(): Array<T> {
        const callbacks = this.callbacks;
        const matchingValues = new Array<T>();

        this.table.forEach(function (value: T) {
            for (let cb of callbacks) {
                if (!cb(value)) {
                    return
                }
            }
            matchingValues.push(value)
        })

        return matchingValues;
    }
}
