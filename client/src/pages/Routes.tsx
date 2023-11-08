import { BrowserRouter, Route, Routes } from 'react-router-dom';
import TableExplorer from '../components/TableExplorer';
import Game from './Game';
import { Box, Text } from '@chakra-ui/react';
import { useEffect } from 'react';
import { worldState } from 'index';
import { createPlayer } from 'core/utils';

export const AppRouter = () => {
  useEffect(() => {
    const startup = async () => {
      await worldState.connectToKeystone();
      createPlayer();
    };

    startup();
  });

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/explore" element={<TableExplorer />} />
        <Route path="/" element={<Game />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  );
};

const NotFound = () => {
  return (
    <Box p="10">
      <Text fontSize="2xl">Page not found :-/</Text>
    </Box>
  );
};
