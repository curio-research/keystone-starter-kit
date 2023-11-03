// api requests types

import { api } from 'core/config';
i
import {ethers, Signature} from 'ethers';
import {ECDSASignature, ecsign} from 'ethereumjs-util';
import {Buffer} from "buffer";
import {isNil} from "lodash";
import {NewKeystoneTx} from "./middleware";

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
  return api.post('/player', NewKeystoneTx(request, null));
};

export const Fire = async (request: CreateProjectileRequest) => {
  return api.post('/fire', NewKeystoneTx(request));
};

export const Move = async (request: MoveRequest) => {
  return api.post('/move', NewKeystoneTx(request));
};
