<script>
    import { goto } from "$app/navigation";
    import class_color from "$lib/assets/class_color.webp";
    import class_bg from "$lib/assets/class_bg.webp";
    import { post } from "$lib/utils/api";

    async function login(event) {
        event.preventDefault();

        let formData = new FormData(event.target);
        const res = await post("login", formData);

        if (res.ok) {
            goto("/beranda");
            return;
        }

        const text = await res.text();

        alert(text);
    }
</script>

    <main class="bg-[#E5CFFE] bg-[url('/src/lib/assets/class_bg.webp')] bg-cover bg-center bg-blend-multiply flex flex-col justify-center items-center size-full">
    <div class="w-full max-w-md md:max-w-none text-center p-6">
        <div
            class="bg-white p-6 shadow-lg rounded-xl md:flex max-w-4xl md:mx-auto gap-4"
        >
            <img
                src={class_color}
                alt="login"
                class="rounded-lg shadow-lg mb-4 md:mb-0 md:max-w-md"
            />
            <div class="flex flex-col gap-4 w-full">
                <h2 class="text-2xl font-bold">Hi, Student</h2>
                <p class="text-semibold">Welcome to PATH!</p>
                <form class="flex flex-col gap-6" onsubmit={login}>
                    <input
                        placeholder="Email"
                        type="email"
                        name="email"
                        required
                    />
                    <input
                        placeholder="Password"
                        type="password"
                        name="password"
                        required
                    />
                    <a
                        href="/lupa"
                        class="text-end text-purple-500 px-2 text-sm"
                        >lupa password?</a
                    >
                    <button class="btn btn-primary rounded-lg" type="submit"
                        >Login</button
                    >
                    <p class="text-sm">
                        Belum punya akun?
                        <a class="text-purple-500" href="/daftar">daftar</a>
                    </p>
                </form>
            </div>
        </div>
    </div>
</main>
