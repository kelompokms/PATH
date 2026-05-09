<main class="bg-purple-100 flex flex-col justify-center items-center size-full">
    <div class="w-full">
        <div
            class="bg-white m-6 p-4 shadow-lg rounded-lg flex flex-col gap-4 text-center"
        >
            <h2 class="text-2xl font-bold">Daftar</h2>
            <form class="flex flex-col gap-6">
                <input placeholder="Nama panjang" type="text" name="Nama" />
                <input placeholder="Email" type="email" name="email" />
                <input placeholder="Password" type="password" name="password" />
                <div class="text-left">
                    <input type="checkbox" id="term" />
                    <label class="text-sm ml-1" for="term"
                        >Accept term & Condition</label
                    >
                </div>
                <button class="btn btn-blue" type="submit">Daftar</button>
            </form>
            <p class="text-sm">
                Sudah punya akun?
                <a class="text-purple-500" href="/login">login</a>
            </p>
        </div>
    </div>
</main>
