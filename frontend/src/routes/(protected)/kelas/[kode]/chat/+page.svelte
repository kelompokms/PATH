<script>
    let resAI = $state();
    let isLoading = $state(false);

    async function handleForm(event) {
        event.preventDefault();
        const formData = new FormData(event.target);

        const data = {
            student_message: formData.get("message"),
        };

        isLoading = true;

        const res = await fetch(
            "http://100.67.36.47:8000/path-ai/tanya-tutor",
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data),
            },
        );
        isLoading = false;
        if (!res.ok) {
            alert("Gagal mendapatkan respon");
            return;
        }
        const json = await res.json();
        resAI = json;
    }
</script>

<div class="flex flex-col justify-center items-center grow">
    <div
        class="bg-white shadow-lg border-2 border-black/10 rounded-lg items-center w-full max-w-2xl"
    >
        <form
            onsubmit={handleForm}
            class="flex flex-col gap-8 p-4 py-8 text-center"
        >
            <h3 class="text-2xl font-bold">Ada yang bisa saya bantu?</h3>
            <input
                name="message"
                placeholder="tanya di sini..."
                class="w-full sm:w-sm mx-auto"
                type="text"
            />
        </form>
        {#if isLoading}
            <div class="w-full flex justify-center items-center p-4">
                <span class="loading loading-dots loading-lg"></span>
            </div>
        {:else}
            {#if resAI}
                <hr class="border-2 border-black/20" />
                <p class="p-8">
                    {resAI.response}
                </p>
            {/if}
        {/if}
    </div>
</div>
