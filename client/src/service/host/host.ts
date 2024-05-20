import { writable, type Writable } from 'svelte/store';
import { GameState, NetService, PacketTypes, type ChangeGameStatePacket, type HostGamePacket, type LeaderboardEntry, type LeaderboardPacket, type Packet, type PlayerJoinPacket, type QuestionShowPacket, type TickPacket } from '../net';
import type { Player, QuizQuestion } from '@/types/quiz';

export const state: Writable<GameState> = writable(GameState.Lobby);
export const players: Writable<Player[]> = writable([]);
export const tick: Writable<number> = writable(0);
export const leaderboard: Writable<LeaderboardEntry[]> = writable([]);
export const currentQuestion: Writable<QuizQuestion | null> = writable(null);

export class HostGame {
  private net: NetService;

  constructor() {
    this.net = new NetService();
    this.net.connectHost();
    this.net.onPacket(p => this.onPacket(p));
  }

  hostQuiz(quizId: string) {
    let packet: HostGamePacket = {
      id: PacketTypes.HostGame,
      quizId: quizId
    }

    this.net.sendPacket(packet);
  }

  start() {
    this.net.sendPacket({ id: PacketTypes.StartGame })
  }

  onPacket(packet: Packet) {
    switch (packet.id) {
      case PacketTypes.ChangeGameState: {
        let data = packet as ChangeGameStatePacket;
        state.set(data.state);
        break;
      }
      case PacketTypes.PlayerJoin: {
        let data = packet as PlayerJoinPacket;
        players.update(p => [...p, data.player]);
        break;
      }
      case PacketTypes.Tick: {
        let data = packet as TickPacket;
        tick.set(data.tick);
        break;
      }
      case PacketTypes.QuestionShow: {
        let data = packet as QuestionShowPacket;
        currentQuestion.set(data.question);
      }
      case PacketTypes.Leaderboard: {
        let data = packet as LeaderboardPacket;
        leaderboard.set(data.points);
        break;
      }
    }
  }
}