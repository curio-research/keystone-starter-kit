import React from "react";
import ReactDOM from "react-dom/client";
import { Provider } from "react-redux";
import { store } from "./store/store";
import { AppRouter } from "./pages/AppRouter";
import { ChakraProvider } from "@chakra-ui/react";

const root = ReactDOM.createRoot(document.getElementById("root") as HTMLElement);

root.render(
  <Provider store={store}>
    <ChakraProvider>
      <React.StrictMode>
        <AppRouter />
      </React.StrictMode>
    </ChakraProvider>
  </Provider>
);
