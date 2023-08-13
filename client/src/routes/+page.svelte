<script>
    import MagnifyIcon from '$lib/assets/icons/MagnifyIcon.svelte'
    import Input from '$lib/components/Input.svelte'
    import Text from '$lib/components/Text.svelte'

    let inputRef = null

    $: if (inputRef && typeof inputRef.focus === 'function') {
        inputRef.focus()
    }
</script>

<main class="flex items-center justify-center">
    <div class="w-full mt-[30vh] flex justify-center">
        <Input bind:this={inputRef} placeholder="Search for a Player / Guild / Leaderboard..." rounded tw="w-4/5">
            <MagnifyIcon slot="left" tw="w-5 h-5 text-neutral-400" />
            <Text slot="right" b type="h6" tw="text-neutral-400">CTRL+K</Text>
        </Input>
    </div>
</main>
