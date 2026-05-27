<script>
    import AngleRight from "$lib/svg/angle-right.svelte";
    import Pencil from "$lib/svg/pencil.svelte";
    import { get } from "$lib/utils/api";
    import { onMount } from "svelte";

    let { params } = $props();
    let tugas = $state([]);

    onMount(async () => {
        const res = await get("class/" + params.kode + "/tugas");
        if (!res) {
            return;
        }
        tugas = res;
    });
</script>

<div class="p-2 max-w-7xl w-full mx-auto">
    <div>
        <h3
            class="p-4 bg-purple-200 font-semibold text-2xl rounded-t-lg border-2 border-purple-900/10"
        >
            Tugas
        </h3>
        {#each tugas as item}
            <div
                class="flex items-center gap-4 p-4 border-2 border-t-0 border-black/10 last:rounded-b-lg"
            >
                <button
                    class="btn btn-secondary rounded-full p-2 *:first:size-6"
                >
                    <Pencil />
                </button>
                <p class="font-semibold grow">{item.Nama}</p>
                <p>
                    <strong>Tenggat:</strong>
                    {item.Tenggat}
                </p>
                <button
                    class="btn btn-secondary btn-square rounded-full *:first:size-6"
                >
                    <AngleRight />
                </button>
            </div>
        {/each}
    </div>
</div>
