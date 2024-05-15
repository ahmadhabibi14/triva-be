import { writable, type Writable } from 'svelte/store';

export const ModeHostPassed: string = 'passed';
export const ModeHostLogin: string = 'login';
export const ModeHostRegister: string = 'register';
export const ModeHostState: Writable<string> = writable(ModeHostLogin);