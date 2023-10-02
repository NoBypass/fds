<script lang="ts">
    import { createEventDispatcher } from 'svelte'

    export let placeholder = ''
    export let rounded = false
    export let disabled = false
    export let color: 'neutral' | 'error' | 'success' | 'warning' = 'neutral'
    export let password = false
    export let value = ''
    export let tw = ''
    export let regex = /./
    export let id = 'c'

    let inputRef: undefined | HTMLInputElement
    let mainRef: undefined | HTMLDivElement
    let isFocused = false

    const classnames = 'ring-1 ring-purple-500 outline outline-1 outline-offset-4 outline-purple-500/40'.split(' ')
    const dispatch = createEventDispatcher()
    const colors = {
        neutral: 'hover:border-neutral-500 border-neutral-700 hover:bg-neutral-700 hover:bg-opacity-10',
        error: 'border-rose-500 hover:border-rose-600',
        success: 'border-emerald-500 hover:border-emerald-600',
        warning: 'border-yellow-500 hover:border-yellow-600',
    }

    $: if (mainRef) {
        mainRef.addEventListener('click', () => {
            isFocused = true
        })
    }

    $: handleFocus(isFocused)
    const handleFocus = (f: boolean) => {
        if (!mainRef) return
        for (const c of classnames) {
            if (f) mainRef.classList.add(c)
            else mainRef.classList.remove(c)
        }
        mainRef.classList.toggle('hover:border-neutral-500')
    }

    const handleInput = () => {
        if (!inputRef) return
        const v = inputRef.value
        if (regex.test(v)) {
            dispatch('input', v)
            value = inputRef.value
        } else inputRef.value = value.slice(0, -1)

        if (!isFocused) isFocused = true
    }
</script>

<div class="pt-8">
    <div bind:this={mainRef} class="{colors[color]} space-between cursor-text transition duration-150 bg-transparent items-center gap-2 py-1 border flex {rounded ? 'rounded-full' : 'rounded-lg'} px-2.5 {tw}">
        <slot name="left" />
        <label for={id} class="z-0 absolute text-white/50 transition duration-150 {isFocused ? '-translate-y-9 -translate-x-2.5' : ''}">{placeholder}</label>
        <input on:blur={() => isFocused = false}
               on:input={handleInput}
               bind:this={inputRef}
               {value} {disabled} id={id}
               type="{password ? 'password' : 'text'}"
               class="z-10 placeholder:text-neutral-400 focus:outline-0 w-full bg-transparent">
        <slot name="right" />
    </div>
</div>
