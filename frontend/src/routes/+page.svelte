<script>
    import { goto } from "$app/navigation";
    import { checkAuth } from "$lib/utils/api";
    import { onMount } from "svelte";

    onMount(async () => {
        const res = await checkAuth();
        console.log(res);
        if (res.ok) {
            goto("/beranda");
            return;
        }

        goto("/login");
    });
</script>
