import { writable, type Writable } from 'svelte/store';
import { NetService, PacketTypes, type HostGamePacket, type Packet } from '../net';
import type { Player } from '../../types/quiz';

export const players: Writable<Player[]> = writable([]);

export class HostGame {
  private net: NetService;

  constructor() {
    this.net = new NetService();
    this.net.connect();
    this.net.onPacket(p => this.onPacket(p));
  }

  hostQuiz(quizId: string) {
    let packet: HostGamePacket = {
      id: PacketTypes.HostGame,
      quizId: quizId
    }

    this.net.sendPacket(packet);
  }

  onPacket(packet: Packet) {
    switch (packet.id) {
      case PacketTypes.PlayerJoin: {
        break;
      }
    }
  }
}