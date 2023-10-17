import React from 'react';
import logo from './logo.svg';
import './App.css';
import WS, {WebSocket} from "ws";
import {addUpdate, TableUpdate} from "../store/store";
import {useDispatch} from "react-redux";

function App() {
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
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
