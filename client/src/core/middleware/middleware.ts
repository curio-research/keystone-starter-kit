import {ethers} from "ethers";

export const playerWallet = ethers.Wallet.createRandom();

export interface KeystoneTx<T> {
    Headers: Map<string, any>
    Req: T
}

export type HeaderEntry = [string, any]
export function NewKeystoneTx<T>(request: T, ...headerEntries: HeaderEntry[]): KeystoneTx<T> {
    const headersMap = new Map<string, any>(headerEntries);

    return {
        Headers: headersMap,
        Req: request
    }
}


