<script lang="ts">
  import { createEventDispatcher, type EventDispatcher } from 'svelte';
  import type { Quiz } from '../types/quiz';
  import Button from './Button.svelte';
  
  const dispatch: EventDispatcher<any> = createEventDispatcher();
  export let quiz: Quiz;
  
  const host = () => dispatch('host', quiz);
</script>

<div class="bg-zinc-900 rounded-lg py-3 px-6 flex flex-row justify-between gap-3">
  <p>{quiz.name}</p>
  <Button on:click={host}>Host</Button>
</div>