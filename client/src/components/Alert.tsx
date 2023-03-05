import React from 'react'
import {CheckIcon, MinusCircleIcon, XMarkIcon} from "@heroicons/react/24/outline";

type Props = {
    title: string
    text: string
    type: 'error' | 'positive'
    width?: string
}

export const Alert = (props: Props) => {
    let color; let icon
    if (props.type == 'error') {
        color = 'red'
        icon = <XMarkIcon className={'h-6 text-red-700'} strokeWidth="3"/>
    } else if (props.type == 'positive') {
        color = 'amber'
        icon = <MinusCircleIcon />
    } else if (props.type == 'success') {
        color = 'green'
        icon = <CheckIcon />
    }

    return (
        <div className={'w-full flex justify-center my-4'}>
            <div className={`${props.width == null? 'w-full': props.width} 10/12 bg-${color}-700/[.3] border border-${color}-700/[.8] p-2 rounded-lg`}>
                <header className={'flex gap-2'}>
                    {icon}
                    <h2 className={'font-bold'}>{props.title}</h2>
                </header>
                <p>{props.text}</p>
            </div>
        </div>
    )
}

export default Alert