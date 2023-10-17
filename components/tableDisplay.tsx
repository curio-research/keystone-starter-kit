import React, {Component} from "react";
import PropTypes from "prop-types";
import {TableAccessor} from "../store/accessors";
import {useSelector} from "react-redux";
import {store, StoreType, TableType} from "../store/store";


function TableDisplay<T extends {Id: number}> (props: {accessor: TableAccessor<T>}) {
    const slice = useSelector((state: TableType<T>) => state.get(props.accessor.name()))!
    const divs = new Array<T>()

    slice.forEach(function (value, key, map) {
        const val = props.accessor.get(value.Id)!
        divs.push(val)
    })

    return (
        <React.Fragment>
            {
                divs.map(function (value, index, array) {
                    return <div>{value as any}</div>
                })
            }
        </React.Fragment>
    )
}