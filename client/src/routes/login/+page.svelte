<script lang="ts">
    import Card from '$lib/components/Card.svelte'
    import Text from '$lib/components/Text.svelte'
    import Input from '$lib/components/Input.svelte'
    import Checkbox from '$lib/components/Checkbox.svelte'
    import Button from '$lib/components/Button.svelte'
    import { alertStore } from '$lib/stores/alertStore'
    import Spinner from '$lib/components/Spinner.svelte'
    import api from '$lib/api/graphql'

    const initialState = {
        name: {
            val: '',
            error: '',
            resolved: false,
            isLoading: false,
            valid: false,
            lastChange: 0,
            prevName: '',
        },
        password: {
            val: '',
            error: '',
            valid: false,
        },
        cpassword: {
            val: '',
            error: '',
            valid: false,
        },
        remember: false,
    }

    let { name, password, cpassword, remember } = { ...initialState }
    const nameRegex = /^(?![0-9])[a-zA-Z0-9_]{0,16}$/
    const pwdRegex = /^.{0,64}$/
    let isLoggingIn: boolean | undefined

    $: {
        name.valid = name.val.length > 0 && name.error === '' && name.resolved
        password.valid = password.val.length > 0 && password.error === ''
        cpassword.error = cpassword.val !== password.val ? 'Passwords do not match.' : ''
        cpassword.valid = cpassword.val.length > 0 && cpassword.error === '' || !!isLoggingIn
    }

    const loginConfig = () => {
        name.error = ''
        name.isLoading = false
        name.resolved = true
        isLoggingIn = true
    }

    $: {
        if (name.val !== name.prevName) {
            name.lastChange = Date.now()
            name.prevName = name.val

            setTimeout(() => {
                if (name.lastChange + 500 < Date.now()) {
                    name.isLoading = true
                    api.account.name(name.val).then(({ data }) => {
                        name.val = data.name
                        loginConfig()
                    }).catch(async (err) => {
                        if (err.status === 404) {
                            return await accountNotTaken()
                        }

                        name.error = err.message
                        name.isLoading = false
                        name.resolved = true
                        alertStore.push(err.message, 'danger')
                    })
                }
            }, 500)
        }
    }

    const accountNotTaken = async () => {
        isLoggingIn = false
        api.player.name(name.val).then(({ data }) => {
            name.val = data.name
            console.log(data)
            name.error = ''
        }).catch(() => {
            name.error = 'This Minecraft account does not exist.'
        })

        name.isLoading = false
        name.resolved = true
    }

    const auth = () => {
        if (password.val.length < 10) {
            password.error = 'Password must be at least 10 characters long.'
        } else if (!/[A-Z]/.test(password.val)) {
            password.error = 'Password must contain at least one uppercase letter.'
        } else if (!/[a-z]/.test(password.val)) {
            password.error = 'Password must contain at least one lowercase letter.'
        } else if (!/[0-9]/.test(password.val)) {
            password.error = 'Password must contain at least one number.'
        } else if (! /[!@#$%^&*()\-_=+{}[\]:;"'<>,.?/\\|~`]/.test(password.val)) {
            password.error = 'Password must contain at least one special character.'
        } else {
            password.error = ''
        }

        if (!name.valid || !password.valid) {
            alertStore.push('Unexpected error, please fill out all the fields correctly.', 'danger')
            return
        }

        api.auth({
            name: name.val,
            password: password.val,
            remember
        }).then(({ data }) => {
            name.val = data.account.name
            localStorage.setItem('token', data.token)
            loginConfig()
            return
        }).catch((e) => {
            if (e.status === 404) {
                alertStore.push('This account does not exist.', 'danger')
            } else if (e.status === 401) {
                alertStore.push('Incorrect password.', 'danger')
            } else {
                alertStore.push(e.message, 'danger')
            }
        })
    }
</script>

<main class="h-[55vh] grid justify-center">
    <Card tw="place-self-end w-96">
        <form class="grid px-2 pb-2">
            <Text tw="mb-2" b type="h2">{isLoggingIn === undefined ? 'Login / Register' : isLoggingIn ? 'Login' : 'Register'}</Text>
            <Input id="ign" bind:value={name.val} color={name.error.length > 0 ? 'danger' : name.valid ? 'success' : 'neutral'} regex={nameRegex} placeholder="Minecraft Name">
                <div slot="right">
                    {#if name.isLoading}
                        <Spinner />
                    {/if}
                </div>
                <div slot="below">
                    {#if name.error.length > 0}
                        <Text size="sm" color="danger" b>{name.error}</Text>
                    {/if}
                </div>
            </Input>
            <Input id="pwd" bind:value={password.val} regex={pwdRegex} password placeholder="Password" />

            {#if isLoggingIn === false}
                <Input id="cpwd" bind:value={cpassword.val} regex={pwdRegex} password placeholder="Confirm Password" />
            {/if}

            <Checkbox tw="mb-8 mt-4" id="remember" bind:checked={remember}>
                <Text size="md">Remember me!</Text>
            </Checkbox>

            <div class="grid grid-cols-3">
                <Button href="/">Back</Button>
                <Button disabled={!name.valid || !password.valid || !cpassword.valid} on:click={auth} tw="col-start-3" type="primary">Continue</Button>
            </div>
        </form>
    </Card>
</main>