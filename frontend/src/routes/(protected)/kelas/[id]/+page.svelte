<script>
    let { params } = $props();
</script>

<main>
    <a href="/kelas" class="btn btn-outline">&lt; Kembali</a>
    <p class="font-bold text-2xl text-center">
        Kamu di kelas {params.id}
    </p>
</main>
