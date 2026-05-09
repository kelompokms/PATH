<script>
    let { icon, name } = $props();
</script>

<li>
    <button
        class="is-drawer-close:tooltip is-drawer-close:tooltip-right hover:bg-purple-200 active:bg-purple-400 text-black"
        data-tip={name}
    >
        {@render icon()}
        <span class="is-drawer-close:hidden text-lg">{name}</span>
    </button>
</li>
