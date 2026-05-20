<script>
    import { goto } from "$app/navigation";
    import Wolf from "$lib/assets/wolf.webp";
    import { post } from "$lib/utils/api";

    let isAgreed = $state(false);

    async function register(event) {
        event.preventDefault();

        let formData = new FormData(event.target);
        const res = await post("register", formData);

        if (res.ok) {
            goto("/beranda");
            return;
        }

        const text = await res.text();
        alert(text);
    }
</script>

<main
    class="bg-purple-100 flex flex-col justify-center items-center h-max md:h-screen"
>
    <div class="w-full max-w-md md:max-w-none text-center p-6">
        <div
            class="bg-white p-6 shadow-lg rounded-xl md:flex max-w-4xl md:mx-auto gap-4"
        >
            <img
                src={Wolf}
                alt="register"
                class="rounded-lg shadow-lg mb-4 md:mb-0 md:max-w-md"
            />
            <div class="flex flex-col gap-4 w-full">
                <h2 class="text-2xl font-bold">Daftar</h2>
                <form class="flex flex-col gap-6" onsubmit={register}>
                    <input placeholder="Nama Lengkap" type="text" name="nama" />
                    <input placeholder="Email" type="email" name="email" />
                    <input placeholder="Telepon" type="text" name="telepon" />
                    <select name="jenis_kelamin">
                        <option value="l">Laki-laki</option>
                        <option value="p">Perempuan</option>
                        <option value="x" disabled>Non-binary</option>
                        <option value="x" disabled>Transgender</option>
                        <option value="x" disabled>Gender Fluid</option>
                    </select>
                    <input
                        placeholder="Password"
                        type="password"
                        name="password"
                    />
                    <label class="text-left text-sm">
                        <input type="checkbox" bind:checked={isAgreed} />
                        Agree Term & Condition
                    </label>
                    <button
                        disabled={!isAgreed}
                        class="btn btn-primary rounded-lg"
                        type="submit">Daftar</button
                    >
                    <p class="text-sm">
                        Sudah akun?
                        <a class="text-purple-500" href="/login">login</a>
                    </p>
                </form>
            </div>
        </div>
    </div>
</main>
