import React, {useEffect, useState, useRef} from "react"
import {Switch} from "@headlessui/react";
import {LoginButton} from "../small/LoginButton";
import {LoginInput} from "../small/LoginInput";
import fastFetch from "../../functions/fast-fetch";
import Alert from "../Alert";
import {createUser, getMojangPlayer, getUser} from "../../functions/api";
import {hash} from "../../functions/hash";

type Input = {
    [key: string]: string | undefined
    username?: string
    discord?: string
    cPassword?: string
    password?: string
}
type Setup = {
    name: string
    button: {
        text: string
        isDisabled: string[]
    },
    inputs: Input
}
type Props = {
    setTitle: (title: string) => void
}

export default function Login(props: Props) {
    const defaultLoginInfo = {
        data: '',
        isValid: false,
        showError: false
    }
    const [stayLogged, setStayLogged] = useState(false)
    const [loginInfo, setLoginInfo] = useState<{ [key: string]: any }>({
        username: defaultLoginInfo,
        password: defaultLoginInfo,
        cPassword: defaultLoginInfo,
        discord: defaultLoginInfo
    })
    const initialSetup = {
        name: 'ign',
        button: {
            text: 'Next',
            isDisabled: ['username']
        },
        inputs: {
            username: ''
        }
    }
    const [setup, setSetup] = useState<Setup>(initialSetup)
    const [inputArray, setInputArray] = useState<{ name: string, value: string, showError: boolean }[]>([])
    const setupRef = useRef<any>(initialSetup)
    const currentChangeRef = useRef<any>(null)
    const [uuid, setUuid] = useState('')

    async function validateInfo(name: string, value: string): Promise<boolean> {
        if (name === 'username') {
            if (value.length > 16) return false
            try {
                const data = await getMojangPlayer(value)
                setUuid(data.uuid)
                return data.success
            } catch (err) {
                return false
            }
        } else if (name === 'cPassword') {
            return value === loginInfo.password.data
        } else if (name === 'discord') {
            let split = value.split('#')
            return !(split.length !== 2 || split[1].length !== 4 || value.length > 32)
        } else if (name === 'password') {
            return value.length >= 8
        }
        return false
    }

    async function handleSubmit(e: any) {
        e.preventDefault()
        if (setup.name === 'ign') {
            const user = await getUser(uuid)
            if (!user.success) {
                props.setTitle(`Registration for ${loginInfo.username.data}`)
                setSetup({
                    name: 'register',
                    button: {
                        text: 'Register',
                        isDisabled: ['password', 'cPassword', 'discord']
                    },
                    inputs: {
                        password: '',
                        cPassword: '',
                        discord: ''
                    }
                })
            } else {
                props.setTitle(`Login for ${loginInfo.username.data}`)
                setSetup({
                    name: 'login',
                    button: {
                        text: 'Log in',
                        isDisabled: ['password']
                    },
                    inputs: {
                        password: ''
                    }
                })
            }
        } else {
            const hashedPwd = await hash(loginInfo.password.data)
            if (!hashedPwd.success) throw new Error('Error while generating password hash')
            const data = await createUser({
                password: hashedPwd.hash,
                discord: loginInfo.discord.data,
                stayLogged,
            }, uuid)
            if (data.success) localStorage.setItem('token', data.token)
        }
    }

    function handleChange(event: any) {
        const inputs = {
            ...setup.inputs,
            [event.target.name]: event.target.value,
        };
        setSetup({
            ...setup,
            inputs
        })
        currentChangeRef.current = { eventTarget: event.target, inputs, loginInfo, }
        setupRef.current = inputs
    }

    useEffect(() => {
        let arr = []
        for (const [key, value] of Object.entries(setup.inputs)) {
            arr.push({ name: key, value: value || '', showError: loginInfo[key].showError })
        }
        setInputArray(arr)
    }, [setup])

    useEffect(() => {
        if (currentChangeRef.current == null) return

        const changeToReference = currentChangeRef.current
        const prevSetup = setupRef.current

        setTimeout(async () => {
            const currentSetup = setupRef.current
            if (prevSetup === currentSetup) {
                const { name, value } = changeToReference.eventTarget
                const validation = await validateInfo(name, value)
                setLoginInfo({
                    ...loginInfo,
                    [name]: {
                        data: value,
                        isValid: validation,
                        showError: !validation && value !== '',
                    },
                })
            }
        }, 1000)
    }, [currentChangeRef.current])

    return (
        <>
            <form className={'flex flex-col mt-4'} onSubmit={e => handleSubmit(e)}>
                {inputArray.map((x) => <LoginInput key={x.name} type={x.name} value={x.value} changeCallback={handleChange} showError={loginInfo[x.name].showError} />)}
                <LoginButton text={setup.button.text} disabledWhen={!setup.button.isDisabled.every(disabledValue => loginInfo[disabledValue].isValid)}/>
            </form>
            <div className={'flex items-center justify-center w-full'}>
                <p className={'mr-4'}>Stay logged in?</p>
                <Switch checked={stayLogged} onChange={setStayLogged} className={`${stayLogged ? 'bg-green-700/[0.1] border-green-700/[.8]' : 'bg-red-700/[0.1] border-red-700/[.8]'} flex items-center border relative inline-flex h-7 w-12 shrink-0 cursor-pointer rounded-full transition-colors duration-200 ease-in-out focus:outline-none focus-visible:ring-2  focus-visible:ring-white focus-visible:ring-opacity-75`}>
                    <span aria-hidden="true" className={`${stayLogged ? 'translate-x-5' : 'translate-x-0'} pointer-events-none inline-block h-6 w-6 transform rounded-full bg-white shadow-lg ring-0 transition duration-200 ease-in-out`}/>
                </Switch>
            </div>
        </>
    )
}