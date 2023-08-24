<script lang="ts">
    import ResponsiveContainer from '$lib/components/ResponsiveContainer.svelte'
    import { browser } from '$app/environment'
    import MagnifyIcon from '$lib/assets/icons/MagnifyIcon.svelte'
    import { createEventDispatcher, onMount } from 'svelte'

    export let open = false
    let dispatch = createEventDispatcher()

    let lastOpenTime = 0
    let paletteRef: HTMLElement | undefined
    let inputRef: HTMLElement | undefined

    $: lastOpenTime = open ? Date.now() : lastOpenTime

    const handleClick = (event: MouseEvent) => {
        const target = event.target as HTMLElement
        if (Date.now() - lastOpenTime < 250) return
        if (!(paletteRef && paletteRef.contains(target))) close(event)
    }

    const handleKeyDown = (event: KeyboardEvent) => {
        if (event.key === 'Escape') close(event)
        else if (event.ctrlKey && event.key === 'k') open = true
        else return
        event.preventDefault()
    }

    const close = (event: Event) => {
        open = false
        dispatch('close', event)
    }

    let slotAnimation = 'max-h-0 opacity-0'
    $: {
        if (open && inputRef) {
            inputRef.focus()
            setTimeout(() => slotAnimation = 'max-h-96 opacity-1', 0)
        }
        else if (!open) slotAnimation = 'max-h-0 opacity-0'
    }
</script>

<svelte:window on:keydown={handleKeyDown} on:click={handleClick} />

{#if (open)}
    <div class="absolute w-full h-full flex items-center justify-center">
        <ResponsiveContainer tw="w-full">
            <div bind:this={paletteRef} class="px-16">
                <div class="border border-neutral-700 w-full h-auto p-3 rounded-xl shadow-md">
                    <div class="flex items-center bg-neutral-100 rounded-md pl-3">
                        <MagnifyIcon tw="h-5 w-5 text-neutral-400" />
                        <input bind:this={inputRef} type="text" placeholder="Search player or guild..." class="w-full px-3 bg-black py-1.5 focus:outline-none rounded-md">
                    </div>
                    <div class="{slotAnimation} transition-all duration-200 ease-in-out mt-2">
                        <slot />
                    </div>
                </div>
            </div>
        </ResponsiveContainer>
    </div>
{/if}