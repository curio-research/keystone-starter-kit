import "reflect-metadata"
import {TableAccessor} from "./tableAccessor";

export interface WithID {
    Id: number
}

interface Position {
    x: number;
    y: number;
}

interface LargeTileSchema {
    Pos: Position
    OwnerId: number
    Level: number
    Visible: boolean
    Locked: boolean
    Id: number
}

const LargeTileTable = new TableAccessor<LargeTileSchema>("LargeTileSchema");

interface SmallTileSchema {
    Pos: Position
    Id: number
}

const SmallTileTable = new TableAccessor<SmallTileSchema>("SmallTileSchema");

type TroopType = string
type Layer = string

interface TroopStackSchema {
    TroopType: TroopType
    Layer: Layer
    Amount: number
    Pos: Position
    OwnerId: number
    MovementStamina: number
    Moving: boolean
    IsGuarding: boolean
    LoadedStackId: number
    Id: number
}

const TroopStackTable = new TableAccessor<TroopStackSchema>("TroopStackSchema");

interface PlayerSchema {
    Position: Position
    Resources: number
    PlayerID: number

    Id: number
}

const PlayerTable = new TableAccessor<PlayerSchema>("PlayerSchema");

interface CardSchema {
    OwnerID: number
    Number: number

    Id: number
}

const CardTable = new TableAccessor<CardSchema>("CardSchema")

export const Accessors: TableAccessor<any>[] = [
    LargeTileTable,
    SmallTileTable,
    TroopStackTable,
    PlayerTable,
    CardTable
]

export const AccessorsMap = new Map<string, TableAccessor<any>>()
for (const accessor of Accessors) {
    AccessorsMap.set(accessor.name(), accessor);
}