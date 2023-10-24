<script lang="ts">
    import type { ComponentType } from 'svelte'
    import DoubleRight from '$lib/assets/icons/DoubleRight.svelte'
    import DoubleLeft from '$lib/assets/icons/DoubleLeft.svelte'
    import Button from '$lib/components/Button.svelte'

    export let children: ComponentType[] = []

    let containerRef: HTMLDivElement | undefined
    let active = 0
    $: width = containerRef?.clientWidth || 0
    $: fullWidth = width * children.length
    $: offset = -width * active

    const left = () => {
        if (active === 0) active = children.length - 1
        else active--
    }

    const right = () => {
        if (active === children.length - 1) active = 0
        else active++
    }
</script>

<div class="grid w-full mb-4 items-center relative" bind:this={containerRef}>
    <Button noRipple
            type="transparent"
            tw="z-20 absolute h-full w-12 flex items-center justify-center"
            on:click={left}>
        <DoubleLeft />
    </Button>

    <div class="overflow-x-hidden" style="width: {width}px">
        <div class="z-0 flex transition-all duration-300" style="transform: translateX({offset}px)">
            {#each children as child}
                <div>
                    <svelte:component this={child} width={containerRef?.clientWidth || 0}/>
                </div>
            {/each}
        </div>
    </div>

    <Button noRipple
            type="transparent"
            tw="z-20 absolute h-full w-12 flex items-center justify-center place-self-end"
            on:click={right}>
        <DoubleRight />
    </Button>
</div>

<div class="flex h-5">
    {#each children as child, i}
        <div on:click={() => active = i} id={i} class="p-1.5 cursor-pointer grid place-content-center">
            <span class="rounded-full block transition-all duration-150 {i === active ? 'w-2 h-2 bg-white' : 'w-1.5 h-1.5 bg-white/60'}" />
        </div>
    {/each}
</div>
