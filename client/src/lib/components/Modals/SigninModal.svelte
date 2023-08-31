<script lang="ts">
    import loginBg from '$lib/assets/login-bg.png'
    import Button from '$lib/components/Button.svelte'
    import Input from '$lib/components/Input.svelte'
    import Checkbox from '$lib/components/Checkbox.svelte'
    import Modal from '$lib/components/Modal.svelte'
    import Text from '$lib/components/Text.svelte'
    import { createEventDispatcher } from 'svelte'
    import Spinner from '$lib/components/Spinner.svelte'
    import { makeGraphQLRequest } from '$lib/api/graphql.js'
    import type { Player } from '$lib/types/hypixel'

    export let open = false

    let isValid = false
    const errors = {
        password: '',
        username: ''
    }
    const pwdChecks = [
        {
            name: 'Password must be between 1 and 32 characters long',
            check: () => password.length > 0 && password.length <= 32
        },
        {
            name: 'Password must be at least 8 characters long',
            check: () => password.length >= 8
        },
        {
            name: 'Password must contain at least one uppercase letter',
            check: () => /[A-Z]/.test(password)
        },
        {
            name: 'Password must contain at least one lowercase letter',
            check: () => /[a-z]/.test(password)
        },
        {
            name: 'Password must contain at least one number',
            check: () => /[0-9]/.test(password)
        },
        {
            name: 'Password must contain at least one special character',
            check: () => /[!@#$%^&*()\-_=+{}[\]:;'<>,.?/\\|~`]/.test(password)
        },
        {
            name: 'Password must not contain any spaces',
            check: () => !password.includes(' ')
        }
    ]
    const usrChecks = [
        {
            name: 'Username must be between 1 and 16 characters long',
            check: () => username.length > 0 && username.length <= 16
        },
        {
            name: 'Username must only contain letters, numbers, and underscores',
            check: () => /^[a-zA-Z0-9_]+$/.test(username)
        }
    ]
    let prev = {
        at: 0,
        val: ''
    }

    let playerStatus: undefined | 'loading' | 'success'
    let username = ''
    let password = ''
    let remember = false

    $: {
        if (username !== '' && username != prev.val) {
            prev.at = Date.now()
            prev.val = username

            playerStatus = 'loading'
            const delay = 700

            setTimeout(() => {
                if (Date.now() - prev.at >= delay && errors.username === '') {
                    fetchPlayer()
                } else playerStatus = undefined
            }, delay)
        }
    }

    const fetchPlayer = async () => {
        await makeGraphQLRequest<Player>(`query {
            player(name: "${username}") {
                name
            }
        }`).then(res => {
            if (typeof res !== 'string' && res.data && res.data.name) {
                playerStatus = 'success'
            } else {
                playerStatus = undefined
                errors.username = 'Player not found'
            }
        })
    }

    $: {
        isValid = pwdChecks.every(check => check.check()) && usrChecks.every(check => check.check())
        const pwd = pwdChecks.filter(check => !check.check())[0]?.name
        errors.password = pwd ? pwd : ''
        const usr = usrChecks.filter(check => !check.check())[0]?.name
        errors.username = usr ? usr : ''
        if (password.length === 0) errors.password = ''
        if (username.length === 0) errors.username = ''
    }

    const dispatch = createEventDispatcher()
    const close = () => {
        open = false
        dispatch('close')
    }

    const submit = () => {
        dispatch('submit', {
            username,
            password,
            remember
        })
    }
    let userInputColor: 'neutral' | 'error' | 'success' | 'warning'
    $: userInputColor = errors.username.length !== 0 ? 'error' : username.length === 0 ? 'neutral' : playerStatus === 'loading' ? 'warning' : 'success'
</script>

<Modal open={open} on:close={close} tw="h-160 flex space-between">
    <img src={loginBg}
         alt="background"
         class="w-1/2 h-full object-none z-20 md:block hidden"
         style="-webkit-mask-image:-webkit-gradient(linear, left bottom, right bottom, from(rgba(0,0,0,1)), to(rgba(0,0,0,0)));
                mask-image: linear-gradient(to right, rgba(0,0,0,1), rgba(0,0,0,0));">
    <div class="w-full">
        <div class="md:p-6 lg:p-12 p-12 z-30 relative w-full h-full grid grid-cols-1 grid-rows-3">
            <div>
                <Text type="h1" size="xl" b>Login/Register</Text>
                <Text o tw="mt-4">You can both register a new account and log into an existing account via the same form.</Text>
            </div>

            <div>
                <Input color={userInputColor}
                       bind:value={username}
                       light rounded placeholder="Minecraft username" tw="mt-12 w-full">
                    <div slot="right">
                        {#if (playerStatus === 'success')}
                            <img src="https://minotar.net/avatar/{username}" alt="mc-head" class="h-6 w-6 rounded-md {userInputColor === 'success' ? '' : 'hidden'}">
                        {/if}
                        <Spinner color="warning" tw="{playerStatus !== 'loading' ? 'hidden' : ''}" />
                    </div>
                </Input>
                <Text tw="mt-2" color="error" b>{errors.username}</Text>

                <Input password color={errors.password.length !== 0 ? 'error' : password.length === 0 ? 'neutral' : 'success'}
                       bind:value={password} light rounded placeholder="Password" tw="mt-10 w-full" />
                <Text tw="mt-2" color="error" b>{errors.password}</Text>

                <Checkbox on:change={(e) => remember = e.detail.target.checked} tw="mt-10"><Text size="md">Remember me!</Text></Checkbox>
            </div>

            <div class="grid w-full grid-cols-2 grid-rows-1 mt-12 h-10 place-self-end">
                <Button on:click={close} rounded color="transparent"><Text b>Cancel</Text></Button>
                <Button on:click={submit} disabled={!isValid} rounded color="neutral"><Text b>Continue</Text></Button>
            </div>
        </div>
        <img src={loginBg} alt="background" class="saturate-150 absolute w-full h-full object-none blur-[100px]  z-10 left-0 top-0">
    </div>
</Modal>