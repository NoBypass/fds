<script lang="ts">
    import { createEventDispatcher } from 'svelte'

    export let tw = ''
    export let checked = false
    export let id = 'c'

    const dispatch = createEventDispatcher()
</script>

<style>
    /* Animation by Dylan Raga (https://codepen.io/dylanraga/pen/Qwqbab) */

    input[type='checkbox'] + label > span{
        transition: all 250ms cubic-bezier(.4,.0,.23,1);
    }

    input[type='checkbox']:checked + label > span{
        animation: shrink-bounce 200ms cubic-bezier(.4,.0,.23,1);
    }

    input[type='checkbox']:checked + label > span:before{
        content: "";
        position: absolute;
        display: flex;
        justify-content: center;
        align-items: center;
        border-right: 2px solid transparent;
        border-bottom: 2px solid transparent;
        margin-left: -3px;
        transform: rotate(45deg);
        animation: checkbox-check 125ms 250ms cubic-bezier(.4,.0,.23,1) forwards;
    }

    @keyframes shrink-bounce{
        0%{
            transform: scale(1);
        }
        33%{
            transform: scale(.85);
        }
        100%{
            transform: scale(1);
        }
    }

    @keyframes checkbox-check{
        0%{
            width: 0;
            height: 0;
            border-color: #ffffff;
            transform: translate3d(0,0,0) rotate(45deg);
        }
        33%{
            width: 6px;
            height: 0;
            transform: translate3d(0,0,0) rotate(45deg);
        }
        100%{
            width: 6px;
            height: 12px;
            border-color: #ffffff;
            transform: translate3d(0,-.5em,0) rotate(45deg);
        }
    }
</style>

<div class="{tw} relative flex">
    <input bind:checked={checked} on:change={(e) => dispatch('change', e)} id={id} type="checkbox" class="h-0 w-0 hidden">
    <label for={id} class="flex items-center w-full gap-2 cursor-pointer">
        <span class="{checked ? 'border-[10px] border-purple-500 w-5' : 'border-2 border-white/30 w-6'} h-5 overflow-hidden rounded-md hover:bg-neutral-400/20" />
        <ins class="text-white hover:text-neutral-200 transition duration-150 w-full no-underline ease">
            <slot />
        </ins>
    </label>
</div>