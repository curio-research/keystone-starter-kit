import {GameSchema, GameTable, PlayerSchema, PlayerTable} from "./schemas";
import {worldState} from "../index";
import {gameConst, playerIdTag, privateTag, testPlayerId} from "./config";
import {ethers} from "ethers/lib.esm";
import {CreatePlayer} from "./requests";


export function createPlayer() {
    const playerID = getPlayerID()
    if (playerID === null) {
        const playerWallet = ethers.Wallet.createRandom();
        const newPlayerID = testPlayerId; // TODO do random
        CreatePlayer({PublicKey: playerWallet.publicKey, PlayerId: newPlayerID})

        setPlayerID(newPlayerID.toString())
        setPrivateKey(playerWallet.privateKey)
    }
}

export function getPlayer(): PlayerSchema | undefined {
    const game = getGame()
    if (game === undefined) {
        return undefined
    }

    const playerTag = playerIdTag(game.GameId)
    const playerIDStr = localStorage.getItem(playerTag);
    if (playerIDStr === null) {
        return undefined
    }

    console.log("Player ID", playerIDStr);

    const playerID = parseInt(playerIDStr, 10)
    const player = PlayerTable.filter(worldState.tableState)
        .WithCondition(p => p.PlayerId === playerID)
        .Execute();
    if (player.length === 0) {
        return undefined
    }

    return player[0]
}

export function getPrivateKey(): string {
    return localStorage.getItem(privateTag)!;
}

export function getPlayerID(): string | undefined {
    const game = getGame()
    if (game === undefined) {
        return undefined
    }

    const playerTag = playerIdTag(game.GameId);
    const playerID = localStorage.getItem(playerTag);
    if (playerID) {
        return playerID!;
    }

    return undefined;
}

function setPlayerID(id: string) {
    const game = getGame()
    if (game === undefined) {
        return undefined
    }

    const playerTag = playerIdTag(game.GameId);
    localStorage.setItem(playerTag, id);
}

function setPrivateKey(privateKey: string) {
    localStorage.setItem(privateTag, privateKey);
}

function getGame(): GameSchema | undefined {
    return GameTable.get(worldState.tableState, gameConst)
}