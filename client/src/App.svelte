<script lang="ts">
  import type { Player, Quiz, QuizQuestion } from './types/quiz';
  import type { HTTPResponse } from './types/http';
  import { NetService, PacketTypes, type ChangeGameStatePacket, GameState, type PlayerJoinPacket, type TickPacket } from './service/net';
  import PlayerView from './views/player/PlayerView.svelte';
  import HostView from './views/host/HostView.svelte';
  import Router from 'svelte-spa-router';

  let quizzes: Quiz[] = [];

  let currentQuestion: QuizQuestion | null = null;
  let state: number = -1;
  let host: boolean = false;
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

  async function GetQuizzes(): Promise<void> {
    let response: Response = await fetch('http://localhost:3000/api/quizzes');
    if (!response.ok) {
      alert('failed');
      return;
    }

    let responseData: HTTPResponse = await response.json() as HTTPResponse;

    let json: Quiz[] = responseData.data as Quiz[];
    quizzes = json;
  }

  let gameCode: string = '', name: string = '';

  function joinGame() {
    netService.sendPacket({
      id: PacketTypes.Connect,
      code: gameCode.trim(),
      name: name.trim()
    })
  }

  function startGame() {
    netService.sendPacket({
      id: PacketTypes.StartGame,
    })
  }

  function hostQuiz(quiz: Quiz) {
    host = true;
    netService.sendPacket({
      id: PacketTypes.HostGame,
      quizId: quiz.id
    });
  }

  let routes = {
    '/': PlayerView,
    '/host': HostView
  };
</script>

<Router {routes} />