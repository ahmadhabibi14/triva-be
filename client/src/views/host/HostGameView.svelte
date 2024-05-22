<script lang="ts">
  import { HostGame, state } from '../../service/host/host';
  import { GameState } from '../../service/net';
  import type { Quiz } from '../../types/quiz';
  import HostIntermissionView from './HostIntermissionView.svelte';
  import HostLobbyView from './HostLobbyView.svelte';
  import HostPlayView from './HostPlayView.svelte';
  import HostQuizListView from './HostQuizListView.svelte';

  let game: HostGame = new HostGame();
  let active: boolean = false;

  const onHost = (event: {detail: Quiz}) => {
    game.hostQuiz(event.detail.id);
    active = true;
  }

  let views: Record<GameState, any> = {
    [GameState.Lobby]: HostLobbyView,
    [GameState.Play]: HostPlayView,
    [GameState.Intermission]: HostIntermissionView,
    [GameState.Reveal]: HostPlayView,
    [GameState.End]: undefined
  }
</script>

{#if active}
  <svelte:component this={views[$state]} {game} />
{:else}
  <HostQuizListView on:host={onHost} />
{/if}
