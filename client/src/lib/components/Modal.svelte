<script lang="ts">
    import { createEventDispatcher } from 'svelte'
    import ResponsiveContainer from '$lib/components/ResponsiveContainer.svelte'
    import Close from '$lib/assets/icons/Close.svelte'
    import { twMerge } from 'tailwind-merge'

    export let preventClose = false
    export let closeX = true
    export let open = false
    export let tw = ''

    let lastOpenTime = 0
    let modalRef: undefined | HTMLElement
    let dispatch = createEventDispatcher()

    const close = (event: Event) => {
        if (preventClose) return
        open = false
        dispatch('close', event)
    }

    $: lastOpenTime = open ? Date.now() : lastOpenTime
    const handleClick = (event: MouseEvent) => {
        const target = event.target as HTMLElement
        if (Date.now() - lastOpenTime < 250) return
        if (!(modalRef && modalRef.contains(target))) close(event)
    }

    const handleKeyDown = (event: KeyboardEvent) => {
        if (event.key === 'Escape') close(event)
        else return
        event.preventDefault()
    }
</script>

<svelte:window on:keydown={handleKeyDown} on:click={handleClick} />

<div class="{open ? 'scale-100 opacity-100' : 'scale-0 opacity-0'} transition-all duration-200 ease-in-out z-40 top-0 left-0 absolute w-full h-screen full flex items-center justify-center bg-black/50"
     style="backdrop-filter: blur(4px)">
    <ResponsiveContainer size="2xl">
        <div bind:this={modalRef} class="{twMerge('bg-neutral-950 rounded-2xl relative overflow-hidden', tw)}">
            {#if closeX}
                <div on:click={close} aria-hidden="true" class="w-full grid absolute z-50">
                    <Close tw="h-8 w-8 place-self-end mr-4 mt-4 opacity-60 hover:opacity-80 transition duration-150 cursor-pointer" />
                </div>
            {/if}
            <slot />
        </div>
    </ResponsiveContainer>
</div>
