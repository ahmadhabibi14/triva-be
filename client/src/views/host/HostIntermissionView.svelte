<script lang="ts">
  import { leaderboard } from '../../service/host/host';
</script>

<div class="p-8 bg-purple-500 min-h-screen w-full">
  <div>
    {#each Object.entries($leaderboard) as [name, points]}
      <p>{name} - {points}</p>
    {/each}
  </div>
</div>