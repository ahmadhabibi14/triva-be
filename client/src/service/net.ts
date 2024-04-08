export class NetService {
  private websocket!: WebSocket;
  private textDecoder: TextDecoder = new TextDecoder();

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
    }
  }
}