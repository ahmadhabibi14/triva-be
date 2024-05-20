<script lang="ts">
  import type { Player } from '@/types/quiz';

  export let player: Player;
</script>

<button on:click>
  <h3
    class="text-3xl px-4 rounded-xl bg-zinc-900 hover:line-through"
  >
    {player.name}
  </h3>
</button>