<script lang="ts">
  import Leaderboard from '../../lib/Leaderboard.svelte';
  import { leaderboard } from '../../service/host/host';
</script>

<div class="flex justify-center bg-purple-500 min-h-screen">
  <div class="mt-32">
    <h2 class="text-center text-3xl font-bold">Game Ended</h2>
    <div class="flex flex-wrap gap-4 mt-10">
      <Leaderboard leaderboard={$leaderboard} />
    </div>
  </div>
</div>