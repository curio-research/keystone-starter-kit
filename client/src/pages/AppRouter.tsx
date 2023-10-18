import { BrowserRouter, Route, Routes } from "react-router-dom";
import TableExplorer from "../components/TableExplorer";
import { Box, Text } from "@chakra-ui/react";

export const AppRouter = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/explore" element={<TableExplorer />} />
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
