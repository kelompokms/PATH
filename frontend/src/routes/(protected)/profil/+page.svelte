<script>
    import CircleUser from "$lib/svg/circle-user.svelte";
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
                    <p>Siswa 1</p>
                </div>
                <button class="btn btn-ghost">Edit</button>
            </div>
            <div>
                <div>
                    <h3 class="font-semibold text-lg">Email</h3>
                    <p>siswa@mail.com</p>
                </div>
                <button class="btn btn-ghost">Edit</button>
            </div>
            <div>
                <div>
                    <h3 class="font-semibold text-lg">Phone</h3>
                    <p>085123123</p>
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
        </div>
    </div>
</main>
