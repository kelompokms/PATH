<script>
	import ArrowRight from '$lib/assets/ArrowRight.svelte';
	import Bars3 from '$lib/assets/Bars3.svelte';
	import Calendar from '$lib/assets/Calendar.svelte';
	import Gear from '$lib/assets/Gear.svelte';
	import Home from '$lib/assets/Home.svelte';
	import PlusCircle from '$lib/assets/PlusCircle.svelte';
	import Question from '$lib/assets/Question.svelte';
	import Teacher from '$lib/assets/Teacher.svelte';
	import UserCircle from '$lib/assets/UserCircle.svelte';
	import UserGroup from '$lib/assets/UserGroup.svelte';
	import { onMount } from 'svelte';

	let { children } = $props();
	let sidebarActive = $state(false);
	let navItems = $state();

	onMount(() => {
		navItems = window.location.pathname.split('/').filter((val) => val != '');
	});
</script>

<div class="flex flex-col">
	<header class="flex items-center gap-2 p-3 text-center md:gap-4">
		<button
			class="rounded p-2 hover:bg-black/5 active:bg-black/10"
			onclick={(sidebarActive = !sidebarActive)}
		>
			<Bars3 />
		</button>
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
	<main class="flex grow gap-4 bg-purple-100">
		<div
			class={`fixed h-full bg-white p-4 shadow-md transition-all md:block ${sidebarActive ? 'block' : 'hidden'}`}
		>
			<nav class="flex flex-col gap-4">
				<a href="/beranda" class="flex items-center gap-2">
					<Home />
					<p class={`${sidebarActive ? 'block' : 'hidden'}`}>Beranda</p>
				</a>
				<a href="/kelas" class="flex items-center gap-2">
					<UserGroup />
					<p class={`${sidebarActive ? 'block' : 'hidden'}`}>Kelas</p>
				</a>
				<a href="/pengajar" class="flex items-center gap-2">
					<Teacher />
					<p class={`${sidebarActive ? 'block' : 'hidden'}`}>Pengajar</p>
				</a>
				<a href="/kalender" class="flex items-center gap-2">
					<Calendar />
					<p class={`${sidebarActive ? 'block' : 'hidden'}`}>Kalender</p>
				</a>
				<a href="/pengaturan" class="flex items-center gap-2">
					<Gear />
					<p class={`${sidebarActive ? 'block' : 'hidden'}`}>Pengaturan</p>
				</a>
			</nav>
		</div>
		<div class="grow overflow-auto">
			{@render children()}
		</div>
	</main>
</div>
