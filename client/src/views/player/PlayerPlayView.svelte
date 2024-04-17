<script lang="ts">
  import QuizChoiceCard from '../../lib/play/QuizChoiceCard.svelte';
import type { PlayerGame } from '../../service/player/player';
  import { COLORS } from '../../types/quiz';

  export let game: PlayerGame;

  function onClick(i: number) {
    game.answer(i);
  }
</script>

<div class="flex flex-wrap w-full min-h-screen">
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
</div>