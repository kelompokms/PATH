<script>
    import Wolf from "$lib/assets/wolf.webp";
</script>

<main class="flex justify-center items-center">
    <div class="flex flex-col p-6 w-full gap-4 md:max-w-md">
        <div class="bg-white p-6 text-center rounded-lg shadow-md">
            <!-- Gambar Di sini -->
            <img
                class="aspect-square rounded-full mb-4"
                src={Wolf}
                alt="gambar beranda"
            />
            <p>Gabung Kelas atau Buat Kelas</p>
        </div>
        <div class="flex gap-2">
            <button class="grow p-2 rounded-xl btn btn-primary"
                >Gabung Kelas</button
            >
            <button
                class="grow p-2 rounded-xl btn btn-primary btn-ghost text-blue-900 border-2 border-blue-100"
                >Buat Kelas</button
            >
        </div>
    </div>
</main>
