<script lang="ts">
  import { link } from 'svelte-spa-router';
</script>

<div class="flex justify-center h-screen items-center">
  <div class="flex flex-col gap-6 items-center justify-center">
    <h3 class="text-6xl text-emerald-500 font-bold">Triva</h3>
    <div class="grid grid-cols-2 w-[400px] h-[70px] gap-4">
      <a
        href="/host"
        use:link
        class="bg-zinc-900 hover:bg-sky-400/10 border-[2px] border-zinc-800 hover:border-sky-400 hover:text-sky-600 select-none cursor-pointer text-xl rounded-full flex justify-center items-center"
        >
        <span>Host a game</span>
      </a>
      <a
        href="/player"
        use:link
        class="bg-zinc-900 hover:bg-sky-400/10 border-[2px] border-zinc-800 hover:border-sky-400 hover:text-sky-600 select-none cursor-pointer text-xl rounded-full flex justify-center items-center"
        >
        <span>Join a game</span>
      </a>
    </div>
  </div>
</div>