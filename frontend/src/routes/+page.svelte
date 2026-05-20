<script>
    import { goto } from "$app/navigation";
    import { checkAuth } from "$lib/utils/api";
    import { onMount } from "svelte";

    onMount(async () => {
        const res = await checkAuth();
        if (res.ok) {
            goto("/beranda");
        }

        goto("/login");
    });
</script>
