<script>
    import { onNavigate } from "$app/navigation";
    import { onMount } from "svelte";

    const navItems = [
        { href: "forum", nama: "Forum" },
        { href: "anggota", nama: "Anggota" },
        { href: "tugas", nama: "Tugas Kelas" },
        { href: "chat", nama: "Chat AI" },
    ];
    const regex = /post|buat/;

    let { children, params } = $props();
    let active = $state("");

    let navActive = $state(true);

    onMount(() => {
        const pathname = window.location.pathname.split("/");
        active = pathname[pathname.length - 1];

        navActive = !regex.test(window.location.pathname);
    });

    onNavigate(() => {
        navActive = !regex.test(window.location.pathname);
    });
</script>

<main class="flex flex-col">
    {#if navActive}
        <div role="tablist" class="tabs tabs-border p-2">
            {#each navItems as item}
                <a
                    onclick={() => {
                        active = item.href;
                    }}
                    role="tab"
                    class={`tab ${active == item.href ? "tab-active" : ""}`}
                    href={"/kelas/" + params.kode + "/" + item.href}
                >
                    {item.nama}
                </a>
            {/each}
        </div>
    {/if}
    {@render children()}
</main>
