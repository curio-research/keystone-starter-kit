import {ECDSAPublicKeyAuthHeader} from "../requests";
import {HeaderEntry, playerWallet} from "./middleware";
import sjcl from "sjcl";

interface ECDSAPublicKeyAuth {
    Base64Signature: string
    Base64Hash: string
    Base64PublicKey: string
}

export function WithECDSAAuth<T>(request: T): HeaderEntry {
    const jsonReq = JSON.stringify(request);

    const hashBits = sjcl.hash.sha256.hash(jsonReq);
    const hashHex = sjcl.codec.hex.fromBits(hashBits);
    const hashBase64 = sjcl.codec.base64.fromBits(hashBits);

    const signatureHex = playerWallet.signMessageSync(hashHex);
    const signatureBits = sjcl.codec.hex.toBits(signatureHex);
    const signatureBase64 = sjcl.codec.base64.fromBits(signatureBits);

    const publicKey = playerWallet.publicKey;
    const publicKeyBits = sjcl.codec.hex.toBits(publicKey);
    const publicKeyBase64 = sjcl.codec.base64.fromBits(publicKeyBits);

    const publicKeyAuth: ECDSAPublicKeyAuth = {
        Base64Hash: hashBase64,
        Base64Signature: signatureBase64,
        Base64PublicKey: publicKeyBase64,
    }
    return [ECDSAPublicKeyAuthHeader, publicKeyAuth];
}
