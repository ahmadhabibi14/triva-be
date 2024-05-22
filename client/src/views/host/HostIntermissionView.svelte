<script lang="ts">
  import Leaderboard from '../../lib/Leaderboard.svelte';
  import { HostGame, leaderboard } from '../../service/host/host';

  export let game: HostGame;

  function skip() {
    game.start();
  }
</script>

<div class="p-8 bg-purple-500 min-h-screen w-full flex flex-col gap-6">
  <div>
    <button class="bg-zinc-800 py-2 px-4 rounded-lg" on:click={skip}>Skip</button>
  </div>
  <div>
    <Leaderboard leaderboard={$leaderboard} />
  </div>
</div>