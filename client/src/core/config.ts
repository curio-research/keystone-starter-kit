import axios from 'axios';

// configs
export const KeystoneWebsocketUrl = 'ws://localhost:9001';
export const KeystoneServerUrl = 'http://localhost:9000';

export const api = axios.create({
  baseURL: KeystoneServerUrl, // Replace with your API's base URL
});
