import { KeystoneAPI } from 'index';
import { NewKeystoneTx } from '../keystone/middleware';
import { WithEthereumWalletAuth } from './middleware/ethereumWalletAuth';

import {api, playerIdTag} from 'core/config';
import {HeaderEntry, NewKeystoneTx} from "./middleware/middleware";
import {WithEthereumWalletAuth} from "./middleware/ethereumWalletAuth";
import {getPlayerID} from "./utils";

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
  PublicKey: string;
  PlayerId: number;
}

export const CreatePlayer = async (request: CreatePlayerRequest) => {
  return KeystoneAPI.getAPI().post('/player', NewKeystoneTx(request, WithEthereumWalletAuth(request)));
};

// fire projectile
export interface FireRequest {
  Direction: string;
  PlayerId: number;
}

export const Fire = async (request: FireRequest) => {
  return KeystoneAPI.getAPI().post('/fire', NewKeystoneTx(request, WithEthereumWalletAuth(request)));
};

// move player
export interface MoveRequest {
  Direction: string;
  PlayerId: number;
}

export const Move = async (request: MoveRequest) => {
  return KeystoneAPI.getAPI().post('/move', NewKeystoneTx(request, WithEthereumWalletAuth(request)));
};
