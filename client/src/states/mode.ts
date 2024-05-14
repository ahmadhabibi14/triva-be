import { writable, type Writable } from 'svelte/store';

export const ModeHostView: Writable<string> = writable('login');