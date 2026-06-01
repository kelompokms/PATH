<script>
    import { goto } from "$app/navigation";
    import CircleUser from "$lib/svg/circle-user.svelte";
    import { get, post } from "$lib/utils/api";
    import { onMount } from "svelte";

    let { data } = $props();

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
    <div class="bg-white">
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
                    <p>{data.Nama}</p>
                </div>
                <button class="btn btn-ghost">Edit</button>
            </div>
            <div>
                <div>
                    <h3 class="font-semibold text-lg">Email</h3>
                    <p>{data.Email}</p>
                </div>
                <button class="btn btn-ghost">Edit</button>
            </div>
            <div>
                <div>
                    <h3 class="font-semibold text-lg">Phone</h3>
                    <p>{data.Telepon}</p>
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
                        {data.JenisKelamin == "l" ? "Laki-laki" : "Perempuan"}
                    </p>
                </div>
                <button class="btn btn-ghost">Edit</button>
            </div>
        </div>
    </div>
    <button onclick={logout} class="mt-4 btn btn-error btn-outline ml-auto"
        >Logout</button
    >
</main>
