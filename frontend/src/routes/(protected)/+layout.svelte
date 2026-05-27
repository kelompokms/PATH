<script>
    import { goto, onNavigate } from "$app/navigation";
    import DrawerItem from "$lib/components/DrawerItem.svelte";
    import AngleRight from "$lib/svg/angle-right.svelte";
    import Bars from "$lib/svg/bars.svelte";
    import Calendar from "$lib/svg/calendar.svelte";
    import ChalkboardUser from "$lib/svg/chalkboard-user.svelte";
    import Chalkboard from "$lib/svg/chalkboard.svelte";
    import CircleUser from "$lib/svg/circle-user.svelte";
    import Gear from "$lib/svg/gear.svelte";
    import House from "$lib/svg/house.svelte";
    import Question from "$lib/svg/question.svelte";
    import { checkAuth } from "$lib/utils/api";
    import { redirect } from "@sveltejs/kit";
    import { onMount } from "svelte";

    let { children } = $props();
    let isAuth = $state(false);
    let breadcrumbs = $state([]);

    onMount(async () => {
        const res = await checkAuth();
        if (!res.ok) goto("/login");
        isAuth = true;

        breadcrumbs = getNavData(window.location.pathname);
    });

    onNavigate((navigation) => {
        breadcrumbs = getNavData(navigation.to.url.pathname);
    });

    function getNavData(pathname) {
        const pathArr = pathname.split("/").slice(1);
        let pathData = [];
        let tempPath = "";

        for (let i = 0; i < pathArr.length; i++) {
            tempPath += "/" + pathArr[i];
            pathData.push({ name: pathArr[i], href: tempPath });
        }
        return pathData;
    }
</script>

{#if isAuth}
    <div class="drawer lg:drawer-open h-screen">
        <input id="drawer-toggle" type="checkbox" class="drawer-toggle" />
        <div class="drawer-content flex flex-col h-screen">
            <!-- Navbar -->
            <nav class="navbar w-full shadow-md/5">
                <label
                    for="drawer-toggle"
                    aria-label="open sidebar"
                    class="btn btn-square btn-secondary btn-ghost"
                >
                    <Bars />
                </label>
                <button onclick={() => goto("/")} class="btn btn-ghost px-2">
                    <h1 class="text-3xl font-bold">PATH</h1>
                </button>
                {#each breadcrumbs as item}
                    <div class="*:first:size-6 flex items-center">
                        <AngleRight />
                        <button
                            onclick={() => goto(item.href)}
                            class="btn btn-ghost p-1">{item.name}</button
                        >
                    </div>
                {/each}
            </nav>

            <!-- Content -->
            <div class="p-4 grow bg-purple-50 overflow-auto">
                {@render children()}
            </div>
        </div>

        <div class="drawer-side is-drawer-close:overflow-visible">
            <label
                for="drawer-toggle"
                aria-label="close sidebar"
                class="drawer-overlay"
            ></label>
            <div
                class="flex min-h-full flex-col items-start bg-base-200 is-drawer-close:w-14 is-drawer-open:w-64"
            >
                <ul class="menu w-full grow">
                    <!-- List item -->
                    <DrawerItem href="/beranda" icon={House} name="Beranda" />
                    <DrawerItem href="/kelas" icon={Chalkboard} name="Kelas" />
                    <DrawerItem
                        href="/pengajar"
                        icon={ChalkboardUser}
                        name="Pengajar"
                    />
                    <DrawerItem
                        href="/kalender"
                        icon={Calendar}
                        name="Kalender"
                    />
                    <DrawerItem
                        href="/pengaturan"
                        icon={Gear}
                        name="Pengaturan"
                    />
                    <DrawerItem
                        href="/bantuan"
                        icon={Question}
                        name="Bantuan"
                    />
                    <DrawerItem
                        href="/profil"
                        icon={CircleUser}
                        name="Profil"
                    />
                </ul>
            </div>
        </div>
    </div>
{/if}
