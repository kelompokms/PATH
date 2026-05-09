<main class="bg-purple-100 flex flex-col justify-center items-center size-full">
    <div class="w-full">
        <div
            class="bg-white m-6 p-4 shadow-lg rounded-lg flex flex-col gap-4 text-center"
        >
            <h2 class="text-2xl font-bold">Hi, Student</h2>
            <p class="text-semibold">Welcome to PATH!</p>
            <form class="flex flex-col gap-6">
                <input placeholder="Email" type="email" name="email" />
                <input placeholder="Password" type="password" name="password" />
                <a href="/lupa" class="text-end text-purple-500 px-2 text-sm"
                    >lupa password?</a
                >
                <button class="btn btn-blue" type="submit">Login</button>
                <p class="text-sm">
                    Belum punya akun?
                    <a class="text-purple-500" href="/daftar">daftar</a>
                </p>
            </form>
        </div>
    </div>
</main>
