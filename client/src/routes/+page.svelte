<script>
    import MagnifyIcon from "$lib/icons/MagnifyIcon.svelte"

    let inputRef

    $: inputRef?.focus()
</script>

<main class="flex items-center justify-center">
    <div class="w-full mt-[30vh] grid grid-rows-3 grid-cols-8 items-center">
        <div class="flex z-50 bg-white items-center gap-3 border border-neutral-300 rounded-full px-4 py-1 col-span-6 col-start-2 row-start-2">
            <MagnifyIcon tw="w-5 h-5 text-neutral-400" />
            <input bind:this={inputRef} type="text" placeholder="Search for a Player / Guild / Leaderboard..." class="w-full focus:outline-none">
        </div>

        <div class="h-px blur-[2px] w-full col-start-2 row-start-1 place-self-end col-span-3"
             style="background: rgba(255,255,255,0);
             background: linear-gradient(90deg, rgba(255,255,255,0) 0%, rgba(0,0,255,1) 50%, rgba(255,255,255,0) 100%);"></div>
        <div class="h-px blur-[2px] w-full col-start-2 row-start-1 place-self-end col-span-3"
             style="background: rgba(255,255,255,0);
             background: linear-gradient(90deg, rgba(255,255,255,0) 0%, rgba(0,0,255,1) 50%, rgba(255,255,255,0) 100%);"></div>
        <div class="h-px blur-[1px] w-full col-start-4 row-start-1 place-self-end col-span-2"
             style="background: rgba(255,255,255,0);
             background: linear-gradient(90deg, rgba(255,255,255,0) 0%, rgb(255,0,0) 50%, rgba(255,255,255,0) 100%);"></div>
        <div class="h-px blur-[1px] w-full col-start-6 row-start-1 place-self-end col-span-1 translate-x-1/2"
             style="background: rgba(255,255,255,0);
             background: linear-gradient(90deg, rgba(255,255,255,0) 0%, rgba(0,0,255,1) 50%, rgba(255,255,255,0) 100%);"></div>

        <div class="h-px blur-[2px] w-full col-start-3 row-start-3 place-self-start col-span-3"
             style="background: rgba(255,255,255,0);
             background: linear-gradient(90deg, rgba(255,255,255,0) 0%, rgba(255,0,0,1) 50%, rgba(255,255,255,0) 100%);"></div>
        <div class="h-px blur-[2px] w-full col-start-3 row-start-3 place-self-start col-span-3"
             style="background: rgba(255,255,255,0);
             background: linear-gradient(90deg, rgba(255,255,255,0) 0%, rgba(255,0,0,1) 50%, rgba(255,255,255,0) 100%);"></div>
        <div class="h-px blur-[1px] w-full col-start-5 row-start-3 place-self-start col-span-2"
             style="background: rgba(255,255,255,0);
             background: linear-gradient(90deg, rgba(255,255,255,0) 0%, rgb(0,0,255) 50%, rgba(255,255,255,0) 100%);"></div>
    </div>
</main>
