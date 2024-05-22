<script lang="ts">
  import { onMount } from 'svelte';
  import QuizChoiceCard from '../../lib/play/QuizChoiceCard.svelte';
  import type { PlayerGame } from '../../service/player/player';
  import { COLORS } from '../../types/quiz';

  export let game: PlayerGame;
  let answered: boolean = false; 

  function onClick(i: number) {
    game.answer(i);
    answered = true;
  }

  onMount(() => {
    answered = false;
  });
</script>

<div class="flex flex-wrap w-full min-h-screen">
  {#if !answered}
    {#each COLORS as color, i}
      <QuizChoiceCard {color} >
        <button
          class="h-full w-full"
          on:click={()=> onClick(i)}
        >
        X
        </button>
      </QuizChoiceCard>
    {/each}
  {:else}
    <div class="p-8">
      <p class="text-2xl text-white">Lightning fast ?</p>
    </div>
  {/if}
</div>