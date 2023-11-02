<script lang="ts">
    import Tri from '$lib/assets/icons/Tri.svelte'
    import Button from '$lib/components/Button.svelte'
    import { twMerge } from 'tailwind-merge'

    export let tw = ''
    export let limit = 20
    export let offset = 0
    export let data: any[] = []
    export let columns: string[] = []
    let sorting: {[key: string]: boolean} = {}

    const getBorder = (i: number) => {
        if (i === 0) return 'rounded-l-lg'
        if (i === columns.length-1) return 'rounded-r-lg'
        return ''
    }

    const sort = (col: string) => {
        sorting[col] = !sorting[col]
        data = data.sort((a, b) => {
            if (sorting[col]) return b[col].localeCompare(a[col])
            return a[col].localeCompare(b[col])
        })
    }

    $: keys = Object.keys(data[0] || {})
    $: if (!columns || columns.length !== keys.length) columns = keys
</script>

<div class="overflow-x-auto">
    <table class={twMerge('border-collapse table-auto w-full', tw)}>
        <tr class="bg-slate-950/60">
            {#each columns as col, i}
                <th class="text-left {getBorder(i)}">
                    <Button noAnim on:click={() => sort(col)} type="transparent" tw="group hover:opacity-70 flex gap-2 py-1.5 px-3 w-full items-center cursor-pointer">
                        {col[0].toUpperCase() + col.slice(1)}
                        <div class="w-6 h-6 cursor-pointer grid place-content-center transition-opacity duration-150 group-hover:opacity-100 {sorting[col] === undefined ? 'opacity-0' : ''}">
                            <Tri tw="w-3.5 h-3.5 transition-transform duration-150 {sorting[col] ? 'rotate-180' : 'rotate-0'}" stroke={2.5} />
                        </div>
                    </Button>
                </th>
            {/each}
        </tr>
        {#each data.slice(offset, limit+offset) as row}
            <tr class="hover:bg-slate-600/20 cursor-pointer">
                {#each columns as col, i}
                    <td class="px-3 py-1.5 {getBorder(i)}">
                        <slot item={row[col]} col={col} />
                    </td>
                {/each}
            </tr>
        {/each}
    </table>
</div>
