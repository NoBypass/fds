import * as React from 'react'

type Props = {
    type: string
    changeCallback: Function
    value: string
    showError: boolean
}

export function LoginInput(props: Props) {
    let name
    let inputType = 'text'
    if (props.type == 'cPassword') name = 'Confirm Password'
    if (props.type == 'cPassword' || props.type == 'password') inputType = 'password'

    return (
        <div className={'flex items-center justify-between my-2'}>
            <p>{name == null ? props.type.split('')[0].toUpperCase() + props.type.substring(1) : name}:</p>
            <div className={'flex flex-col items-center'}>
                <input name={props.type} className={`${props.showError? 'bg-red-600/[.1] hover:bg-red-600/[.2] border-red-600/[.8]': 'bg-blue-600/[.1] hover:bg-blue-600/[.2] border-blue-600/[.8]'} w-80 rounded-full border placeholder:italic placeholder:opacity-25 pl-3 focus:outline-offset-2 focus:outline-blue-600 focus:outline outline-2 hover:shadow-md duration-100 hover:shadow-blue-600/[.2] h-8`} onChange={c => props.changeCallback(c)} type={inputType} />
            </div>
        </div>
    )
}