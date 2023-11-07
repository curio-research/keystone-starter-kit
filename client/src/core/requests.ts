// api requests types

import {api, playerIdTag} from 'core/config';
import {HeaderEntry, NewKeystoneTx} from "./middleware/middleware";
import {WithEthereumWalletAuth} from "./middleware/ethereumWalletAuth";
import {getPlayerID} from "./utils";

export interface CreatePlayerRequest {
  PublicKey: string;
  PlayerId: number;
}

export interface CreateProjectileRequest {
  Direction: string;
  PlayerId: number;
}

export interface MoveRequest {
  Direction: string;
  PlayerId: number;
}

// api requests

const playerIDHeader = "playerIDHeader";
function WithCustomEthereumWalletAuth<T>(req: T): HeaderEntry<any>[] {
  const playerIDTag = getPlayerID()!

  return [
      [playerIDHeader, playerIDTag],
      WithEthereumWalletAuth(req)
  ]
}

export const CreatePlayer = async (request: CreatePlayerRequest) => {
  return api.post('/player', NewKeystoneTx(request, ...WithCustomEthereumWalletAuth(request)));
};

export const Fire = async (request: CreateProjectileRequest) => {
  return api.post('/fire', NewKeystoneTx(request, ...WithCustomEthereumWalletAuth(request)));
};

export const Move = async (request: MoveRequest) => {
  return api.post('/move', NewKeystoneTx(request, ...WithCustomEthereumWalletAuth(request)));
};
