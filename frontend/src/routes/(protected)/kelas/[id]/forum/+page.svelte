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
        <div class="border p-2">
            <p class="font-semibold text-xl">{item.nama}</p>
            <p>{item.deskripsi}</p>
        </div>
    {/each}
</div>
