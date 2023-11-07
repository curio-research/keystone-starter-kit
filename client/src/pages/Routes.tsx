import { BrowserRouter, Route, Routes } from 'react-router-dom';
import TableExplorer from '../components/TableExplorer';
import Game from './Game';
import { Box, Text } from '@chakra-ui/react';

export const AppRouter = () => {
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
