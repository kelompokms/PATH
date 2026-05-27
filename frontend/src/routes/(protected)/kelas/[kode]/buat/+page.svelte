<script>
    import { goto } from "$app/navigation";
    import AngleLeft from "$lib/svg/angle-left.svelte";
    import { post } from "$lib/utils/api";
    import { onMount } from "svelte";

    let { params } = $props();

    let hasTenggat = $state(false);
    const today = new Date().toISOString().slice(0, 16);

    async function handleForm(event) {
        event.preventDefault();
        const formData = new FormData(event.target);
        if (formData.get("tenggat")) {
            const formDate = new Date(formData.get("tenggat")).toISOString();
            formData.set("tenggat", formDate);
        }
        const res = await post("class/" + params.kode + "/post", formData);

        if (!res.ok) {
            const text = await res.text();
            alert(text);
            return;
        }

        goto("forum");
    }
</script>

<div class="grow flex flex-col justify-center items-center">
    <form
        onsubmit={handleForm}
        class="flex flex-col border-2 border-black/10 p-4 gap-4 bg-white shadow-lg rounded-lg w-full max-w-md"
    >
        <button
            onclick={() => goto("forum")}
            class="btn btn-ghost pl-1 pr-2 text-lg self-start mb-2 *:first:size-6"
            ><AngleLeft />
            <span>Kembali</span>
        </button>
        <input
            type="text"
            name="nama"
            placeholder="Judul"
            maxlength="64"
            required
        />
        <textarea
            rows="6"
            type="text"
            name="deskripsi"
            placeholder="Deskripsi"
            required
        ></textarea>
        <select
            name="tipe"
            required
            onchange={(event) => (hasTenggat = event.target.value == "tugas")}
        >
            <option value="materi">Materi</option>
            <option value="kuis">Kuis</option>
            <option value="tugas">Tugas</option>
        </select>
        {#if hasTenggat}
            <input
                min={today}
                type="datetime-local"
                name="tenggat"
                placeholder="Tenggat"
                required
            />
        {/if}
        <button class="btn btn-primary" type="submit">Buat Post</button>
    </form>
</div>
