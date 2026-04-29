<script>
	import ArrowRight from '$lib/assets/ArrowRight.svelte';
	import Bars3 from '$lib/assets/Bars3.svelte';
	import PlusCircle from '$lib/assets/PlusCircle.svelte';
	import UserCircle from '$lib/assets/UserCircle.svelte';
	import { onMount } from 'svelte';

	let { children } = $props();
	let navItems = $state();

	onMount(() => {
		navItems = window.location.pathname.split('/').filter((val) => val != '');
	});
</script>

<div class="flex flex-col">
	<header class="flex items-center gap-2 p-3 text-center md:gap-4">
		<checkbox class="rounded p-2 hover:bg-black/5 active:bg-black/10">
			<Bars3 />
		</checkbox>
		<a
			href="/beranda"
			class="mb-1 rounded p-2 text-2xl font-black hover:bg-black/5 active:bg-black/10">PATH</a
		>
		{#each navItems as item}
			<ArrowRight />
			<p clpss="overflow-auto text-sm md:text-lg">{item}</p>
		{/each}
		<div class="grow"></div>
		<button class="rounded-full bg-purple-200 hover:bg-purple-300 active:bg-purple-400">
			<PlusCircle />
		</button>
		<button class="rounded-full bg-purple-200 hover:bg-purple-300 active:bg-purple-400">
			<UserCircle />
		</button>
	</header>
	<main class="flex grow gap-4">
		<div class="">Lorem</div>
		<div class="grow overflow-auto">
			{@render children()}
		</div>
	</main>
</div>
