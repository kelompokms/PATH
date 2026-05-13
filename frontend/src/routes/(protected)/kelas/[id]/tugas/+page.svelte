<script>
    import AngleRight from "$lib/svg/angle-right.svelte";
    import Pencil from "$lib/svg/pencil.svelte";

    const tugasItems = [
        { id: 1, nama: "Tugas 1", deadline: "2026-05-05" },
        { id: 2, nama: "Tugas 2", deadline: "2026-05-05" },
        { id: 3, nama: "Tugas 3", deadline: "2026-05-05" },
        { id: 4, nama: "Tugas 4", deadline: "2026-05-05" },
        { id: 5, nama: "Tugas 5", deadline: "2026-05-05" },
        { id: 6, nama: "Tugas 6", deadline: "2026-05-05" },
    ];
</script>

<div class="p-2">
    <div>
        <h3 class="p-4 bg-purple-200 font-semibold text-2xl rounded-t-lg">
            Tugas
        </h3>
        {#each tugasItems as item}
            <div
                class="flex items-center gap-4 p-4 border-2 border-t-0 border-black/10 last:rounded-b-lg"
            >
                <button
                    class="btn btn-secondary rounded-full p-2 *:first:size-6"
                >
                    <Pencil />
                </button>
                <p class="font-semibold grow">{item.nama}</p>
                <p>
                    <strong>Tenggat:</strong>
                    {item.deadline}
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
