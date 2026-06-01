<script>
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";

    let { params } = $props();

    onMount(() => {
        goto("/kelas/" + params.kode + "/forum");
    });
</script>
