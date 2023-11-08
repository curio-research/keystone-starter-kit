// configs
export const KeystoneWebsocketUrl = 'ws://localhost:9001';
export const KeystoneServerUrl = 'http://localhost:9000';

// constants
export const gameEntity = 200;
export const testPlayerId = -100;

const playerIDTag = "existingPlayerID";
export const base64PublicKeyTag = "base64PublicKey"
export const privateKeyTag = "privateKey"

export function playerIdTag(gameID: string): string {
    return playerIDTag + "_" + gameID
}