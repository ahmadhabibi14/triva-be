<script lang="ts">
  import QuizCard from './lib/QuizCard.svelte';
  import type { Quiz } from './types/quiz';
  import type { HTTPResponse } from './types/http';

  let quizzes: Quiz[] = [];

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

  let code: string = '', msg: string = '';

  function ConnectSocket() {
    let websocket: WebSocket = new WebSocket('ws://localhost:3000/api/ws');
    websocket.onopen = () => {
      console.log('opened connection');
      websocket.send(`join:${code}`)
    }

    websocket.onmessage = (event: MessageEvent) => {
      msg = event.data;
      console.log(event.data);
    }
  }

  function HostQuiz(quiz: Quiz) {
    let websocket: WebSocket = new WebSocket('ws://localhost:3000/api/ws');
    websocket.onopen = () => {
      console.log('opened connection');
      websocket.send(`host:${quiz.id}`)
    }
  }
</script>

<main class="container flex justify-center flex-col gap-6">
  <div class="mt-10 flex flex-row gap-6">
    <button class="bg-emerald-600 py-2 px-6 rounded-full text-white" on:click={GetQuizzes}>Get Quizzes</button>
    <button class="bg-sky-500 py-2 px-6 font-semibold text-white rounded-full">Cool</button>
  </div>
  {#if quizzes && quizzes.length}
    {#each quizzes as q, _ (q.id)}
      <QuizCard quiz={q} />
    {/each}
  {/if}
</main>