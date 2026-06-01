<script>
    import { goto } from "$app/navigation";
    import Pencil from "$lib/svg/pencil.svelte";
    import { get } from "$lib/utils/api";
    import { onMount } from "svelte";

    import ClassBG from "$lib/assets/placeholder_bg.webp";

    let { data, params } = $props();

    let posts = $state();

    onMount(async () => {
        const res = await get("class/" + params.kode + "/post");
        if (!res) return;

        posts = res;
    });
</script>

<div
    class="flex flex-col gap-4 p-2 max-w-7xl w-full mx-auto grow overflow-auto"
>
    {#if data}
        <div class="flex flex-col bg-purple-50">
            <div
                class="bg-purple-200 p-4 rounded-t-md border-2 border-purple-900/20 flex justify-between items-center"
            >
                <div>
                    <h2 class="text-3xl font-bold">{data.NamaKelas}</h2>
                    <p class="mt-2 font-semibold">{data.Bagian}</p>
                </div>
                <button
                    onclick={(event) => {
                        navigator.clipboard.writeText(data.Kode);
                        alert("Kode kelas berhasil disalin");
                    }}
                    class="btn btn-ghost hover:bg-purple-900/10 active:bg-purple-900/20"
                    >{data.Kode}</button
                >
            </div>
            <img
                src={ClassBG}
                class="h-32 object-cover rounded-b-md border-2 border-black/20 border-t-0"
                alt="class background"
            />
        </div>
        <!-- <div
            class="flex flex-col border-2 rounded-md border-black/20 bg-purple-50"
        >
            <h2 class="bg-purple-200 p-4 text-2xl font-semibold rounded-t-md">
                {data.NamaKelas}
            </h2>
            <p class="p-4">{data.Bagian}</p>
            <button
                onclick={(event) => {
                    navigator.clipboard.writeText(data.Kode);
                    alert("Kode kelas berhasil disalin");
                }}
                class="btn btn-ghost mt-4 font-bold text-xl text-black/50"
            >
                {data.Kode}
            </button>
            {#if data.IsPengajar}
                <button
                    onclick={() => goto("/kelas/" + params.kode + "/buat")}
                    class="btn btn-primary border-0 rounded-t-0"
                >
                    Buat Postingan
                </button>
            {/if}
        </div> -->
    {:else}
        <div class="border-2 border-black/10 p-4 rounded-md shadow-md">
            <div class="font-bold text-2xl skeleton h-4"></div>
            <div class="mt-4 h-8 skeleton"></div>
        </div>
    {/if}
    {#if !posts}
        {#each Array(3) as i}
            <div class="skeleton h-32"></div>
        {/each}
    {:else if posts.length <= 0}
        <div
            class="p-4 grow content-center justify-center bg-white border-2 border-black/10 shadow-md rounded-md font-bold text-center text-black/50 flex flex-col gap-2"
        >
            Belum ada postingan.
            {#if data.IsPengajar}
                <button
                    onclick={() => goto("buat")}
                    class="btn btn-primary max-w-sm w-full mx-auto">Buat</button
                >
            {/if}
        </div>
    {/if}
    {#each posts as post}
        <button
            onclick={() => goto("/kelas/" + params.kode + "/post/" + post.ID)}
            class="text-left border-2 border-transparent active:border-black/40 hover:border-black/20 rounded-md"
        >
            <p
                class={`p-4 rounded-t-md border-2 border-purple-900/15 border-b-0 font-semibold text-lg md:text-2xl ${post.Tipe == "tugas" ? "bg-primary" : "bg-secondary"}`}
            >
                {post.Nama}
            </p>
            <div
                class="p-4 border-2 border-t-0 border-black/10 bg-white rounded-b-md"
            >
                <p>
                    {post.Deskripsi}
                </p>
                {#if post.Tenggat}
                    <p class="text-red-900 text-sm mt-2">
                        Deadline: {new Date(post.Tenggat).toLocaleString(
                            "id-ID",
                            { dateStyle: "medium", timeStyle: "short" },
                        )}
                    </p>
                {/if}
            </div>
        </button>
    {/each}
</div>
