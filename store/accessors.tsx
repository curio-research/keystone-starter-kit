import {store} from "./store";


export class TableAccessor<T extends {Id: number}> {
    private tableName: string
    constructor(tableName: string) {
        let t: T;
        this.tableName = tableName;
    }

    name(): string {
        return this.tableName;
    }
    get(id: number): T | null {
        const state = store.getState();
        const table = state.get(this.tableName);

        if (!table) {
            return null;
        }
        return table.get(id)
    }

    set(id: number, val: T) {
        const state = store.getState();
        if (!state.get(this.tableName)) {
            state.set(this.tableName, new Map<number, T>());
        }

        const table = state.get(this.tableName)!;
        return table.set(id, val)
    }

    filter(val: T): Array<T> {
        const state = store.getState();
        const table = state.get(this.tableName)!;

        let filter = new FilterArgs<T>(table);
        for (const field in val) {
            if (val[field]) {
                filter = filter.WithCondition((t: T): boolean => {
                        return t[field] === val[field]
                    }
                )
            }
        }

        return filter.Execute()
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
