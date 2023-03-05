import * as React from 'react';
import {useEffect, useState} from "react";

type Props = {
    text: string
    disabledWhen?: boolean
}

export function LoginButton({ text, disabledWhen }: Props) {
    return (
        <section className={'w-full flex justify-center mb-3'}>
            <button className={'my-4 bg-blue-700/[0.1] border-blue-700/[.8] border p-2 w-20 rounded-lg enabled:hover:bg-blue-700/[0.3] enabled:hover:border-blue-700 enabled:hover:shadow-md duration-100 enabled:hover:shadow-blue-600/[.2]'} type={'submit'} disabled={disabledWhen}>
                {text}
            </button>
        </section>
    )
}