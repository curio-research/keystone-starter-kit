import { useEffect } from "react";
import { Box, Select } from "@chakra-ui/react";
// eslint-disable-next-line max-len
// import { addTableUpdateToPendingUpdates, addUpdate, applyAllPendingUpdates, GetStateResponse, setIsFetchingState, setSelectedTableDisplay, store, StoreState, TableOperationType, TableUpdate } from "../store/store";
import { Accessors } from "../core/schemas";
import { observer } from "mobx-react";
import Table from "./Table";
import axios from "axios";
import { GetStateResponse, TableOperationType, TableUpdate } from "../store/types";
import { stateStore, uiStore } from "..";

const TableExplorer = observer(() => {
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
        if (stateStore.isFetchingState) {
          stateStore.addTableUpdateToPendingUpdates(update);
        } else {
          stateStore.addUpdate(update);
        }
      }
    };

    ws.onerror = (event: Event) => {
      console.log(event);
    };

    stateStore.setIsFetchingState(true);

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

        stateStore.addUpdate(tableUpdate);
      }
    }

    stateStore.applyAllPendingUpdates();
    stateStore.setIsFetchingState(false);
  };

  useEffect(() => {
    startup();
  }, []);

  return (
    <Box m={10}>
      <Box w="200px" mb={10}>
        <Select
          value={uiStore.selectedTableToDisplay || ""}
          placeholder="Select table"
          onChange={(e) => {
            uiStore.setSelectedTableToDisplay(e.target.value);
          }}
        >
          {Accessors.map((accessor) => {
            return (
              <option value={accessor.name()} key={accessor.name()}>
                {accessor.name()}
              </option>
            );
          })}
        </Select>
      </Box>

      {Accessors.map((accessor, index) => {
        return uiStore.selectedTableToDisplay === accessor.name() && <Table key={index} accessor={accessor} />;
      })}
    </Box>
  );
});

export default TableExplorer;
