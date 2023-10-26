<script lang="ts">
    import { createEventDispatcher } from 'svelte'

    type Item = {
        value: string
    } & any
    export let items: Item[] = []
    export let value = ''

    const dispatch = createEventDispatcher()
    const click = (item: Item, e: any) => {
        const event = { item, e }
        dispatch('click', event)
        value = item.value
    }
</script>

<ul>
    {#each items as item}
        <li class="hover:bg-slate-600/30 text-white rounded-lg py-1">
            <button class="w-full h-full border-none bg-transparent" on:click={(e) => click(item, e)}>
                <slot item={item} />
            </button>
        </li>
    {/each}
</ul>
