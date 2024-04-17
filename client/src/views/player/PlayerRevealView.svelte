<script lang="ts">
  import { points } from '../../service/player/player';

  $: correct = $points > 0;
</script>

<div class="min-h-full w-full {correct
    ? 'bg-green-500'
    : 'bg-red-500'} flex justify-center items-center"
>
  {#if correct}
    <h2>Correct !</h2>
    <p>+ {$points} point</p>
  {:else}
    <h2>Incorrect</h2>
  {/if}
</div>