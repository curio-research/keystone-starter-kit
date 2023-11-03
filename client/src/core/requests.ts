// api requests types

import { api } from 'core/config';
import {NewKeystoneTx} from "./middleware/middleware";
import {WithECDSAAuth} from "./middleware/ecdsaPublicKeyAuth";

export const ECDSAPublicKeyAuthHeader = "ecdsaPublicKeyAuth"

export interface CreatePlayerRequest {
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

export const CreatePlayer = async (request: CreatePlayerRequest) => {
  return api.post('/player', NewKeystoneTx(request, WithECDSAAuth(request)));
};

export const Fire = async (request: CreateProjectileRequest) => {
  return api.post('/fire', NewKeystoneTx(request, WithECDSAAuth(request)));
};

export const Move = async (request: MoveRequest) => {
  return api.post('/move', NewKeystoneTx(request, WithECDSAAuth(request)));
};
