<script lang="ts">
    import Table from '$lib/components/Table.svelte'
    import Text from '$lib/components/Text.svelte'
    import Dropdown from '$lib/components/Dropdown.svelte'
    import Tri from '$lib/assets/icons/Tri.svelte'
    import Page from './Page.svelte'
    import DropdownOptions from '$lib/components/DropdownOptions.svelte'
    import DoubleRight from '$lib/assets/icons/DoubleRight.svelte'
    import DoubleLeft from '$lib/assets/icons/DoubleLeft.svelte'

    let rand: any[] = []
    for (let i = 0; i < 100; i++) {
        rand.push({
            a: 'a'+ Math.random(),
            b: 'b'+ Math.random(),
            c: 'c'+ Math.random()
        })
    }

    const playerOptions = [{ value: 10 }, { value: 20 }, { value: 50 }]
    let playersPerPage = '20'
    let page = 0
    $: pages = Math.ceil(rand.length/playersPerPageNum)
    $: playersPerPageNum = parseInt(playersPerPage)

    const minMax = (num: number): number => {
        if (num > pages) return pages
        if (num < 0) return 0
        return num
    }

    const generateBetween = (start: number, end: number): number[] => {
        let arr = []
        for (let i = start; i < end; i++) {
            arr.push(i)
        }
        return arr
    }

    $: showing = generateBetween(minMax(page-2), minMax(page+3))
    $: if (showing.length < 5) {
        if (showing[0] == 0) {
            showing = generateBetween(0, 5)
        } else {
            showing = generateBetween(pages-5, pages)
        }
    }
</script>

<Text type="h1" tw="mb-10">Leaderboards</Text>

<div class="w-full flex justify-between">
    <Text color="darkened">Total: <Text type="s">{rand.length}</Text></Text>

    <Dropdown>
        <Text slot="title" color="darkened">Players per page: <Text type="s">{playersPerPage}</Text></Text>
        <DropdownOptions bind:value={playersPerPage} items={playerOptions} let:item>
            {item.value}
        </DropdownOptions>
    </Dropdown>
</div>

<Table tw="mt-2" columns={['a', 'b', 'c']} data={rand.slice(page*playersPerPageNum, (page+1)*playersPerPageNum)} />

<div class="w-full flex justify-center mt-8 gap-1">
    <ul class="flex border border-slate-600/40 rounded-lg">
        <Page on:click={() => page = 0}>
            <DoubleLeft tw="w-4 h-4" />
        </Page>
        <Page disabled={page == 0} on:click={() => page--}>
            <Tri tw="rotate-90 w-4 h-4" />
        </Page>

        {#each showing as i}
            <Page active={i == page} on:click={() => page = i}>{i+1}</Page>
        {/each}

        <Page disabled={page == pages-1} on:click={() => page++}>
            <Tri tw="-rotate-90 w-4 h-4" />
        </Page>
        <Page on:click={() => page = pages-1}>
            <DoubleRight tw="w-4 h-4" />
        </Page>
    </ul>
</div>
