import React from "react";
import {TableAccessor} from "../core/tableAccessor";
import {useSelector} from "react-redux";
import {TableType, WorldType} from "../store/store";
import {WithID} from "../core/schemas";


interface TableProps<T extends WithID> {
    table: TableType<T>
    accessor: TableAccessor<T>
}

export default function TableDisplay<T extends {Id: number}> (props: {accessor: TableAccessor<T>}) {
    const accessor = props.accessor;
    const slice = useSelector((state: WorldType) => state.get(accessor.name()))!

    return (
        <React.Fragment>
            <div>Table Name: {accessor.name()}</div>
            <Table table={slice} accessor={accessor} />
        </React.Fragment>
    )
}

function Table<T extends WithID>(props: TableProps<T>) {
    const {table, accessor} = props;

    const anyVal = accessor.getAny(table)
    if (anyVal === undefined) {
        return null
    }

    const columnNames = new Array<string>();
    for (const field in anyVal) {
        columnNames.push(field);
    }

    const allEntities = accessor.allEntities(table)
    return <React.Fragment>
        <table>
            {
                <tr>
                    {
                        columnNames.map((value, index) => {
                            return <th key={index}>{value}</th>
                        })
                    }
                </tr>
            }
            {
                allEntities.map((entity, index) => {
                    const obj = accessor.get(table, entity)! as any
                    return <tr key={index}>
                        {
                            columnNames.map((columnName, index) => {
                                return <td key={index}>{obj[columnName]}</td>
                            })
                        }
                    </tr>
                })
            }
        </table>
    </React.Fragment>
}
