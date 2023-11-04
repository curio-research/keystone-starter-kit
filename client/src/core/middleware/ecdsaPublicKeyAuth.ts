import {ECDSAPublicKeyAuthHeader} from "../requests";
import {HeaderEntry, playerWallet} from "./middleware";
import sjcl from "sjcl";

interface ECDSAPublicKeyAuth {
    Base64Signature: string
    Base64Hash: string
    Base64PublicKey: string
}

export function WithECDSAAuth<T>(request: T): HeaderEntry<ECDSAPublicKeyAuth> {
    // Serialize the request to a JSON string
    const jsonReq = JSON.stringify(request);

    // Compute a SHA256 hash of the JSON request
    const hashBits = sjcl.hash.sha256.hash(jsonReq);
    const hashHex = sjcl.codec.hex.fromBits(hashBits);
    const hashBase64 = sjcl.codec.base64.fromBits(hashBits);

    // Sign the hash with the wallet's private key
    const signature = playerWallet.signingKey.sign("0x" + hashHex).serialized;
    const signatureBits = sjcl.codec.hex.toBits(signature);
    const signatureBase64 = sjcl.codec.base64.fromBits(signatureBits);

    // Extract and encode the public key to base64
    const publicKeyBits = sjcl.codec.hex.toBits(playerWallet.signingKey.publicKey);
    const publicKeyBase64 = sjcl.codec.base64.fromBits(publicKeyBits);

    const publicKeyAuth: ECDSAPublicKeyAuth = {
        Base64Hash: hashBase64,
        Base64Signature: signatureBase64,
        Base64PublicKey: publicKeyBase64,
    };

    return [ECDSAPublicKeyAuthHeader, publicKeyAuth];
}

