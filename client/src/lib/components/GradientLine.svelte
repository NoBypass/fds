<script lang="ts">
    import { twMerge } from 'tailwind-merge'

    export let tw = ''
</script>

<div class={twMerge('relative z-50 h-px w-screen bg-gradient-to-r from-transparent via-white/30 to-transparent', tw)} />