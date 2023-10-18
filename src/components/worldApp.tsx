import React from 'react';

import {addUpdate, TableUpdate} from "../store/store";
import {useDispatch} from "react-redux";
import {Accessors, AccessorsMap} from "../core/schemas";
import Table from "./table";


function WorldApp() {
  const dispatch = useDispatch();
  const ws = new window.WebSocket("ws://localhost:9001/subscribeAllTableUpdates");

  ws.onopen = () => {
    console.log("connection opened!")
  }

  ws.onmessage = (event: MessageEvent) => {
    const jsonObj: any = JSON.parse(event.data)
    const updates = jsonObj as unknown as Array<TableUpdate>;
    for (const update of updates) {
      dispatch(addUpdate({
        entity: update.entity,
        op: update.op,
        table: update.table,
        time: update.time,
        value: update.value
      }))
    }
  }

  ws.onerror = (event: Event) => {
    console.log(event)
  }

  return (
      <React.Fragment>
        {
          Accessors.map((accessor, index) => {
            return <Table key={index} accessor={accessor} />
          })
        }
      </React.Fragment>
  );
}

export default WorldApp;
