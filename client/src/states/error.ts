import { writable, type Writable } from 'svelte/store';

export const errorState: Writable<string> = writable('');