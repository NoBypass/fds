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
        normal: 'border border-gray-500/60 bg-gradient-glass opacity-80 hover:opacity-100',
        primary: 'bg-primary-glass border border-purple-500/60 bg-gradient-primary-glass'
    }

    let buttonRef: undefined | HTMLElement
    let dispatch = createEventDispatcher()
    const styleClass = `${disabled ? styles[type].split(' ').filter(c => c.startsWith('hover:')).join(' ') : styles[type]} ${disabled ? 'opacity-70' : ''}`
    
    const handleClick = (e: MouseEvent) => {
        dispatch('click', e)
        if (buttonRef) {
            if (href != '') window.open(href)

            const ripple = document.createElement('span')
            ripple.style.left = `${e.offsetX}px`
            ripple.style.top = `${e.offsetY}px`
            'absolute rounded-full bg-white/30 -translate-x-1/2 -translate-y-1/2 w-4 h-4 animate-[ripple_1s_ease-in-out_infinite]'
                .split(' ').forEach((c) => ripple.classList.add(c))
            buttonRef.classList.add('animate-resize')
            buttonRef.appendChild(ripple)
            setTimeout(() => {
                if (buttonRef) buttonRef.classList.remove('animate-resize')
                ripple.remove()
            }, 400)
        }
    }
</script>

<button bind:this={buttonRef}
        on:click={handleClick}
        type="button"
        disabled={disabled}
        class="{twMerge(`${styleClass} z-10 overflow-hidden ${rounded ? 'rounded-full' : 'rounded-md'} px-5 py-1.5 transition-all duration-150 relative`, tw)}">
    <slot />
</button>
