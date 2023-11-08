import { KeystoneAPI } from 'index';
import {getPlayerID} from "./utils";
import {HeaderEntry, NewKeystoneTx} from "../keystone/middleware";
import {WithEthereumWalletAuth} from "./middleware/ethereumWalletAuth";

const playerIDHeader = "playerIDHeader";
function WithCustomEthereumWalletAuth<T>(req: T): HeaderEntry<any>[] {
    const playerID = getPlayerID()!

    return [
        [playerIDHeader, playerID],
        WithEthereumWalletAuth(req)
    ]
}

// create player
export interface CreatePlayerRequest {
  EthBase64PublicKey: string;
  PlayerId: number;
}

export const CreatePlayer = async (request: CreatePlayerRequest) => {
  return KeystoneAPI.getAPI().post('/createPlayer', NewKeystoneTx(request, ...WithCustomEthereumWalletAuth(request)));
};

// fire projectile
export interface FireRequest {
  Direction: string;
  PlayerId: number;
}

export const Fire = async (request: FireRequest) => {
  return KeystoneAPI.getAPI().post('/fire', NewKeystoneTx(request, ...WithCustomEthereumWalletAuth(request)));
};

// move player
export interface MoveRequest {
  Direction: string;
  PlayerId: number;
}

export const Move = async (request: MoveRequest) => {
  return KeystoneAPI.getAPI().post('/move', NewKeystoneTx(request, ...WithCustomEthereumWalletAuth(request)));
};
