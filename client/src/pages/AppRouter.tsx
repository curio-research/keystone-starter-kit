import { BrowserRouter, Route, Routes } from 'react-router-dom';
import TableExplorer from '../components/TableExplorer';
import Game from './Game';
import { Box, Text } from '@chakra-ui/react';
import { GetStateResponse, TableOperationType, TableUpdate } from '../store/types';
import { stateStore } from '..';
import axios from 'axios';
import { useEffect } from 'react';

export const AppRouter = () => {
  const startup = async () => {
    // TODO: this is being pinged twice for some reason
    // TODO: move this to init file
    const ws = new WebSocket('ws://localhost:9001/subscribeAllTableUpdates');

    ws.onopen = () => {
      console.log('connection opened!');
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
    const url = 'http://localhost:9000/getState';
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
    <BrowserRouter>
      <Routes>
        <Route path="/explore" element={<TableExplorer />} />
        <Route path="/" element={<Game />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </BrowserRouter>
  );
};

const NotFoundPage = () => {
  return (
    <Box p="10">
      <Text fontSize="2xl">Page not found :-/</Text>
    </Box>
  );
};
