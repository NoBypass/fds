<script lang="ts">
    import { twMerge } from 'tailwind-merge'

    export let tw = ''
    export let status: 'on' | 'off' | 'idle' = 'on'

    let self: undefined | string | null
    $: if (typeof localStorage !== 'undefined') {
        self = localStorage.getItem('self')
    }
</script>

<div class="w-12">
    {#if (status !== 'off')}
        <span class="absolute -translate-x-1/2 -translate-y-1/2 flex h-3.5 w-3.5">
            <span class="outline outline-2 outline-black relative inline-flex rounded-full h-3.5 w-3.5 {status === 'on' ? 'bg-emerald-500' : 'bg-amber-300'}"></span>
            {#if (status === 'on')}
                <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-emerald-400 opacity-75"></span>
            {/if}
        </span>
    {/if}
    <img src="https://minotar.net/avatar/{self}" alt="mc-head" class="{twMerge(`rounded-xl outline outline-2 outline-offset-2 ${status === 'on' ? 'outline-emerald-500' : status === 'idle' ? 'outline-amber-300' : 'outline-neutral-800' } h-8 w-8`, tw)}">
</div>
