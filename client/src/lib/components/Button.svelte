<script lang="ts">
    let buttonRef = null
    let maxW = 0

    const handleClick = (e) => {
        if (buttonRef) {
            maxW = buttonRef.clientWidth

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
        class="overflow-hidden bg-purple-500 text-white rounded-md shadow-md shadow-purple-500/50 hover:bg-purple-600 transition duration-150 relative">
    <div class="w-full h-full px-5 py-1.5">
        <slot />
    </div>
</button>
