<script lang="ts">
  import HostLoginView from './HostLoginView.svelte';
  import HostRegisterView from './HostRegisterView.svelte';
  import { ModeHostState, ModeHostLogin, ModeHostRegister } from '../../states/mode';
</script>

<div class="flex flex-col w-full justify-center h-fit mt-10">
  {#if $ModeHostState === ModeHostLogin}
    <HostLoginView />
  {/if}
  {#if $ModeHostState === ModeHostRegister}
    <HostRegisterView />
  {/if}
</div>
