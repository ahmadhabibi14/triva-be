<script lang="ts">
  import GuestView from './views/guest/GuestView.svelte';
  import PlayerView from './views/player/PlayerView.svelte';
  import HostView from './views/host/HostView.svelte';
  import Router from 'svelte-spa-router';

  let routes = {
    '/': GuestView,
    '/player': PlayerView,
    '/host': HostView
  };
</script>

<Router {routes} />