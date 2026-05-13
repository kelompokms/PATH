<script>
    import CircleUser from "$lib/svg/circle-user.svelte";

    const MURID = ["John Doe", "Foo", "Bar"];
</script>

<div class="p-2">
    <div class="mb-6">
        <h3
            class="font-semibold text-2xl bg-purple-200 p-4 border-2 border-purple-900/10 rounded-t-lg"
        >
            Pengajar
        </h3>
        <p class="p-4 border-2 border-t-0 border-black/15 rounded-b-lg">
            Vereh
        </p>
    </div>
    <div>
        <h3
            class="font-semibold text-2xl bg-purple-200 p-4 border-2 border-purple-900/15 rounded-t-lg"
        >
            Teman Sekelas
        </h3>
        <div>
            {#each MURID as item}
                <div
                    class="p-4 border-b-2 border-black/15 last:rounded-b-lg border-x-2 flex gap-2 items-center *:first:size-8 *:first:bg-purple-200 *:first:rounded-full"
                >
                    <CircleUser />
                    <p>{item}</p>
                </div>
            {/each}
        </div>
    </div>
</div>
