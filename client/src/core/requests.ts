import { api } from 'core/config';
import { NewKeystoneTx } from '../keystone/middleware';
import { WithEthereumWalletAuth } from './middleware/ethereumWalletAuth';

export const ECDSAPublicKeyAuthHeader = 'ecdsaPublicKeyAuth';

// ----------------------------------------
// API requests to Keystone server
// ----------------------------------------

// create player
export interface CreatePlayerRequest {
  PlayerId: number;
}

export const CreatePlayer = async (request: CreatePlayerRequest) => {
  return api.post('/player', NewKeystoneTx(request, WithEthereumWalletAuth(request)));
};

// fire projectile
export interface FireRequest {
  Direction: string;
  PlayerId: number;
}

export const Fire = async (request: FireRequest) => {
  return api.post('/fire', NewKeystoneTx(request, WithEthereumWalletAuth(request)));
};

// move player
export interface MoveRequest {
  Direction: string;
  PlayerId: number;
}

export const Move = async (request: MoveRequest) => {
  return api.post('/move', NewKeystoneTx(request, WithEthereumWalletAuth(request)));
};
