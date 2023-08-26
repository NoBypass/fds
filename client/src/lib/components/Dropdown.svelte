<script lang="ts">
    import { twMerge } from 'tailwind-merge'

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

<div>
    <div aria-hidden="true" class="cursor-pointer" on:click={show}>
        <slot name="trigger" />
    </div>
    <div class="{twMerge(`${open ? 'scale-100 opacity-100' : 'scale-0 opacity-0'} p-2 transition-all duration-100 ease-in-out absolute bg-neutral-900 rounded-lg min-w-[180px]`, tw)}">
        <slot name="content" />
    </div>
</div>
