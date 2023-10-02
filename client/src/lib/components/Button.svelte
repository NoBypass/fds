<script lang="ts">
    import { createEventDispatcher } from 'svelte'
    import { twMerge } from 'tailwind-merge'

    export let disabled = false
    export let rounded = false
    export let type: 'fancy' | 'primary' | 'normal' | 'transparent' = 'normal'
    export let href = ''
    export let tw = ''

    const styles = {
        fancy: 'bg-gradient-primary normal-b',
        transparent: 'bg-transparent text-white',
        normal: 'border border-gray-500/60 bg-gray-500/[.12] shadow-inset-neutral opacity-80 hover:opacity-100',
        primary: 'border border-purple-500/60 bg-purple-500/[.12] shadow-inset-primary opacity-80 hover:opacity-100'
    }

    let buttonRef: undefined | HTMLElement
    let dispatch = createEventDispatcher()
    $: styleClass = `${disabled ? styles[type].split(' ').filter(c => !c.startsWith('hover:')).join(' ') : styles[type]} ${disabled ? 'opacity-50 saturate-50 ' : 'active:scale-[.97]'}`

    const handleClick = (e: MouseEvent) => {
        dispatch('click', e)
        if (buttonRef) {
            if (href != '') window.location.replace(href)

            const ripple = document.createElement('span')
            const bounds = buttonRef.getBoundingClientRect()
            ripple.style.left = `${e.clientX - bounds.left}px`
            ripple.style.top = `${e.clientY - bounds.top}px`
            'absolute rounded-full bg-white/30 -translate-x-1/2 -translate-y-1/2 w-4 h-4 animate-[ripple_1s_ease-in-out_infinite]'
                .split(' ').forEach((c) => ripple.classList.add(c))
            buttonRef.classList.add('animate-resize')
            buttonRef.appendChild(ripple)
            setTimeout(() => {
                ripple.remove()
            }, 400)
        }
    }
</script>

<button bind:this={buttonRef}
        on:mousedown={handleClick}
        type="button"
        {disabled}
        class="{twMerge(`${styleClass} z-10 overflow-hidden ${rounded ? 'rounded-full' : 'rounded-md'} px-5 py-1.5 transition-all duration-150 relative`, tw)}">
    <slot />
</button>
