import ReactDOM from 'react-dom/client';
import { AppRouter } from './pages/Routes';
import { ChakraProvider } from '@chakra-ui/react';
import { WorldState } from './keystone/stateStore';
import { UIState } from './core/uiStore';
import { makeAutoObservable } from 'mobx';
import { KeystoneAPIBase } from 'keystone/util';
import { KeystoneServerUrl, KeystoneWebsocketUrl } from 'core/keystoneConfig';
import { PositionWrapperManager } from 'core/positionWrapperManager';

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);

export const positionWrapperState = makeAutoObservable(new PositionWrapperManager());

// Initialize Keystone API object
export const KeystoneAPI = new KeystoneAPIBase(KeystoneServerUrl, KeystoneWebsocketUrl);

// Keystone world state
export const worldState = makeAutoObservable(new WorldState());

// Helper UI state
export const uiState = makeAutoObservable(new UIState());

root.render(
  <ChakraProvider>
    <AppRouter />
  </ChakraProvider>
);
