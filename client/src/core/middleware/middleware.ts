import {ethers} from "ethers";

export interface KeystoneTx<T> {
    Headers: { [p: string]: any }
    Data: T
}

export type HeaderEntry<T> = [string, T]

export function NewKeystoneTx<T>(request: T, ...headerEntries: HeaderEntry<any>[]): KeystoneTx<T> {
    const headersMapJSON = Object.fromEntries(headerEntries);
    return {
        Headers: headersMapJSON,
        Data: request
    }
}


