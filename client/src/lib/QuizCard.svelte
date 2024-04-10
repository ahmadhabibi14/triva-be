<script lang="ts">
  import type { Quiz } from "../types/quiz";
  export let quiz: Quiz;
</script>

<div class="bg-zinc-900 rounded-lg py-3 px-6 flex flex-row justify-between gap-3">
  <p>{quiz.name}</p>
  <button
    on:click
    class="bg-emerald-600 hover:bg-emerald-500 text-white rounded-md py-1 px-3 text-sm"
  >Host</button>
</div>