<script lang="ts">
  import { ModeHostState, ModeHostLogin} from '../../states/mode';
  import InputBox from '../../lib/InputBox.svelte';
  import { link } from 'svelte-spa-router';
  import { Icon } from 'svelte-icons-pack';
  import { TrOutlineHome } from 'svelte-icons-pack/tr';
  import axios from 'axios';
  import toast, { Toaster } from 'svelte-french-toast';
  import { type HTTPResponse } from '../../types/http';

  let email: string = '';
  let username: string = '';
  let full_name: string = '';
  let password: string = '';

  let isSubmitted: boolean = false;

  type registerIn = {
    email: string;
    username: string;
    full_name: string;
    password: string;
	}
  
  type registerOut = {
    message: string;
  }

  async function Register() {
    isSubmitted = true;

    const payload: registerIn = { email, username, full_name, password }
    const url: string = import.meta.env.VITE_WEB_HOST+'/auth/register';

    await axios.post(url, payload,
      { headers: { 'Content-Type': 'application/json' } }
    ).then((resp) => {
      isSubmitted = false;

      const body: HTTPResponse = resp.data as HTTPResponse;
      const out: registerOut = body.data as registerOut;

      toast.success(out.message);
      setTimeout(() => ModeHostState.set(ModeHostLogin), 1200);
    }).catch((err) => {
      isSubmitted = false;

      const body: HTTPResponse = err.response.data as HTTPResponse;

      toast.error(body.errors);
      console.log(err);
    });
  }
</script>

<Toaster />

<div class="w-[400px] bg-zinc-900 p-5 rounded-xl mx-auto flex flex-col gap-7">
  <header class="flex flex-row justify-between items-center">
    <h1 class="text-3xl font-bold text-end">Register <span class="text-emerald-600">&gt;</span></h1>
    <a href="/" use:link class="focus:outline-none flex justify-center items-center p-2 bg-zinc-800 rounded-lg group hover:bg-violet-500/10">
      <Icon src={TrOutlineHome} size="20" className="group-hover:stroke-violet-500" />
    </a>
  </header>
  <form class="flex flex-col gap-4">
    <InputBox
      id="email"
      type="email"
      bind:value={email}
      placeholder="john@example.com"
      label="Email"
    />
    <InputBox
      id="username"
      type="text"
      bind:value={username}
      placeholder="john_doe"
      label="Username"
    />
    <InputBox
      id="full_name"
      type="text"
      bind:value={full_name}
      placeholder="John Doe"
      label="Full Name"
    />
    <InputBox
      id="password"
      type="password"
      bind:value={password}
      placeholder="password1234"
      label="Password"
    />
    <div class="flex flex-row gap-1 justify-center text-sm">
      <span>Already have an account?</span>
      <button
        class="text-sky-500 hover:underline focus:outline-none"
        on:click|preventDefault={() => ModeHostState.set(ModeHostLogin)}
      >Login</button>
    </div>
    <button
      disabled={isSubmitted}
      class="focus:outline-none bg-violet-500 disabled:cursor-not-allowed py-2.5 rounded-lg font-semibold
        text-white hover:bg-violet-500/70 disabled:bg-zinc-800"
      on:click|preventDefault={Register}
    >
      {#if !isSubmitted}
        Register
      {/if}
      {#if isSubmitted}
        Register...
      {/if}
    </button>
  </form>
</div>