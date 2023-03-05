import * as React from 'react'
import {Listbox, Transition} from "@headlessui/react";
import {ChevronUpDownIcon} from "@heroicons/react/24/outline";
import {Fragment} from "react";

type Selected = {
    name: string
    icon: JSX.Element
    placeholder: string
}

type Props = {
    selected: Selected
    setSelected: Function
    options: Selected[]
}

export default function SelectionBox(props: Props) {
    const { selected, setSelected, options } = props

    return (
        <Listbox value={selected} onChange={(v) => setSelected(v)}>
            <Listbox.Button className="px-2 rounded-r-full h-8 -skew-x-20 flex items-center justify-center hover:shadow-md duration-100 hover:shadow-blue-600/[.2] hover:bg-blue-600/[.2] cursor-default shadow-md sm:text-sm content-between">
                <span className="skew-x-20 h-5">{selected.name}</span>
                <span className="skew-x-20 pointer-events-none">
                    <ChevronUpDownIcon className="h-5 w-5 text-gray-400" aria-hidden="true" />
                </span>
            </Listbox.Button>
            <Transition as={Fragment} leave="transition ease-in duration-100" leaveFrom="opacity-100" leaveTo="opacity-0">
                <Listbox.Options className="z-50 absolute mt-1 max-h-60 rounded-md bg-blue-600/[.2] backdrop-blur-sm border border-blue-600/[.8] py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm cursor-pointer">
                    {options.map((type) => (
                        <Listbox.Option key={options.indexOf(type)} value={type} className={({active}) => `text-white relative cursor-default select-none py-1 px-2 ${active ? 'bg-blue-700' : 'text-gray-900'}`}>
                            {type.name}
                        </Listbox.Option>
                    ))}
                </Listbox.Options>
            </Transition>
        </Listbox>
    )
}