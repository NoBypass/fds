<script lang="ts">
    import { twMerge } from 'tailwind-merge'
    import { mouseStore } from '$lib/stores/location'
    import type { Point } from '$lib/types/common'

    export let tw = ''
    export let noAnim = false

    let mouseLocation: Point = { x: 0, y: 0 }
    let ref: HTMLSpanElement | undefined
    const rad = 200

    $: boundingBox = ref?.getBoundingClientRect()
    mouseStore.subscribe((point: Point) => {
        if (boundingBox && mouseStore.intersects(boundingBox)) {
            mouseLocation = { x: point.x - boundingBox.left - rad, y: point.y - boundingBox.top - rad }
        }
    })
</script>

<div class={twMerge(tw, 'group bg-white/[.01] border border-white/10 hover:border-white/[.15] bg-gradient-neutral p-4 rounded-3xl relative transition duration-300 overflow-hidden backdrop-blur-xl', noAnim ? '' : '')}>
    <span bind:this={ref} class="z-0 absolute group-hover:opacity-[.03] opacity-0 transition-opacity duration-300"
          style="background: radial-gradient(rgba(255,255,255) 0, transparent {rad}px);
              width: {rad*2}px;
              height: {rad*2}px;
              top: {mouseLocation.y}px;
              left: {mouseLocation.x}px;"
    />
    <div class="z-10 relative">
        <slot />
    </div>
</div>
