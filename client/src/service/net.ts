export class NetService {
  private websocket!: WebSocket;
  private textDecoder: TextDecoder = new TextDecoder();
  private textEncoder: TextEncoder = new TextEncoder();

  connect() {
    this.websocket = new WebSocket('ws://localhost:3000/ws');
    this.websocket.onopen = () => {
      console.log('opened connection');
    };

    this.websocket.onmessage = async (event: MessageEvent) => {
      const arrayBuffer: Iterable<number> = await event.data.arrayBuffer();
      const bytes: Uint8Array = new Uint8Array(arrayBuffer);
      const packetId: number = bytes[0];

      const packet = JSON.parse(this.textDecoder.decode(bytes.subarray(1)));

      console.log(packetId);
      console.log(packet);
    }
  }

  sendPacket(packet: any) {
    const packetId: number = 1337;
    const packetData: string = JSON.stringify(packet);

    const packetIdArray: Uint8Array = new Uint8Array([packetId]);
    const packetDataArray: Uint8Array = this.textEncoder.encode(packetData);

    const mergedArray: Uint8Array = new Uint8Array(
      packetIdArray.length + packetDataArray.length,
    );
    mergedArray.set(packetIdArray);
    mergedArray.set(packetDataArray, packetIdArray.length);

    this.websocket.send(mergedArray);
  }
}