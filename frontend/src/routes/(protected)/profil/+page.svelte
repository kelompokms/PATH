<script>
    import { goto } from "$app/navigation";
    import CircleUser from "$lib/svg/circle-user.svelte";
    import { get, post } from "$lib/utils/api";
    import { onMount } from "svelte";

    let user = $state();

    onMount(async () => {
        const res = await get("user");
        user = res;
    });

    async function logout() {
        const res = await post("logout", {});
        if (res.ok) {
            goto("/login");
            return;
        }

        const text = res.text();
        alert("Gagal Logout:", text);
    }
</script>

<main class="p-0 md:p-4">
    {#if user}
        <div class="bg-white max-w-7xl mx-auto">
            <h2
                class="text-2xl font-semibold p-4 bg-purple-200 border-2 rounded-t-lg border-purple-900/10"
            >
                Profil
            </h2>
            <div
                class="*:flex *:justify-between *:items-center *:p-4 *:border-2 *:border-black/10 *:border-t-0 *:last:rounded-b-lg"
            >
                <div>
                    <div
                        class="flex grow *:first:size-12 *:first:bg-purple-200 *:first:rounded-full gap-2 items-center"
                    >
                        <CircleUser />
                        <p>Foto Profil</p>
                    </div>
                    <button class="btn btn-ghost">Edit</button>
                </div>
                <div>
                    <div>
                        <h3 class="font-semibold text-lg">Display Name</h3>
                        <p>{user.Nama}</p>
                    </div>
                    <button class="btn btn-ghost">Edit</button>
                </div>
                <div>
                    <div>
                        <h3 class="font-semibold text-lg">Email</h3>
                        <p>{user.Email}</p>
                    </div>
                    <button class="btn btn-ghost">Edit</button>
                </div>
                <div>
                    <div>
                        <h3 class="font-semibold text-lg">Phone</h3>
                        <p>{user.Telepon}</p>
                    </div>
                    <button class="btn btn-ghost">Edit</button>
                </div>
                <div>
                    <div>
                        <h3 class="font-semibold text-lg">Password</h3>
                        <p>********</p>
                    </div>
                    <button class="btn btn-ghost">Edit</button>
                </div>

                <div>
                    <div>
                        <h3 class="font-semibold text-lg">Jenis Kelamin</h3>
                        <p>
                            {user.JenisKelamin == "l"
                                ? "Laki-laki"
                                : "Perempuan"}
                        </p>
                    </div>
                    <button class="btn btn-ghost">Edit</button>
                </div>
            </div>
        </div>
        <div class="max-w-7xl mx-auto">
            <button onclick={logout} class="mt-4 btn btn-error btn-outline"
                >Logout</button
            >
        </div>
    {:else}
        <div class="flex flex-col gap-4 max-w-7xl mx-auto">
            <div class="w-full h-20 skeleton"></div>
            <div class="w-full h-20 skeleton"></div>
            <div class="w-full h-20 skeleton"></div>
            <div class="w-full h-20 skeleton"></div>
        </div>
    {/if}
</main>
