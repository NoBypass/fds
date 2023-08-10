<script lang="ts">
    import { createEventDispatcher } from 'svelte'

    export let placeholder = ''
    export let rounded = false
    export let disabled = false
    export let light = false
    export let color: 'neutral' | 'error' | 'success' | 'warning' = 'neutral'
    export let password = false
    export let tw = ''

    let inputRef = null
    let mainRef = null

    const colors = {
        neutral: `hover:border-neutral-500 ${light ? 'border-neutral-300' : 'border-neutral-700'}`,
        error: 'border-rose-500 hover:border-rose-600',
        success: 'border-emerald-500 hover:border-emerald-600',
        warning: 'border-yellow-500 hover:border-yellow-600',
    }

    const dispatch = createEventDispatcher()
    const change = (e) => {
        dispatch('change', e)
    }

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

<div bind:this={mainRef} class="{colors[color]} space-between cursor-text transition duration-150 {light ? 'bg-white text-black' : 'bg-black'} items-center gap-2 py-1 border-2 flex {rounded ? 'rounded-full' : ''} px-4 {tw}">
    <slot name="left" />
    <input on:input={change} bind:this={inputRef} {disabled} type="{password ? 'password' : 'text'}" placeholder="{placeholder}" class="placeholder:text-neutral-400 focus:outline-0 w-full {light ? 'bg-white text-black' : 'bg-black'}">
    <slot name="right" />
</div>