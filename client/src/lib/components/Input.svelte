<script lang="ts">
    import { createEventDispatcher } from 'svelte'
    import { mouseStore } from '$lib/stores/location'

    export let placeholder = ''
    export let rounded = false
    export let disabled = false
    export let color: 'neutral' | 'danger' | 'success' | 'warning' = 'neutral'
    export let password = false
    export let value = ''
    export let tw = ''
    export let regex: RegExp | undefined = undefined
    export let id = 'c'

    let inputRef: undefined | HTMLInputElement
    let mainRef: undefined | HTMLDivElement
    let isFocused = false

    $: placeholderIsTop = isFocused || value.length > 0
    const classnames = 'ring-1 ring-purple-500 outline outline-1 outline-offset-4 outline-purple-500/40'.split(' ')
    const dispatch = createEventDispatcher()
    const colors = {
        neutral: 'hover:border-neutral-500 border-neutral-700 hover:bg-neutral-700 hover:bg-opacity-10',
        danger: 'border-rose-500 hover:border-rose-600',
        success: 'border-emerald-500 hover:border-emerald-600',
        warning: 'border-yellow-500 hover:border-yellow-600',
    }

    $: if (mainRef) {
        mouseStore.onClick(() => {
            isFocused = true
        }, mainRef.getBoundingClientRect())
    }

    $: handleFocus(isFocused)
    const handleFocus = (f: boolean) => {
        if (!mainRef) return
        for (const c of classnames) {
            if (f) {
                inputRef?.focus()
                mainRef.classList.add(c)
            } else mainRef.classList.remove(c)
        }
        mainRef.classList.toggle('hover:border-neutral-500')
    }

    const handleInput = () => {
        if (!inputRef) return
        const v = inputRef.value
        if (!regex) {
            dispatch('input', v)
            value = inputRef.value
        } else {
            if (regex && regex.test(v)) {
                dispatch('input', v)
                value = inputRef.value
            } else inputRef.value = value.slice(0, -1)
        }

        if (!isFocused) isFocused = true
    }
</script>

<div class="pt-8 {tw}">
    <div bind:this={mainRef}
         class="{colors[color]} space-between cursor-text transition duration-150 bg-transparent items-center gap-2 py-1 border flex {rounded ? 'rounded-full' : 'rounded-lg'} px-2.5">
        <slot name="left" />
        <label for={id} class="z-0 {placeholderIsTop ? 'absolute -translate-y-9 -translate-x-2.5' : 'whitespace-nowrap relative hover:cursor-text'} text-white/50 transition duration-150">{placeholder}</label>
        <input on:blur={() => isFocused = false}
               on:input={handleInput}
               on:focus={() => isFocused = true}
               bind:this={inputRef}
               {value} {disabled} id={id}
               type="{password ? 'password' : 'text'}"
               class="z-10 placeholder:text-neutral-400 focus:outline-0 w-full bg-transparent">
        <slot name="right" />
    </div>
    <div class="mt-1">
        <slot name="below" />
    </div>
</div>
