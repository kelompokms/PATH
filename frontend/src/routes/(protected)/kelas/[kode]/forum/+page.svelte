<script>
    import { goto } from "$app/navigation";
    import Pencil from "$lib/svg/pencil.svelte";
    import { get } from "$lib/utils/api";
    import { onMount } from "svelte";

    let { data, params } = $props();

    let posts = $state([]);

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
    {#if !posts}
        <div
            class="p-4 grow content-center justify-center bg-white border-2 border-black/10 shadow-lg rounded-lg font-bold text-center text-black/50 flex flex-col gap-2"
        >
            Belum ada postingan.
            <button
                onclick={() => goto("buat")}
                class="btn btn-primary max-w-sm w-full mx-auto">Buat</button
            >
        </div>
    {:else}
        <button
            onclick={() => goto("buat")}
            class="btn btn-primary *:first:size-4 mb-4"
        >
            Buat Postingan
        </button>
    {/if}
    {#each posts as post}
        <div>
            <p
                class={`p-4 rounded-t-lg border-2 border-purple-900/15 border-b-0 font-semibold text-2xl ${post.Tipe == "tugas" ? "bg-blue-200" : "bg-purple-200"}`}
            >
                {post.Nama}
            </p>
            <div
                class="p-4 border-2 border-t-0 border-black/10 bg-white rounded-b-lg"
            >
                <p>
                    {post.Deskripsi}
                </p>
                {#if post.Tenggat}
                    <p class="text-red-800/60 text-sm mt-2">
                        Deadline: {new Date(post.Tenggat).toLocaleString(
                            "id-ID",
                            { dateStyle: "medium", timeStyle: "short" },
                        )}
                    </p>
                {/if}
            </div>
        </div>
    {/each}
</div>
