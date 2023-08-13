<script lang="ts">
    import { onMount } from 'svelte'
    import { browser } from '$app/environment'
    import { twMerge } from 'tailwind-merge'

    export let tw = ''

    let open = false
    let openedAt = 0

    const show = () => {
        openedAt = new Date().getTime()
        open = true
    }

    const click = () => {
        if (open && new Date().getTime() - openedAt > 10) {
            open = false
        }
    }

    const escape = (event) => {
        if (event.key === 'Escape') {
            open = false
        }
    }

    onMount(() => {
        document.addEventListener('click', click)
        document.addEventListener('keydown', escape)
        return () => {
            if (browser) {
                document.removeEventListener('click', click)
                document.removeEventListener('keydown', escape)
            }
        }
    })
</script>

<div>
    <div aria-hidden="true" class="cursor-pointer" on:click={show}>
        <slot name="trigger" />
    </div>
    <div class="{twMerge(`${open ? 'scale-100 opacity-100' : 'scale-0 opacity-0'} p-2 transition-all duration-100 ease-in-out absolute bg-neutral-900 rounded-lg min-w-[180px]`, tw)}">
        <slot name="content" />
    </div>
</div>
