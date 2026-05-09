<script>
    import DrawerItem from "$lib/components/DrawerItem.svelte";
    import Bars from "$lib/svg/bars.svelte";
    import Calendar from "$lib/svg/calendar.svelte";
    import ChalkboardUser from "$lib/svg/chalkboard-user.svelte";
    import Chalkboard from "$lib/svg/chalkboard.svelte";
    import CircleUser from "$lib/svg/circle-user.svelte";
    import Gear from "$lib/svg/gear.svelte";
    import House from "$lib/svg/house.svelte";
    import Question from "$lib/svg/question.svelte";

    let { children } = $props();
</script>

<div class="drawer lg:drawer-open h-screen">
    <input id="drawer-toggle" type="checkbox" class="drawer-toggle" />
    <div class="drawer-content flex flex-col h-screen">
        <!-- Navbar -->
        <nav class="navbar w-full">
            <label
                for="drawer-toggle"
                aria-label="open sidebar"
                class="btn btn-square btn-ghost"
            >
                <Bars />
            </label>
            <div class="px-4">
                <h1 class="text-3xl font-bold">PATH</h1>
            </div>
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
                <DrawerItem icon={House} name="Beranda" />
                <DrawerItem icon={Chalkboard} name="Kelas" />
                <DrawerItem icon={ChalkboardUser} name="Pengajar" />
                <DrawerItem icon={Calendar} name="Kalender" />
                <DrawerItem icon={Gear} name="Pengaturan" />
                <DrawerItem icon={Question} name="Bantuan" />
                <DrawerItem icon={CircleUser} name="Profil" />
            </ul>
        </div>
    </div>
</div>
