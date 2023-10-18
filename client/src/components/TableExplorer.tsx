import React, { useEffect } from "react";
import { addUpdate, setSelectedTableDisplay, store, StoreState, TableUpdate } from "../store/store";
import { Accessors } from "../core/schemas";
import Table from "./table";
import { Box } from "@chakra-ui/react";
import { Select } from "@chakra-ui/react";
import { useSelector } from "react-redux";

function TableExplorer() {
  const uiControls = useSelector((state: StoreState) => state.uiControls);

  useEffect(() => {
    // TODO: this is being pinged twice for some reason
    // TODO: move this to init file
    const ws = new WebSocket("ws://localhost:9001/subscribeAllTableUpdates");

    ws.onopen = () => {
      console.log("connection opened!");
    };

    ws.onmessage = (event: MessageEvent) => {
      const jsonObj: any = JSON.parse(event.data);
      const updates = jsonObj as unknown as Array<TableUpdate>;
      for (const update of updates) {
        store.dispatch(
          addUpdate({
            entity: update.entity,
            op: update.op,
            table: update.table,
            time: update.time,
            value: update.value,
          })
        );
      }
    };

    ws.onerror = (event: Event) => {
      console.log(event);
    };
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
