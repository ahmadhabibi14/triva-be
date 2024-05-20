<script lang="ts">
  import { GameState } from '@/service/net';
  import { PlayerGame } from '@/service/player/player';
  import PlayerJoinView from '@/views/player/PlayerJoinView.svelte';
  import PlayerLobbyView from '@/views/player/PlayerLobbyView.svelte';
  import PlayerPlayView from '@/views/player/PlayerPlayView.svelte';
  import { state } from '@/service/player/player';
  import PlayerRevealView from '@/views/player/PlayerRevealView.svelte';

  let game: PlayerGame = new PlayerGame();
  let active: boolean = false;

  const onJoin = () => active = true;

  let views: Record<GameState, any> = {
    [GameState.Lobby]: PlayerLobbyView,
    [GameState.Play]: PlayerPlayView,
    [GameState.Reveal]: PlayerRevealView,
    [GameState.Intermission]: undefined,
    [GameState.End]: undefined
  }
</script>

{#if active}
  <svelte:component this={views[$state]} {game} />
{:else}
  <PlayerJoinView on:join={onJoin} {game} />
{/if}