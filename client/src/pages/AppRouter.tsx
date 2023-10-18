import { BrowserRouter, Route, Routes } from "react-router-dom";
import TableExplorer from "../components/TableExplorer";

export const AppRouter = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path={"/explore"} element={<TableExplorer />} />
      </Routes>
    </BrowserRouter>
  );
};
