<script>
    const BANTUAN = [
        {
            judul: "Panduan AI",
            deskripsi:
                "Lorem ipsum dolor sit amet consectetur adipisicing elit. Tempore excepturi aperiam rem in vel itaque expedita distinctio minus voluptate, beatae nostrum dolorem provident eius, ducimus totam earum. Exercitationem, deserunt excepturi.",
        },
        {
            judul: "Etika Penggunaan AI",
            deskripsi:
                "Lorem ipsum dolor sit amet consectetur adipisicing elit. Tempore excepturi aperiam rem in vel itaque expedita distinctio minus voluptate, beatae nostrum dolorem provident eius, ducimus totam earum. Exercitationem, deserunt excepturi.",
        },
        {
            judul: "Who's Joe?",
            deskripsi: "Joe mama.",
        },
    ];
</script>

<main class="p-0 md:p-4">
    <div class="bg-white shadow-lg">
        <h2 class="bg-purple-200 p-4 text-2xl font-semibold">Bantuan</h2>
        <div class="join join-vertical w-full">
            {#each BANTUAN as item}
                <div
                    class="collapse collapse-arrow join-item border-base-300 border"
                >
                    <input
                        type="radio"
                        name="my-accordion-4"
                        checked="checked"
                    />
                    <div class="collapse-title font-semibold">
                        {item.judul}
                    </div>
                    <div class="collapse-content text-sm">
                        {item.deskripsi}
                    </div>
                </div>
            {/each}
        </div>
    </div>
</main>
