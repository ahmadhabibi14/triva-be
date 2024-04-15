<script lang="ts">
  import Button from '../../lib/Button.svelte';
import type { HostGame } from '../../service/host/host';


  export let game: HostGame;

  const start = () => game.start();
</script>

<div class="p-8">
  <div class="flex justify-end">
    <Button on:click={start}>Start Game</Button>
  </div>
</div>