import {ethers, Signature} from "ethers";
import {Buffer} from "buffer";
import {ECDSASignature, ecsign} from "ethereumjs-util";
import {ECDSAPublicKeyAuthHeader} from "./requests";
import any = jasmine.any;

export const playerWallet = ethers.Wallet.createRandom();

export interface KeystoneTx<T> {
    Headers: Map<string, any>
    Req: T
}

interface ECDSAPublicKeyAuth {
    Base64Signature: string
    Base64Hash:      string
    Base64PublicKey: string
}

export function NewKeystoneTx<T>(request: T, headers: Map<string,any> | null): KeystoneTx<T> {
    const headersMap = new Map<string, any>();
    if (headers != null) {
        headers.forEach((value, key) => {
            headersMap.set(key, value)
        })
    }

    return {
        Headers: headersMap,
        Req: request
    }
}

function ECDSAAuthHeaders<T>(request: T) {
    const hash = Buffer.from(JSON.stringify(request), "utf8");
    const base64Hash = bytesToBase64(hash);

    const privateKey = Buffer.from(playerWallet.privateKey, "hex");
    const signature = ecsign(hash, privateKey);
    const base64Sig = ecdsaSignatureToBase64(signature)!;

    const publicKey = Buffer.from(playerWallet.publicKey, "hex");
    const base64PublicKey = bytesToBase64(publicKey);

    const publicKeyAuth: ECDSAPublicKeyAuth = {
        Base64Hash: base64Hash,
        Base64Signature: base64Sig,
        Base64PublicKey: base64PublicKey,
    }
    return publicKeyAuth;
}

function ecdsaSignatureToBase64(e: ECDSASignature): string | null {
    // Step 1: Convert the signature to the desired format (if necessary)
    // If you're already getting the signature in the right format, you can skip this step
    let ecdsaSignature: Signature
    if (e.v === 27 || e.v === 28) {
        ecdsaSignature = new Signature(null, e.r.toString(), e.s.toString(), e.v);
    } else {
        return null // TODO when to return null or undefined?
    }

    // Step 2: Get the 'r' and 's' components in hexadecimal format
    const rHex = ecdsaSignature.r.padStart(64, '0');  // Ensuring it's 32 bytes
    const sHex = ecdsaSignature.s.padStart(64, '0');  // Ensuring it's 32 bytes

    // Step 3: Concatenate the 'r' and 's' components as a single hexadecimal string
    const concatenatedHex = rHex + sHex;

    // Step 4: Convert the concatenated hexadecimal string to bytes
    const concatenatedBytes = Buffer.from(concatenatedHex, 'hex');

    // Step 5: Encode the bytes as Base64
    return bytesToBase64(concatenatedBytes);
}

function bytesToBase64(b: Buffer): string {
    return btoa(String.fromCharCode.apply(null, Array.from(b)));
}
