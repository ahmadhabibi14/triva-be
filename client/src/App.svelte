<script lang="ts">
  import QuizCard from './lib/QuizCard.svelte';

  let quizzes: {id: string, name: string}[] = [];

  async function getQuizzes() {
    let response = await fetch('http://localhost:3000/api/quizzes');
    if (!response.ok) {
      alert('failed');
      return;
    }

    let json = await response.json();
    quizzes = json;
  }
</script>

<main class="container flex justify-center flex-col gap-6">
  <div class="mt-10 flex flex-row gap-6">
    <button class="bg-emerald-600 py-2 px-6 rounded-full text-white" on:click={getQuizzes}>Get Quizzes</button>
    <button class="bg-sky-500 py-2 px-6 font-semibold text-white rounded-full">Cool</button>
  </div>
  {#if quizzes && quizzes.length}
    {#each quizzes as q, _ (q.id)}
      <QuizCard quiz={q} />
    {/each}
  {/if}
</main>