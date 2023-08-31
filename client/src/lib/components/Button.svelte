<script lang="ts">
    import { createEventDispatcher } from 'svelte'
    import { twMerge } from 'tailwind-merge'

    export let rounded = false
    export let color: 'primary' | 'neutral' | 'transparent' = 'primary'
    export let disabled = false
    export let tw = ''

    const colors = {
        primary: 'bg-purple-500 hover:bg-purple-600 shadow-purple-500/50 text-white shadow-md',
        neutral: 'bg-white hover:bg-gray-200 shadow-gray-500/50 text-gray-800 shadow-md',
        transparent: 'bg-transparent text-white'
    }

    let buttonRef: undefined | HTMLElement
    let dispatch = createEventDispatcher()
    
    const handleClick = (e: MouseEvent) => {

        dispatch('click', e)
        if (buttonRef) {
            const ripple = document.createElement('span')
            ripple.style.left = `${e.offsetX}px`
            ripple.style.top = `${e.offsetY}px`
            'absolute rounded-full bg-white -translate-x-1/2 -translate-y-1/2 w-4 h-4 animate-[ripple_1s_ease-in-out_infinite]'
                .split(' ').forEach((c) => ripple.classList.add(c))
            buttonRef.appendChild(ripple)
            setTimeout(() => ripple.remove(), 400)
        }

    }
</script>

<style>
    @keyframes ripple {
        0% {
            scale: 0;
            opacity: .7;
        }
        100% {
            scale: 1;
            opacity: 0;
        }
    }
</style>

<button bind:this={buttonRef}
        on:click={handleClick}
        type="button"
        disabled={disabled}
        class="{twMerge(`${disabled ? colors[color].split(' ').filter((c) => !c.startsWith('hover:')).join(' ') : colors[color]} ${disabled ? 'opacity-70' : ''} z-10 overflow-hidden ${rounded ? 'rounded-full' : 'rounded-md'} px-5 py-1.5 transition duration-150 relative`, tw)}">
    <slot />
</button>
