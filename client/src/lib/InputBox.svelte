<script lang="ts">
  import { onMount } from 'svelte';

  export let label: string = '';
  export let value: string = '';
  export let type: string = 'text';
  export let id: string = '';
  export let placeholder: string = '';

  let inputElm: HTMLInputElement;

  onMount(() => inputElm.type = type);
</script>

<div class="flex flex-col gap-2">
  <label for={id} class="text-sm ml-2">{label}</label>
  <input
    bind:this={inputElm}
    bind:value

    {id}
    {placeholder}
    autocomplete="off"
    class="bg-zinc-800 py-2 px-3 rounded-lg outline-none focus:outline-violet-500"
  />
</div>