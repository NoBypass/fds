<script lang="ts">
    import { createEventDispatcher } from 'svelte'

    export let tw = ''

    export let checked = false

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
    <input bind:checked={checked} on:change={(e) => dispatch('change', e)} id="c" type="checkbox" class="h-0 w-0 hidden">
    <label for="c" class="flex items-center w-full gap-2">
        <span class="{checked ? 'border-[10px] border-purple-500' : 'border-2 border-neutral-400'} h-5 w-5 overflow-hidden rounded-md cursor-pointer hover:bg-neutral-400/20"></span>
        <ins class="text-white hover:text-neutral-200 transition duration-150 w-full no-underline ease">
            <slot />
        </ins>
    </label>
</div>