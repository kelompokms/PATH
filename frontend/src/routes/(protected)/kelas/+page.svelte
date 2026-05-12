<script>
    import Wolf from "$lib/assets/wolf.webp";
    let kelas = [
        {
            ID: 1,
            NamaKelas: "Mapel 1",
            Subjek: "Jaringan Nirkabel",
            NamaPengguna: "Vereh",
            Kode: "6969",
            Dibuat: "2026-03-03",
        },
        {
            ID: 2,
            NamaKelas: "Mapel 2",
            Subjek: "Jaringan Nirkabel",
            NamaPengguna: "Vereh",
            Kode: "6969",
            Dibuat: "2026-03-03",
        },
        {
            ID: 3,
            NamaKelas: "Mapel 3",
            Subjek: "Jaringan Nirkabel",
            NamaPengguna: "Vereh",
            Kode: "6969",
            Dibuat: "2026-03-03",
        },
        {
            ID: 4,
            NamaKelas: "Mapel 4",
            Subjek: "Jaringan Nirkabel",
            NamaPengguna: "Vereh",
            Kode: "6969",
            Dibuat: "2026-03-03",
        },
    ];
</script>

<main class="flex flex-wrap gap-8 p-0 md:p-4">
    {#each kelas as i}
        <a
            href={"/kelas/" + i.ID + "/forum"}
            class="card w-100 h-fit text-left shadow-sm transition-all bg-purple-200 hover:bg-purple-300 active:bg-purple-400"
        >
            <figure>
                <img src={Wolf} alt="" class="aspect-video" />
            </figure>
            <div class="card-body p-4">
                <h2 class="card-title">{i.NamaKelas}</h2>
                <p>{i.NamaPengguna}</p>
            </div>
        </a>
    {/each}
</main>
