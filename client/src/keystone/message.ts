import { CMD, S2CErrorMessage, S2CTestevent } from '../clientpb/proto/schemas/base';
import { match } from 'ts-pattern';
import { Simulate } from 'react-dom/test-utils';

import { blobToArrayBuffer } from './utils';

type IS2C_MessageType = S2CErrorMessage | S2CTestevent;

interface ProtoUnmarshalResponse {
  flag: number;
  command: number;
  param: number;
  bodyLength: number;
  data: IS2C_MessageType;
}

const MESSAGE_HEAD_LENGTH = 13;
export async function decode(blob: Blob): Promise<ProtoUnmarshalResponse | undefined> {
  try {
    const buffer = await blobToArrayBuffer(blob);
    const header = buffer.slice(0, MESSAGE_HEAD_LENGTH);

    const view = new DataView(header);

    const flag = view.getUint32(0, true);
    const command = view.getUint32(1, true);
    const param = view.getUint32(5, true);
    const bodyLength = view.getUint32(9, true);

    const dataSlice = buffer.slice(MESSAGE_HEAD_LENGTH);
    const data = decodeMessageData(command, new Uint8Array(dataSlice))!;

    return {
      flag,
      command,
      param,
      bodyLength,
      data,
    };
  } catch (e) {
    return undefined;
  }
}

function decodeMessageData(command: CMD, data: Uint8Array) {
  return match(Number(command))
    .returnType<IS2C_MessageType | null>()
    .with(CMD.S2C_Error, () => S2CErrorMessage.decode(data))
    .with(CMD.S2C_TestEvent, () => S2CTestevent.decode(data))
    .otherwise(() => {
      return null;
    });
}
