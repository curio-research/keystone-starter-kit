import { BrowserRouter, Route, Routes } from 'react-router-dom';
import TableExplorer from '../components/TableExplorer';
import Game from './Game';
import { Box, Text } from '@chakra-ui/react';
import { GetStateResponse, TableOperationType, TableUpdate } from '../store/types';
import { stateStore } from '..';
import { useEffect } from 'react';
import { KeystoneWebsocketUrl, api } from 'core/config';

export const AppRouter = () => {
  const startup = async () => {
    const ws = new WebSocket(`${KeystoneWebsocketUrl}/subscribeAllTableUpdates`);

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
    const res = await api.post('/getState', {});

    const data = res.data as GetStateResponse;

    for (const table of data.tables) {
      for (const value of table.values) {
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
