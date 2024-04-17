<script lang="ts">
  export let color: string;
</script>

<div class="w-1/2 text-white text-6xl font-bold flex justify-center items-center {color}">
  <slot />
</div>