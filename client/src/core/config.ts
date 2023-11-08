import axios from 'axios';

// configs
export const KeystoneWebsocketUrl = 'ws://localhost:9001';
export const KeystoneServerUrl = 'http://localhost:9000';

export const api = axios.create({
  baseURL: KeystoneServerUrl, // Replace with your API's base URL
});

// constants
export const gameConst = 200;
export const testPlayerId = -100;

const playerIDTag = "existingPlayerID";
export const base64PublicKeyTag = "base64PublicKey"
export const privateKeyTag = "privateKey"

export function playerIdTag(gameID: string): string {
  return playerIDTag + "_" + gameID
}