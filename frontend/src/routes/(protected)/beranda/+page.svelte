<script>
    import { goto } from "$app/navigation";
    import Wolf from "$lib/assets/wolf.webp";
</script>

<main class="flex justify-center items-center">
    <div class="flex flex-col p-8 w-full gap-4 md:max-w-md">
        <div class="bg-white p-10 text-center rounded-md shadow-md">
            <!-- Gambar Di sini -->
            <img
                class="aspect-square rounded-full mb-4"
                src={Wolf}
                alt="gambar beranda"
            />
            <p class="text-md font-semibold">
                <span class="text-sky-800">Gabung Kelas</span>
                atau
                <span class="text-sky-800">Buat Kelas</span>
            </p>
        </div>
        <div class="flex gap-2">
            <button
                onclick={() => goto("/kelas/gabung")}
                class="grow p-2 rounded-md btn btn-primary">Gabung Kelas</button
            >
            <a
                href="/kelas/buat"
                class="grow p-2 rounded-md btn btn-primary btn-ghost text-blue-900 border-2 border-blue-100"
                >Buat Kelas</a
            >
        </div>
    </div>
</main>
