
import { match } from 'ts-pattern'
import {CMD, S2CErrorMessage} from "../../clientpb/base";

type IS2C_MessageType =
    | CMD.S2C_Error

const MESSAGE_HEAD_LENGTH = 13

export function decode (buffer: ArrayBuffer) {
    const header = buffer.slice(0, MESSAGE_HEAD_LENGTH)

    const view = new DataView(header)

    const flag = view.getUint32(0, true)
    const command = view.getUint32(1, true)
    const param = view.getUint32(5, true)
    const bodyLength = view.getUint32(9, true)

    const dataSlice = buffer.slice(MESSAGE_HEAD_LENGTH)
    const data = decodeMessageData(command, new Uint8Array(dataSlice))

    return {
        flag,
        command,
        param,
        bodyLength,
        data,
    }
}
function decodeMessageData(command: CMD, data: Uint8Array) {
    return match(Number(command))
        .returnType<IS2C_MessageType | null>()
        .with(
            CMD.S2C_Error,
            () => S2CErrorMessage.decode(data),
        ).otherwise(() => {
            return null;
        })
}

