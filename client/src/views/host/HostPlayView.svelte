<script lang="ts">
  import QuizChoiceCard from '../../lib/play/QuizChoiceCard.svelte';
  import { type HostGame, currentQuestion, tick, state} from '../../service/host/host';
  import { GameState } from '../../service/net';
  import { COLORS, type QuizChoice } from '../../types/quiz';

  function getCardColor(choice: QuizChoice, state: GameState, defaultColor: string) {
    if (state !== GameState.Reveal) return defaultColor;

    return choice.correct ? 'bg-green-500' : 'bg-red-500';
  }

  export let game: HostGame;

  console.log(game);
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
        <QuizChoiceCard color={getCardColor($currentQuestion.choices[i], $state, color)}>
          {$currentQuestion.choices[i].name}
        </QuizChoiceCard>
      {/each}
    </div>
  </div>
{/if}