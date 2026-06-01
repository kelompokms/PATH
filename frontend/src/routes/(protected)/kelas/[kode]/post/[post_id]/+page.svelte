<script>
    import { goto } from "$app/navigation";
    import { PUBLIC_API_URL } from "$env/static/public";
    import AngleLeft from "$lib/svg/angle-left.svelte";
    import { get } from "$lib/utils/api";
    import { onMount } from "svelte";

    let { params } = $props();

    let post = $state();

    onMount(async () => {
        const res = await get(
            "class/" + params.kode + "/post/" + params.post_id,
        );
        post = res;
    });
</script>

{#if post == undefined || post == null}
    <div
        class="overflow-auto max-w-6xl z-5 md:py-4 relative w-full h-full mx-auto"
    >
        <div class=" bg-white p-4 pcard h-full flex flex-col gap-4">
            <button
                onclick={() => goto("/kelas/" + params.kode + "/forum")}
                class="*:first:size-6 btn btn-ghost p-2 justify-center items-center self-start"
            >
                <AngleLeft />
                Kembali
            </button>
            <div class="h-12 skeleton"></div>
            <div class="h-32 skeleton"></div>
            <div class="h-32 skeleton"></div>
            <div class="h-32 skeleton"></div>
        </div>
    </div>
{:else}
    <div
        class="overflow-auto max-w-6xl z-5 md:py-4 relative w-full h-full mx-auto"
    >
        <div class=" bg-white p-4 pcard h-full flex flex-col gap-2">
            <button
                onclick={() => goto("/kelas/" + params.kode + "/forum")}
                class="*:first:size-6 btn btn-ghost p-2 justify-center items-center self-start"
            >
                <AngleLeft />
                Kembali
            </button>
            <div class="flex justify-between items-center">
                <h2 class="font-bold text-2xl">
                    {post.Nama}
                </h2>
                <p
                    class={`font-semibold ${post.Tipe == "tugas" ? "text-blue-900" : "text-purple-900"}`}
                >
                    {post.Tipe}
                </p>
            </div>
            <p class="italic text-sm text-black/80 mb-2">
                Dibuat pada {new Date(post.Dibuat).toLocaleString("id-ID", {
                    dateStyle: "medium",
                    timeStyle: "short",
                })}
            </p>
            <p>{post.Deskripsi}</p>
            {#if post.Tenggat}
                <div class="absolute bottom-0 left-0 w-full">
                    <div class="p-4 max-w-7xl mx-auto md:pb-8">
                        <p class="text-red-900 mb-4 text-center">
                            Tenggat Pengerjaan:
                            {new Date(post.Dibuat).toLocaleString("id-ID", {
                                dateStyle: "medium",
                                timeStyle: "short",
                            })}
                        </p>
                        <button
                            onclick={(event) => {
                                event.target.disabled = true;
                                window.location.href =
                                    "https://youtu.be/QDia3e12czc?si=Q3j0nePGrsr4-ubU";
                            }}
                            class="btn btn-primary w-full">Kumpulkan</button
                        >
                    </div>
                </div>
            {/if}
            {#each post.File as file}
                <a
                    class="btn btn-outline w-fit"
                    target="_blank"
                    href={PUBLIC_API_URL + file.slice(1)}
                >
                    {file.split("/").at(-1)}</a
                >
            {/each}
        </div>
    </div>
{/if}
