<script lang="ts">
  import type { Player, QuizQuestion } from './types/quiz';
  import { NetService, PacketTypes, type ChangeGameStatePacket, type PlayerJoinPacket, type TickPacket } from './service/net';
  import GuestView from './views/guest/GuestView.svelte';
  import PlayerView from './views/player/PlayerView.svelte';
  import HostView from './views/host/HostView.svelte';
  import Router from 'svelte-spa-router';

  let currentQuestion: QuizQuestion | null = null;
  let state: number = -1;
  let tick: number = 0;

  let players: Player[] = [];

  let netService = new NetService();

  setTimeout(() => {
    netService.connect();
  }, 500);
  
  netService.onPacket((packet: any) => {
    console.log(packet);
    switch (packet.id) {
      case PacketTypes.QuestionShow: {
        currentQuestion = packet.question;
        break;
      }
      case PacketTypes.ChangeGameState: {
        let data = packet as ChangeGameStatePacket;
        state = data.state;
        break;
      }
      case PacketTypes.PlayerJoin: {
        let data = packet as PlayerJoinPacket; 
        players = [...players, data.player];
        break;
      }
      case PacketTypes.Tick: {
        let data = packet as TickPacket;
        tick = data.tick;
        break;
      }
    }
  });

  let routes = {
    '/': GuestView,
    '/player': PlayerView,
    '/host': HostView
  };
</script>

<Router {routes} />