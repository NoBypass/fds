<script>
    import Modal from '$lib/components/Modal.svelte'
    import { createEventDispatcher } from 'svelte'
    import Wave from '$lib/assets/shapes/Wave.svelte'
    import CircledCheckmark from '$lib/assets/icons/CircledCheckmark.svelte'
    import { tweened } from 'svelte/motion'
    import { cubicInOut } from 'svelte/easing'
    import Text from '$lib/components/Text.svelte'

    export let open = false

    let isSmall = false
    let progress = 0
    const dispatch = createEventDispatcher()
    const close = () => {
        dispatch('close')
    }

    const pathDrawProgress = tweened(0, {
        duration: 2000,
        easing: cubicInOut
    })

    $: if (open) pathDrawProgress.set(1)
    $: pathDrawProgress.subscribe((v) => progress = v)
    $: if (progress === 1) isSmall = true
</script>

<Modal open={open} on:close={close} tw="bg-emerald-500 w-128 h-64 grid">
    <div class="text-center grid">
        <CircledCheckmark progress={progress} tw="{isSmall ? 'w-20 h-20' : 'w-28 h-28'} transition-all duration-150 place-self-center" />
        {#if (isSmall)}
            <Text tw="text-white text-2xl font-bold">Success!</Text>
        {/if}
    </div>
    <Wave tw="absolute w-full text-emerald-600 place-self-end" />
</Modal>