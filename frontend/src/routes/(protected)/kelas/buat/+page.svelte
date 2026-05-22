<script>
    import { goto } from "$app/navigation";
    import { post } from "$lib/utils/api";

    async function handleForm(event) {
        event.preventDefault();

        let formData = new FormData(event.target);
        console.log(formData);

        const res = await post("class", formData);
        const text = await res.text();

        if (!res.ok) {
            alert(text);
            return;
        }

        goto("/kelas/" + text);
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
                name="nama_kelas"
            />
            <label for="bagian" class="w-full label">Bagian</label>
            <input id="bagian" class="w-full input" type="text" name="bagian" />
            <button class="mt-4 btn btn-primary" type="submit"
                >Buat Kelas</button
            >
        </fieldset>
    </form>
</main>
