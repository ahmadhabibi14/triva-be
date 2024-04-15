<script lang="ts">
  import { PlayerGame } from '../../service/player/player';
  import PlayerJoinView from './PlayerJoinView.svelte';

  let game: PlayerGame = new PlayerGame();
  let active: boolean = false;

  const onJoin = () => active = true;
</script>

{#if active}
  active
{:else}
  <PlayerJoinView on:join={onJoin} {game} />
{/if}