// api requests types

import { api } from 'core/config';

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
  return api.post('/player', request);
};

export const Fire = async (request: CreateProjectileRequest) => {
  return api.post('/fire', request);
};

export const Move = async (request: MoveRequest) => {
  return api.post('/move', request);
};
