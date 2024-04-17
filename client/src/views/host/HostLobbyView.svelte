<script lang="ts">
  import Button from '../../lib/Button.svelte';
  import PlayerNameCard from '../../lib/lobby/PlayerNameCard.svelte';
  import { type HostGame, players } from '../../service/host/host';


  export let game: HostGame;

  const start = () => game.start();
</script>

<div class="p-8">
  <div class="flex justify-end">
    <Button on:click={start}>Start Game</Button>
  </div>
  <div class="flex flex-wrap gap-2">
    {#each $players as player (player.id)}
      <PlayerNameCard {player} />
    {/each}
  </div>
</div>