<script>
    import { get } from "$lib/utils/api";
    import { onMount } from "svelte";

    let { data, params } = $props();

    const kelasItems = [
        {
            id: 1,
            nama: "Materi 1",
            deskripsi:
                "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Modi, iusto unde sint aspernatur, ducimus illum ipsam maiores voluptate ab perspiciatis blanditiis voluptatibus veniam similique. Molestiae itaque magni dolorum vel ducimus.",
            kode_kelas: 1,
            tipe: "materi",
        },
        {
            id: 2,
            nama: "Materi 2",
            deskripsi:
                "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Modi, iusto unde sint aspernatur, ducimus illum ipsam maiores voluptate ab perspiciatis blanditiis voluptatibus veniam similique. Molestiae itaque magni dolorum vel ducimus.",
            kode_kelas: 1,
            tipe: "materi",
        },
        {
            id: 3,
            nama: "Kuis 1",
            deskripsi:
                "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Modi, iusto unde sint aspernatur, ducimus illum ipsam maiores voluptate ab perspiciatis blanditiis voluptatibus veniam similique. Molestiae itaque magni dolorum vel ducimus.",
            kode_kelas: 1,
            tipe: "quiz",
        },
    ];
</script>

<div class="flex flex-col gap-4 p-2 max-w-7xl mx-auto">
    {#if data}
        <div class="border-2 border-black/10 p-4 rounded-lg shadow-lg">
            <h2 class="font-bold text-2xl">{data.NamaKelas}</h2>
            <p class="mt-2">{data.Bagian}</p>
            <button
                onclick={(event) => {
                    navigator.clipboard.writeText(data.Kode);
                    alert("Kode kelas berhasil disalin");
                }}
                class="btn btn-ghost mt-4 font-bold text-xl text-black/50"
            >
                {data.Kode}
            </button>
        </div>
    {:else}
        <div class="border-2 border-black/10 p-4 rounded-lg shadow-lg">
            <div class="font-bold text-2xl skeleton h-4"></div>
            <div class="mt-4 h-8 skeleton"></div>
        </div>
    {/if}
    {#each kelasItems as item}
        <div>
            <p
                class="p-4 rounded-t-lg border-2 border-purple-900/15 border-b-0 font-semibold text-2xl bg-purple-200"
            >
                {item.nama}
            </p>
            <p
                class="p-4 border-2 border-t-0 border-black/10 bg-white rounded-b-lg"
            >
                {item.deskripsi}
            </p>
        </div>
    {/each}
</div>
