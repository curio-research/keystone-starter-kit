import {HeaderEntry} from "./middleware";
import sjcl from "sjcl";
import {ethers} from "ethers";

import {getPrivateKey} from "../utils";

interface EthereumWalletAuth {
    Base64Signature: string
    Base64Hash: string
    Base64PublicKey: string
}

export const ECDSAPublicKeyAuthHeader = "ecdsaPublicKeyAuth"

export function WithEthereumWalletAuth<T>(request: T): HeaderEntry<EthereumWalletAuth> {
    // Serialize the request to a JSON string
    const jsonReq = JSON.stringify(request);

    // Compute a SHA256 hash of the JSON request
    const hashBits = sjcl.hash.sha256.hash(jsonReq);
    const hashHex = "0x" + sjcl.codec.hex.fromBits(hashBits);
    const hashBase64 = sjcl.codec.base64.fromBits(hashBits);

    // Sign the hash with the wallet's private key
    const privateKey = getPrivateKey();
    const playerWallet = new ethers.Wallet(privateKey)

    const signature = playerWallet.signingKey.sign(hashHex).serialized;
    const signatureBits = sjcl.codec.hex.toBits(signature);
    const signatureBase64 = sjcl.codec.base64.fromBits(signatureBits);

    // Extract and encode the public key to base64
    const publicKeyBits = sjcl.codec.hex.toBits(playerWallet.signingKey.publicKey);
    const publicKeyBase64 = sjcl.codec.base64.fromBits(publicKeyBits);

    const publicKeyAuth: EthereumWalletAuth = {
        Base64Hash: hashBase64,
        Base64Signature: signatureBase64,
        Base64PublicKey: publicKeyBase64,
    };

    return [ECDSAPublicKeyAuthHeader, publicKeyAuth];
}

