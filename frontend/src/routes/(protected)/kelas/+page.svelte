<script>
    import { goto } from "$app/navigation";
    import Placeholder from "$lib/assets/placeholder_bg.webp";
    import { get } from "$lib/utils/api";
    import { onMount } from "svelte";

    let kelas = $state();

    onMount(async () => {
        kelas = await get("class");
    });
</script>

{#if typeof kelas == "undefined"}
    <main class="flex flex-wrap gap-6 justify-center h-fit">
        <div class="skeleton w-full max-w-sm aspect-video h-fit"></div>
        <div class="skeleton w-full max-w-sm aspect-video h-fit"></div>
        <div class="skeleton w-full max-w-sm aspect-video h-fit"></div>
        <div class="skeleton w-full max-w-sm aspect-video h-fit"></div>
    </main>
{:else if typeof kelas == "object" && kelas.length <= 0}
    <main class="flex flex-col justify-center items-center">
        <div class="flex flex-col max-w-sm gap-4">
            <div
                class="flex flex-col p-10 gap-6 bg-white rounded-lg shadow-lg w-full text-center"
            >
                <img src={Placeholder} class="rounded-full aspect-square" alt="" />
                <h3 class="font-semibold text-xl text-sky-900">
                    Tidak ada Kelas
                </h3>
            </div>
            <button onclick={() => goto("/kelas/buat")} class="btn btn-primary"
                >Buat</button
            >
        </div>
    </main>
{:else}
    <main class="flex flex-wrap gap-4 p-0 md:p-4 h-fit">
        {#each kelas as i}
            <a
                href={"/kelas/" + i.Kode + "/forum"}
                class="card rounded-md w-100 h-min text-left shadow-md transition-all bg-purple-200 hover:bg-purple-300 active:bg-purple-400"
            >
                <figure>
                    <img src={Wolf} alt="" class="aspect-video" />
                </figure>
                <div class="card-body p-3 border-2 border-purple-900/10 gap-1">
                    <h2 class="card-title">{i.NamaKelas}</h2>
                    <p>{i.NamaPengguna}</p>
                </div>
            </a>
        {/each}
    </main>
{/if}
