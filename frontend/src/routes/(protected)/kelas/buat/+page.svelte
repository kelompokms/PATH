<script>
    import { post } from "$lib/utils/api";

    function handleForm(event) {
        event.preventDefault();

        post("");
    }
</script>

<main class="flex justify-center items-center">
    <form class="w-full max-w-sm" onsubmit={handleForm}>
        <fieldset class="fieldset rounded-lg shadow-lg p-8 bg-white">
            <legend class="fieldset-legend">Buat Kelas</legend>
            <label for="nama" class="w-full label">Nama Kelas *</label>
            <input
                required
                id="nama"
                class="w-full input"
                type="text"
                name="nama"
            />
            <label for="bagian" class="w-full label">Bagian</label>
            <input id="bagian" class="w-full input" type="text" name="bagian" />
            <label for="deskripsi" class="w-full label">Deskripsi kelas</label>
            <input
                id="deskripsi"
                class="w-full input"
                type="text"
                name="deskripsi"
            />
            <button class="mt-4 btn btn-primary" type="submit"
                >Buat Kelas</button
            >
        </fieldset>
    </form>
</main>
