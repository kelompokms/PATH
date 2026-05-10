<script>
    let { icon, name, href } = $props();
</script>

<li>
    <a
        {href}
        class="is-drawer-close:tooltip is-drawer-close:tooltip-right hover:bg-purple-200 active:bg-purple-400 text-black"
        data-tip={name}
    >
        {@render icon()}
        <span class="is-drawer-close:hidden text-lg">{name}</span>
    </a>
</li>
