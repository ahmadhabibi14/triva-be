<script lang="ts">
  import QuizCard from '../../lib/QuizCard.svelte';
  import type { HTTPResponse } from '../../types/http';
  import type { Quiz } from '../../types/quiz';

  let quizzes: Quiz[] = [];

  async function getQuizzes(): Promise<Quiz[]> {
    let response: Response = await fetch('http://localhost:3000/api/quizzes');
    if (!response.ok) {
      alert('failed to get quizzes');
      return [];
    }

    let responseData: HTTPResponse = await response.json() as HTTPResponse;
    let json: Quiz[] = responseData.data as Quiz[];

    return json;
  }

  (async () => {
    quizzes = await getQuizzes();
  })();
</script>

<div class="p-8 max-w-7xl mx-auto">
  <h2 class="text-4xl font-bold">Your Quizzes</h2>
  <div class="flex flex-col gap-2 mt-4">
    {#each quizzes as quiz (quiz.id)}
      <QuizCard on:host {quiz} />
    {/each}
  </div>
</div>