import {GameSchema, GameTable, PlayerSchema, PlayerTable} from "./schemas";
import {worldState} from "../index";
import {ethers} from "ethers";
import {CreatePlayer} from "./requests";
import sjcl from "sjcl";
import {base64PublicKeyTag, gameEntity, playerIdTag, privateKeyTag, testPlayerId} from "./keystoneConfig";


export function createPlayer() {
    const playerID = getPlayerID()
    if (playerID === undefined) {
        const playerWallet = ethers.Wallet.createRandom();
        const newPlayerID = testPlayerId; // TODO do random
        const base64PublicKey = hexToBase64(playerWallet.publicKey);

        CreatePlayer({PublicKey: base64PublicKey, PlayerId: newPlayerID}); // TODO call it base64PublicKey?

        setPlayerID(newPlayerID.toString())
        setBase64PublicKey(hexToBase64(playerWallet.publicKey))
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

    const playerID = parseInt(playerIDStr, 10)
    const player = PlayerTable.filter(worldState.tableState)
        .WithCondition(p => p.PlayerId === playerID)
        .Execute();
    if (player.length === 0) {
        return undefined
    }

    return player[0]
}

export function getPublicKeyBase64(): string {
    return localStorage.getItem(base64PublicKeyTag)!;
}

function setBase64PublicKey(base64PublicKey: string) {
    localStorage.setItem(base64PublicKeyTag, base64PublicKey)
}

export function getPrivateKey(): string {
    return localStorage.getItem(privateKeyTag)!;
}

function setPrivateKey(privateKey: string) {
    localStorage.setItem(privateKeyTag, privateKey)
}

export function getPlayerID(): number | undefined {
    const game = getGame()
    if (game === undefined) {
        return undefined
    }
    const playerTag = playerIdTag(game.GameId);

    const playerID = localStorage.getItem(playerTag);
    if (playerID) {
        return parseInt(playerID, 10);
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
export function hexToBase64(hex: string) {
    const bits = sjcl.codec.hex.toBits(hex);
    return sjcl.codec.base64.fromBits(bits);
}

function getGame(): GameSchema | undefined {
    return GameTable.get(worldState.tableState, gameEntity)
}
