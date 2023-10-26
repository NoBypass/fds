<script lang="ts">
    import { twMerge } from 'tailwind-merge'
    import Tri from '$lib/assets/icons/Tri.svelte'

    export let tw = ''

    let open = false
    let openedAt = 0

    const show = () => {
        openedAt = new Date().getTime()
        open = !open
    }

    const handleClick = () => {
        if (open && new Date().getTime() - openedAt > 1) {
            open = false
        }
    }

    const handleKeyDown = (event: KeyboardEvent) => {
        if (event.key === 'Escape') {
            open = false
        }
    }
</script>

<svelte:window on:keydown={handleKeyDown} on:click={handleClick} />

<button class="gap-1 cursor-pointer" on:click={show}>
    <div class="flex items-center gap-1">
        <slot name="title" />
        <Tri tw="w-4 h-4" />
    </div>
    <div class="{twMerge(`${open ? 'scale-100 opacity-100' : 'scale-0 opacity-0'} border border-slate-600/40 z-50 backdrop-blur-sm p-2 transition-all duration-100 ease-in-out absolute bg-slate-950/60 rounded-lg min-w-[180px]`, tw)}">
        <slot />
    </div>
</button>
