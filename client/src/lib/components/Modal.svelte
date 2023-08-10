<script>
    import { createEventDispatcher, onMount } from 'svelte'
    import { browser } from '$app/environment'
    import ResponsiveContainer from '$lib/components/ResponsiveContainer.svelte'

    export let open = false
    export let tw = ''

    let modalRef = null
    let dispatch = createEventDispatcher()

    const close = (event) => {
        open = false
        dispatch('close', event)
    }

    $: lastOpenTime = open ? Date.now() : lastOpenTime
    const handleOutsideClick = (event) => {
        if (Date.now() - lastOpenTime < 250) return
        if (!(modalRef && modalRef.contains(event.target))) close(event)
    }

    const handleKeyDown = (event) => {
        if (event.key === 'Escape') close(event)
        else return
        event.preventDefault()
    }

    onMount(() => {
        document.addEventListener('click', handleOutsideClick)
        document.addEventListener('keydown', handleKeyDown)
        return () => {
            if (browser) {
                document.removeEventListener('click', handleOutsideClick)
                document.removeEventListener('keydown', handleKeyDown)
            }
        }
    })
</script>


<div class="{open ? 'scale-100 opacity-100' : 'scale-0 opacity-0'} transition-all duration-200 ease-in-out z-50 top-0 left-0 absolute w-full h-screen full flex items-center justify-center bg-black/50"
     style="backdrop-filter: blur(4px)">
    <ResponsiveContainer size="2xl">
        <div bind:this={modalRef} class="{tw} bg-neutral-950 rounded-2xl relative overflow-hidden">
            <slot />
        </div>
    </ResponsiveContainer>
</div>
