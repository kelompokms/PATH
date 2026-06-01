<script>
    import CircleUser from "$lib/svg/circle-user.svelte";
    import { get } from "$lib/utils/api.js";
    import { onMount } from "svelte";

    let { data, params } = $props();

    let murid = $state([]);

    onMount(async () => {
        const res = await get("class/" + params.kode + "/murid");
        murid = res;
    });
</script>

<div class="max-w-7xl w-full mx-auto p-2">
    <div class="mb-6">
        <h3
            class="font-semibold text-2xl bg-purple-200 p-4 border-2 border-purple-900/10 rounded-t-lg"
        >
            Pengajar
        </h3>
        <p class="p-4 border-2 border-t-0 border-black/15 rounded-b-lg">
            {data.NamaPengajar}

            {#if data.IsPengajar}
                <span class="font-semibold">(saya)</span>
            {/if}
        </p>
    </div>
    {#if murid}
        <div>
            <h3
                class="font-semibold text-2xl bg-purple-200 p-4 border-2 border-purple-900/15 rounded-t-lg"
            >
                Teman Sekelas
            </h3>
            <div>
                {#each murid as item}
                    <div
                        class="p-4 border-b-2 border-black/15 last:rounded-b-lg border-x-2 flex gap-2 items-center *:first:size-8 *:first:bg-purple-200 *:first:rounded-full"
                    >
                        <CircleUser />
                        <p class="grow">{item.Nama}</p>
                        <p class="text-black/60">{item.Email}</p>
                    </div>
                {/each}
            </div>
        </div>
    {/if}
</div>
