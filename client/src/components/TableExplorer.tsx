import React, { useEffect } from "react";
// eslint-disable-next-line max-len
import { addTableUpdateToPendingUpdates, addUpdate, applyAllPendingUpdates, GetStateResponse, setIsFetchingState, setSelectedTableDisplay, store, StoreState, TableOperationType, TableUpdate } from "../store/store";
import { Accessors } from "../core/schemas";
import Table from "./table";
import { Box } from "@chakra-ui/react";
import { Select } from "@chakra-ui/react";
import { useSelector } from "react-redux";
import axios from "axios";

function TableExplorer() {
  const uiControls = useSelector((state: StoreState) => state.uiControls);

  const startup = async () => {
    // TODO: this is being pinged twice for some reason
    // TODO: move this to init file
    const ws = new WebSocket("ws://localhost:9001/subscribeAllTableUpdates");

    ws.onopen = () => {
      console.log("connection opened!");
    };

    ws.onmessage = (event: MessageEvent) => {
      const jsonObj: any = JSON.parse(event.data);
      const updates = jsonObj as TableUpdate[];

      for (const update of updates) {
        if (store.getState().isFetchingState) {
          store.dispatch(addTableUpdateToPendingUpdates(update));
        } else {
          store.dispatch(addUpdate(update));
        }
      }
    };

    ws.onerror = (event: Event) => {
      console.log(event);
    };

    store.dispatch(setIsFetchingState(true));

    // call api
    const url = "http://localhost:9000/getState";
    const res = await axios.post(url, {});

    const data = res.data as GetStateResponse;

    for (const table of data.tables) {
      for (const value of table.values) {
        // TODO: fix time
        const date = new Date();

        const tableUpdate: TableUpdate = {
          op: TableOperationType.Update,
          entity: value.entity,
          table: table.name,
          value: value.value,
          time: date,
        };

        store.dispatch(addUpdate(tableUpdate));
      }
    }

    store.dispatch(applyAllPendingUpdates(0));

    store.dispatch(setIsFetchingState(false));
  };

  useEffect(() => {
    startup();
  }, []);

  return (
    <Box m={10}>
      <Box w={"200px"} mb={10}>
        <Select
          value={uiControls.selectedTableDisplay || ""}
          placeholder="Select table"
          onChange={(e) => {
            store.dispatch(setSelectedTableDisplay(e.target.value));
          }}
        >
          {Accessors.map((accessor, index) => {
            return <option value={accessor.name()}>{accessor.name()}</option>;
          })}
        </Select>
      </Box>

      {Accessors.map((accessor, index) => {
        return uiControls.selectedTableDisplay === accessor.name() && <Table key={index} accessor={accessor} />;
      })}
    </Box>
  );
}

export default TableExplorer;
