import ReactDOM from 'react-dom/client';
import { AppRouter } from './pages/AppRouter';
import { ChakraProvider } from '@chakra-ui/react';
import { WorldState } from './store/stateStore';
import { UIState } from './store/uiStore';
import { makeAutoObservable } from 'mobx';

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);

export const worldState = makeAutoObservable(new WorldState());
export const uiState = makeAutoObservable(new UIState());

root.render(
  <ChakraProvider>
    <AppRouter />
  </ChakraProvider>
);
