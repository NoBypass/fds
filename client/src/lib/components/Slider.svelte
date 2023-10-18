<script lang="ts">
    import type { ComponentType } from 'svelte'

    export let children: ComponentType[] = []

    let fullRef: HTMLDivElement | undefined
    let singleRef: HTMLDivElement | undefined
    let offset = 0
    $: width = singleRef?.offsetWidth || 0
    $: fullWidth = fullRef?.offsetWidth || 0

    const left = () => {
        offset += width
        if (offset > fullWidth - (2 * width)) {
            offset -= fullWidth
        }
    }

    const right = () => {
        offset -= width
        if (offset < -fullWidth + (2 * width)) {
            offset += fullWidth
        }
    }
</script>

<div class="flex">
    <button class="z-20" on:click={left}>Left</button>

    <div bind:this={fullRef} class="z-0 flex transition-all duration-200" style="transform: translateX({offset}px)">
        {#each children as child}
            <div bind:this={singleRef} class="border border-red-500">
                <svelte:component this={child}/>
            </div>
        {/each}
    </div>

    <button class="z-20" on:click={right}>Right</button>
</div>
<div class="flex gap-2">
    {#each children as child}
        <span id={children.indexOf(child).toString()} class="w-1.5 h-1.5 bg-white/60 rounded-full" />
    {/each}
</div>
