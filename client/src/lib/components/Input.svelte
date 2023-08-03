<script lang="ts">
    export let placeholder = ''
    export let rounded = false
    export let disabled = false
    export let light = false
    export let tw = ''

    let inputRef = null
    let mainRef = null

    $: if (inputRef && !light) {
        const classnames = 'shadow-md shadow-purple-500/50 border-purple-500'.split(' ')
        inputRef.addEventListener('focus', () => {
            for (const c of classnames) {
                mainRef.classList.add(c)
            }
            mainRef.classList.remove('hover:border-neutral-500')
        })
        inputRef.addEventListener('blur', () => {
            for (const c of classnames) {
                mainRef.classList.remove(c)
            }
            mainRef.classList.add('hover:border-neutral-500')
        })
    }

    $: if (mainRef) {
        mainRef.addEventListener('click', () => {
            if (inputRef) inputRef.focus()
        })
    }
</script>

<div bind:this={mainRef} class="space-between hover:border-neutral-500 cursor-text transition duration-150 {light ? 'bg-white text-black border-neutral-300' : 'bg-black border-neutral-700'} items-center gap-2 py-1 border-2 flex {rounded ? 'rounded-full' : ''} px-4 {tw}">
    <slot name="leftIcon" />
    <input bind:this={inputRef} {disabled} type="text" placeholder="{placeholder}" class="placeholder:text-neutral-400 focus:outline-0 w-full {light ? 'bg-white text-black' : 'bg-black'}">
    <slot name="rightIcon" />
</div>