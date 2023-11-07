import 'reflect-metadata';
import { TableAccessor } from './tableAccessor';

// ------------------------
// game schemas
// TODO: add automatic code-generation for this step
// ------------------------
interface LocalRandSeedSchema {
  RandValue: number;
  Id: number;
}

export enum Weather {
  Sunny = 1,
  Windy = 2,
}

export interface GameSchema {
  Id: number;
  Weather: Weather;
  GameId: string;
}

export interface TileSchema {
  Id: number;
  Position: Position;
  Terrain: boolean; // true: ground, false: obstacle
}

export interface PlayerSchema {
  Id: number;
  Position: Position;
  Resources: number;
  PlayerId: number;
  PublicKey: string;
}

export interface ProjectileSchema {
  Id: number;
  Position: Position;
}

export interface AnimalSchema {
  Id: number;
  Position: Position;
}

export interface ResourceSchema {
  Id: number;
  Position: Position;
  Amount: number;
}

export interface Position {
  x: number;
  y: number;
}

// ---------------------------
// table accessors
// ---------------------------

export const PlayerTable = new TableAccessor<PlayerSchema>('PlayerSchema');
export const TileTable = new TableAccessor<TileSchema>('TileSchema');
export const GameTable = new TableAccessor<GameSchema>('GameSchema');
export const AnimalTable = new TableAccessor<AnimalSchema>('AnimalSchema');
export const ProjectileTable = new TableAccessor<ProjectileSchema>('ProjectileSchema');
export const LocalRandSeedTable = new TableAccessor<LocalRandSeedSchema>('LocalRandSeedSchema');
export const ResourceTable = new TableAccessor<ResourceSchema>('ResourceSchema');

export const AllTableAccessors: TableAccessor<any>[] = [
  PlayerTable,
  TileTable,
  GameTable,
  AnimalTable,
  ProjectileTable,
  LocalRandSeedTable,
  ResourceTable,
];
