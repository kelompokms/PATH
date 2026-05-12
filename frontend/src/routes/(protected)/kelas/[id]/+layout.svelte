<script>
    import { onMount } from "svelte";

    const navItems = [
        { href: "forum", nama: "Forum" },
        { href: "anggota", nama: "Anggota" },
        { href: "tugas", nama: "Tugas Kelas" },
        { href: "chat", nama: "Chat AI" },
    ];

    let { children } = $props();
    let active = $state("");

    onMount(() => {
        const pathname = window.location.pathname.split("/");
        active = pathname[pathname.length - 1];
    });
</script>

<main>
    <div role="tablist" class="tabs tabs-border p-2">
        {#each navItems as item}
            <a
                onclick={() => {
                    active = item.href;
                }}
                role="tab"
                class={`tab ${active == item.href ? "tab-active" : ""}`}
                href={item.href}
            >
                {item.nama}
            </a>
        {/each}
    </div>
    {@render children()}
</main>
