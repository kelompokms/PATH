<script>
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";

    onMount(() => {
        goto(window.location.pathname + "/forum");
    });
</script>
