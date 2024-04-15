import { NetService, PacketTypes, type ConnectPacket, type Packet } from '../net';

export class PlayerGame {
  private net: NetService;

  constructor() {
    this.net = new NetService();
    this.net.connect();
    this.net.onPacket(p => this.onPacket(p));
  }

  join(code: string, name: string) {
    let packet: ConnectPacket = {
      id: PacketTypes.Connect,
      code: code,
      name: name,
    }

    this.net.sendPacket(packet);
  }

  onPacket(packet: Packet) {

  }
}