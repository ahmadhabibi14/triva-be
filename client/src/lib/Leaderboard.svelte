<script lang="ts">
  import { type LeaderboardEntry } from '../service/net';
  export let leaderboard: LeaderboardEntry[] = [];
</script>

<div>
  <h3>Leaderboard</h3>
  <div>
    {#each leaderboard as entry}
      <p>{entry.name} - {entry.points}</p>
    {/each}
  </div>
</div>