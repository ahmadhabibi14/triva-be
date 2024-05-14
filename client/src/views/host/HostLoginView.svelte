<script lang="ts">
  import { ModeHostView } from '../../states/mode';
  import InputBox from '../../lib/InputBox.svelte';
  import { link } from 'svelte-spa-router';
  import { Icon } from 'svelte-icons-pack';
  import { TrOutlineHome } from 'svelte-icons-pack/tr';

  let email: string = '';
  let password: string = '';

  let isSubmitted: boolean = false;
</script>

<div class="w-[400px] bg-zinc-900 p-5 rounded-xl mx-auto flex flex-col gap-7">
  <header class="flex flex-row justify-between items-center">
    <h1 class="text-3xl font-bold text-end">Login <span class="text-emerald-600">&gt;</span></h1>
    <a href="/" use:link class="flex justify-center items-center p-2 bg-zinc-800 rounded-lg group hover:bg-violet-500/10">
      <Icon src={TrOutlineHome} size="20" className="group-hover:stroke-violet-500" />
    </a>
  </header>
  <form class="flex flex-col gap-4">
    <InputBox
      id="email"
      type="email"
      bind:value={email}
      placeholder="john@example.com"
      label="Email"
    />
    <InputBox
      id="password"
      type="password"
      bind:value={password}
      placeholder="password123"
      label="Password"
    />
    <div class="flex flex-row gap-1 justify-center text-sm">
      <span>Have no account?</span>
      <button
        class="text-sky-500 hover:underline"
        on:click|preventDefault={() => ModeHostView.set('register')}
      >Register</button>
    </div>
    <button
      disabled={isSubmitted}
      class="bg-violet-500 disabled:cursor-not-allowed py-2.5 rounded-lg font-semibold
        text-white hover:bg-violet-400 disabled:bg-zinc-800"
      on:click|preventDefault={() => isSubmitted = !isSubmitted}
    >
      {#if !isSubmitted}
        Login
      {/if}
      {#if isSubmitted}
        Logging in...
      {/if}
    </button>
  </form>
</div>