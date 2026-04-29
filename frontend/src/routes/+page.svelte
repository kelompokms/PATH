<script>
	import { get } from '$lib/utilitas/api';
	import { onMount } from 'svelte';

	onMount(() => {
		fetch('auth').then((res) => {
			if (!res.ok) {
				window.location.href = '/login';
				return;
			}
			window.location.href = '/dashboard';
		});
	});
</script>
