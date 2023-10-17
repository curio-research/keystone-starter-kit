import React from 'react';
import WS, {WebSocket} from "ws";
import {addUpdate, store, TableUpdate} from "../store/store";
import {Provider, useDispatch} from "react-redux";
import {Accessors} from "../core/schemas";
import Table from "./table";

function WorldApp() {
  const dispatch = useDispatch();
  const ws = new WebSocket("ws://localhost:8080/subscribeAllTableUpdates");

  ws.onopen = () => {
    console.log("connection opened!")
  }

  ws.onmessage = (event: WS.MessageEvent) => {
    const data = event.data as unknown as TableUpdate;
    dispatch(addUpdate({
      entity: data.entity,
      op: data.op,
      table: data.table,
      time: data.time,
      value: data.value
    }))
  }

  ws.onerror = (event: WS.ErrorEvent) => {
    console.log(event.error)
  }

  return (
      <Provider store={store}>
      <React.Fragment>
        {
          Accessors.map((accessor) => {
            return <Table accessor={accessor} />
          })
        }
      </React.Fragment>
      </Provider>
  );
}

export default WorldApp;
