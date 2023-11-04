<script lang="ts">
    import Card from '$lib/components/Card.svelte'
    import Text from '$lib/components/Text.svelte'
    import Input from '$lib/components/Input.svelte'
    import Checkbox from '$lib/components/Checkbox.svelte'
    import Button from '$lib/components/Button.svelte'
    import Spinner from '$lib/components/Spinner.svelte'

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

    $: name.valid = name.val.length > 0 && name.error === '' && name.resolved
    $: {
        cpassword.error = cpassword.val !== password.val && cpassword.val.length > 0 ? 'Passwords do not match.' : ''
        cpassword.valid = cpassword.val.length > 0 && cpassword.error === '' || !!isLoggingIn
    }
    $: {
        if (password.val.length === 0) {
            password.error = ''
        } else if (password.val.length < 10) {
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

        password.valid = password.val.length > 0 && password.error === ''
    }

    const auth = () => {
        // TODO: implement
    }
</script>

<main class="h-[60vh] grid">
    <Card tw="place-self-center sm:w-96 w-full">
        <form class="grid px-2 pb-2">
            <Text tw="mb-2" b type="h2">{isLoggingIn === undefined ? 'Login / Register' : isLoggingIn ? 'Login' : 'Register'}</Text>
            <Input bind:value={name.val}
                   bind:error={name.error}
                   bind:isSuccess={name.valid}
                   regex={nameRegex}
                   placeholder="Minecraft Name">
                <div slot="right">
                    {#if name.isLoading}
                        <Spinner />
                    {/if}
                </div>
            </Input>
            <Input bind:value={password.val}
                   bind:error={password.error}
                   bind:isSuccess={password.valid}
                   regex={pwdRegex}
                   password
                   placeholder={isLoggingIn === false ? 'Password (NOT YOUR MINECRAFT LOGIN)' : 'Password'} />

            {#if isLoggingIn === false}
                <Input bind:error={cpassword.error}
                       bind:value={cpassword.val}
                       bind:isSuccess={cpassword.valid}
                       regex={pwdRegex}
                       password
                       placeholder="Confirm Password" />
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