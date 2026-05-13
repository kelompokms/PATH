<script>
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

<div class="flex flex-col gap-4 p-2">
    {#each kelasItems as item}
        <div>
            <p
                class="p-4 rounded-t-lg border-2 border-purple-900/15 border-b-0 font-semibold text-2xl bg-purple-200"
            >
                {item.nama}
            </p>
            <p class="p-4 border-2 border-t-0 border-black/10 rounded-b-lg">
                {item.deskripsi}
            </p>
        </div>
    {/each}
</div>
