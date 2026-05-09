<script>
    import Wolf from "$lib/assets/wolf.webp";
    function login(event) {
        event.preventDefault();

        window.location.href = "/beranda";
    }
</script>

<main class="bg-purple-100 flex flex-col justify-center items-center size-full">
    <div class="w-full max-w-md md:max-w-none text-center p-6">
        <div
            class="bg-white p-6 shadow-lg rounded-xl md:flex max-w-4xl md:mx-auto gap-4"
        >
            <img
                src={Wolf}
                alt="login"
                class="rounded-lg shadow-lg mb-4 md:mb-0 md:max-w-md"
            />
            <div class="flex flex-col gap-4 w-full">
                <h2 class="text-2xl font-bold">Daftar</h2>
                <form class="flex flex-col gap-6" onsubmit={login}>
                    <input placeholder="Nama Lengkap" type="text" name="text" />
                    <input placeholder="Email" type="email" name="email" />
                    <input
                        placeholder="Password"
                        type="password"
                        name="password"
                    />
                    <label class="text-left text-sm">
                        <input type="checkbox" />
                        Agree Term & Condition
                    </label>
                    <button class="btn btn-primary rounded-lg" type="submit"
                        >Daftar</button
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
