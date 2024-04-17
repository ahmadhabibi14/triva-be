<script lang="ts">
  import QuizChoiceCard from '../../lib/play/QuizChoiceCard.svelte';
import { type HostGame, currentQuestion, tick } from '../../service/host/host';

  const COLORS: string[] = ['bg-pink-500', 'bg-blue-500', 'bg-yellow-500', 'bg-green-500'];

  export let game: HostGame;
</script>

{#if $currentQuestion != null}
  <div class="min-h-screen h-screen flex flex-col">
    <div class="text-3xl border-b p-4 font-bold">
      {$currentQuestion.name}
    </div>
    <div class="flex-l flex flex-col justify-center pl-4">
      <div class="flex justify-between items-center">
        <div class="bg-purple-500 text-1xl h-16 w-16 rounded-full flex items-center justify-center ml-8">
          {$tick}
        </div>
        <img src="" alt="center" class="max-w-[500px]" />
        <div class="w-24"></div>
      </div>
    </div>
    <div class="flex flex-wrap w-full h-96">
      {#each COLORS as color, i}
        <QuizChoiceCard {color} choice={$currentQuestion.choices[i]} />
      {/each}
    </div>
  </div>
{/if}