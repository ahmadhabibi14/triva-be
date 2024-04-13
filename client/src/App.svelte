<script lang="ts">
  import QuizCard from './lib/QuizCard.svelte';
  import type { Player, Quiz, QuizQuestion } from './types/quiz';
  import type { HTTPResponse } from './types/http';
  import { NetService, PacketTypes, type ChangeGameStatePacket, GameState, type PlayerJoinPacket } from './service/net';

  let quizzes: Quiz[] = [];

  let currentQuestion: QuizQuestion|null = null;
  let state: number = -1;
  let host: boolean = false;

  let players: Player[] = [];

  let netService = new NetService();
  netService.connect();
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

  function connect() {
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
</script>

<main class="flex justify-center flex-col gap-6 p-5 w-full">
  <div class="min-w-[400px] w-[700px] max-w-[700px] mx-auto flex flex-col gap-6">
  {#if state === -1}
    <h2 class="text-5xl text-sky-500 font-bold text-center">Bwizz - Quiz app</h2>
    <div class="flex flex-row justify-center gap-3">
      <input
        type="text"
        placeholder="Enter game code"
        bind:value={gameCode}
        class="py-2 px-4 rounded-lg bg-zinc-900 border border-zinc-800 caret-indigo-500 focus:border-indigo-500 focus:outline focus:outline-indigo-500"
      />
      <input
        type="text"
        placeholder="Name"
        bind:value={name}
        class="py-2 px-4 rounded-lg bg-zinc-900 border border-zinc-800 caret-indigo-500 focus:border-indigo-500 focus:outline focus:outline-indigo-500"
      />
      <button
        on:click={connect}
        class="py-2 px-6 rounded-lg bg-indigo-600 hover:bg-indigo-500 text-white"
      >Join</button>
    </div>
    <button
        on:click={GetQuizzes}
        class="bg-sky-600 hover:bg-sky-500 text-white py-2 px-6 rounded-lg"
      >Get Quizzes</button>
    {#if quizzes && quizzes.length}
      <div class="flex flex-col gap-3">
        {#each quizzes as q, _ (q.id)}
          <QuizCard quiz={q} on:click={() => hostQuiz(q)}/>
        {/each}
      </div>
    {/if}

    {#if currentQuestion && currentQuestion !== null}
      <div class="flex flex-col gap-3">
        <h4 class="text-2xl">{currentQuestion?.name || 'No question'}</h4>
        {#if currentQuestion.choices && currentQuestion.choices.length}
          <div class="flex flex-col gap-2">
            {#each currentQuestion.choices as c}
              <button class="bg-zinc-900 hover:bg-zinc-800 py-2 px-4 rounded-lg">
                {c.name}
              </button>
            {/each}
          </div>
        {/if}
      </div>
    {/if}
  {:else if state === GameState.Lobby}
    {#if host}
      <div class="flex flex-col">
        <button on:click={startGame}>Start game</button>
        <p>Lobby state</p>
        {#if players && players.length}
          <p>Players:</p>
          {#each players as p}
            <p>{p.name}</p>
          {/each}
        {/if}
      </div>
    {:else}
      <p>You successfully joined game</p>
    {/if}
  {/if}
  </div>
</main>