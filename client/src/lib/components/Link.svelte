<script lang="ts">
    export let href = ''
</script>

<a class="shadow-inset-bottom hover:text-white text-white/70 hover:shadow-inset-bottom-2 transition-all duration-150 pb-px" href={href}>
    <slot />
</a>