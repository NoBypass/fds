<script lang="ts">
    import type { ComponentType } from 'svelte'
    import DoubleRight from '$lib/assets/icons/DoubleRight.svelte'
    import DoubleLeft from '$lib/assets/icons/DoubleLeft.svelte'
    import Button from '$lib/components/Button.svelte'

    export let children: ComponentType[] = []

    let containerRef: HTMLDivElement | undefined
    let offset = 0
    $: width = containerRef?.clientWidth || 0
    $: fullWidth = width * children.length

    const left = () => {
        offset += width
        if (offset > 0) {
            offset -= fullWidth
        }
    }

    const right = () => {
        offset -= width
        if (offset < -fullWidth + width) {
            offset += fullWidth
        }
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

<div class="flex gap-2">
    {#each children as child}
        <span id={children.indexOf(child).toString()} class="w-1 h-1 bg-white/60 rounded-full" />
    {/each}
</div>
