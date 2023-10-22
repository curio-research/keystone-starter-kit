import React from 'react';
import ReactDOM from 'react-dom/client';
import { AppRouter } from './pages/AppRouter';
import { ChakraProvider } from '@chakra-ui/react';
import { TableStateStore } from './store/stateStore';
import { UIStore } from './store/uiStore';
import { makeAutoObservable } from 'mobx';

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);

export const stateStore = makeAutoObservable(new TableStateStore());
export const uiStore = makeAutoObservable(new UIStore());

root.render(
  <ChakraProvider>
    <AppRouter />
  </ChakraProvider>
);
