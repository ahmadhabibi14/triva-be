<script lang="ts">
  import { onMount } from 'svelte';
  import { Icon } from 'svelte-icons-pack';
  import { AiOutlineEye, AiOutlineEyeInvisible } from 'svelte-icons-pack/ai';

  export let label: string = '';
  export let value: string = '';
  export let type: string = 'text';
  export let id: string = '';
  export let placeholder: string = '';

  let inputElm: HTMLInputElement;
  let isShowPassword: boolean = false;
  let showPasswordState: string = 'show';

  onMount(() => inputElm.type = type);

  function toggleShowPassword() {
    isShowPassword = !isShowPassword;
    if (isShowPassword) inputElm.type = 'text', showPasswordState = 'hide';
    else inputElm.type = 'password', showPasswordState = 'show';
  }
</script>

<div class="flex flex-col gap-2 relative">
  <label for={id} class="text-sm ml-2">{label}</label>
  <input
    bind:this={inputElm}
    bind:value

    {id}
    {placeholder}
    autocomplete="off"
    class="bg-zinc-800 py-2 px-3 rounded-lg outline-none focus:outline-violet-500"
  />
  {#if type === 'password'}
    <button
      title="{showPasswordState} password"
      on:click={toggleShowPassword}
      class="absolute h-fit w-fit p-0 top-[38px] bottom-auto right-2 cursor-pointer group focus:outline-none"
    >
      {#if !isShowPassword}
        <Icon
          size="20"
          src={AiOutlineEye}
          className="fill-zinc-100 group-hover:fill-violet-400"
        />
      {/if}
      {#if isShowPassword}
        <Icon
          size="20"
          src={AiOutlineEyeInvisible}
          className="fill-zinc-100 group-hover:fill-violet-400"
        />
      {/if}
    </button>
  {/if}
</div>