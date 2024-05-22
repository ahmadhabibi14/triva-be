<script lang="ts">
  import HostLoginView from './HostLoginView.svelte';
  import HostRegisterView from './HostRegisterView.svelte';
  import { ModeHostState, ModeHostPassed, ModeHostLogin, ModeHostRegister } from '../../states/mode';
  import { onMount } from 'svelte';

  const sessionId: string = localStorage.getItem('session_id') || '';

  onMount(() => {
    if (sessionId === '') ModeHostState.set(ModeHostLogin);
    else ModeHostState.set(ModeHostPassed);
  });
</script>

<div class="flex flex-col w-full justify-center h-fit mt-10">
  {#if $ModeHostState === ModeHostPassed}
    <div>
      <h1>Host</h1>
    </div>
  {/if}
  {#if $ModeHostState === ModeHostLogin}
    <HostLoginView />
  {/if}
  {#if $ModeHostState === ModeHostRegister}
    <HostRegisterView />
  {/if}
</div>
