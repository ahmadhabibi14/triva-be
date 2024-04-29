<script lang="ts">
  import GuestView from './views/guest/GuestView.svelte';
  import PlayerView from './views/player/PlayerView.svelte';
  import HostView from './views/host/HostView.svelte';
  import Router from 'svelte-spa-router';

  console.log(`WebSocket host:`, import.meta.env.VITE_WEBSOCKET_HOST);

  let routes = {
    '/': GuestView,
    '/player': PlayerView,
    '/host': HostView
  };
</script>

<Router {routes} />