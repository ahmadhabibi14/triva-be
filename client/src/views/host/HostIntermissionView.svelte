<script lang="ts">
  import { leaderboard } from '../../service/host/host';
</script>

<div class="p-8 bg-purple-500 min-h-screen w-full flex flex-col gap-6">
  <div>
    <button class="bg-zinc-800 py-2 px-4 rounded-lg">Skip</button>
  </div>
  <div>
    {#each $leaderboard as entry}
      <p>{entry.name} - {entry.points}</p>
    {/each}
  </div>
</div>