<script>
    import { goto } from "$app/navigation";
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
        class="h-screen w-screen absolute left-0 top-0 p-4 flex flex-col gap-4 z-3"
    >
        <div class="skeleton h-12"></div>
        <div class="skeleton h-32"></div>
        <div class="skeleton h-32"></div>
        <div class="absolute left-0 bottom-0 w-full">
            <div class="p-4">
                <div class="skeleton h-12 w-full"></div>
            </div>
        </div>
    </div>
{:else}
    <div
        class="h-screen w-screen overflow-auto absolute left-0 top-0 z-3 md:py-4"
    >
        <div
            class="md:max-w-7xl mx-auto bg-white p-4 pcard h-full flex flex-col gap-2"
        >
            <button
                onclick={() => goto("/kelas/" + params.kode + "/post")}
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
        </div>
    </div>
{/if}
