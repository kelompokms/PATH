<script>
    import { get, post } from "$lib/utils/api";

    let code = $state("");

    let kelas = $state();

    $effect(() => {
        if (code.length != 6) {
            kelas = undefined;
            return;
        }
        setTimeout(async () => {
            const res = await get("class/" + code);
            if (res) {
                kelas = res;
            }
        }, 500);
    });
</script>

<main class="flex justify-center items-center">
    <div
        class="bg-white flex flex-col gap-4 p-4 shadow-lg rounded-lg border-2 border-black/10 w-full sm:w-sm"
    >
        <input
            bind:value={code}
            class="p-2"
            placeholder="Masukkan Kode Kelas"
            type="text"
            required
        />
        {#if typeof kelas != "undefined"}
            <div class="border border-black/20 p-2">
                <p class="text-2xl font-bold text-sky-900">{kelas.NamaKelas}</p>
                <p class="text-slate-600 text-sm">{kelas.Bagian}</p>
                <p class="text-slate-600 text-sm mb-2">
                    Pengajar: {kelas.NamaPengajar}
                </p>
                <p class="">{kelas.Deskripsi}</p>
            </div>
        {/if}
        <button
            disabled={code == "" && kelas == undefined}
            onclick={() => {
                post("class/" + code + "/join", {});
            }}
            class="p-2 btn btn-secondary">Gabung</button
        >
    </div>
</main>
