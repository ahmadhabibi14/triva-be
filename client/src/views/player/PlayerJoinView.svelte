<script lang="ts">
  import { createEventDispatcher, type EventDispatcher } from 'svelte';
  import type { PlayerGame } from '../../service/player/player';
  import { errorState } from '../../states/error';
  import toast, { Toaster } from 'svelte-french-toast';

  const dispatch: EventDispatcher<any> = createEventDispatcher();

  let code: string = '';
  let name: string = '';
  export let game: PlayerGame;

  function join() {
    game.join(code, name);
  }

  $: {
    if ($errorState !== '') {
      toast.error($errorState);
      errorState.set('');
    }
  }
</script>

<Toaster />

<div class="min-h-screen w-full flex items-center justify-center">
  <div class="flex flex-col gap-5">
    <h2 class="font-bold text-center text-5xl">Triva</h2>
    <div class="flex flex-col gap-3">
      <input
        type="text"
        placeholder="Enter game code"
        autocomplete="off"
        bind:value={code}
        class="py-2 px-4 rounded-lg bg-zinc-900 border border-zinc-800 caret-indigo-500 focus:border-indigo-500 focus:outline focus:outline-indigo-500"
      />
      <input
        type="text"
        placeholder="Name"
        autocomplete="off"
        bind:value={name}
        class="py-2 px-4 rounded-lg bg-zinc-900 border border-zinc-800 caret-indigo-500 focus:border-indigo-500 focus:outline focus:outline-indigo-500"
      />
      <button
        on:click={join}
        class="bg-emerald-600 hover:bg-emerald-500 py-2 rounded-lg border border-emerald-700 hover:border-emerald-600"
      >
        Join game
      </button>
    </div>
  </div>
</div>