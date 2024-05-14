<script lang="ts">
  import HostLoginView from './HostLoginView.svelte';
  import HostRegisterView from './HostRegisterView.svelte';
  import { ModeHostView } from '../../states/mode';

  const MODE_LOGIN: string = 'login';
  const MODE_REGISTER: string = 'register';
</script>

<div class="flex flex-col w-full justify-center h-fit mt-10">
  {#if $ModeHostView === MODE_LOGIN}
    <HostLoginView />
  {/if}
  {#if $ModeHostView === MODE_REGISTER}
    <HostRegisterView />
  {/if}
</div>
