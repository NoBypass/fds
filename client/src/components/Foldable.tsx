import {Disclosure, Transition} from '@headlessui/react'
import { ChevronUpIcon } from '@heroicons/react/20/solid'

type Props = {
    title: string,
    content: JSX.Element
}

export function Foldable(props:Props) {
    return (
        <div className="w-full pt-4">
            <Disclosure>
                {({ open }:any) => (
                    <>
                        <Disclosure.Button className="flex w-full justify-between rounded-lg bg-slate-800 px-4 py-2 text-left text-sm font-medium text-gray-300 hover:bg-slate-700 duration-150 focus:outline-none focus-visible:ring focus-visible:ring-purple-500 focus-visible:ring-opacity-75">
                            <span>
                                {props.title}
                            </span>
                            <ChevronUpIcon className={`${open ? 'rotate-180 transform' : ''} h-5 w-5`} />
                        </Disclosure.Button>
                        <Transition enter="transition duration-100 ease-out" enterFrom="transform scale-95 opacity-0" enterTo="transform scale-100 opacity-100" leave="transition duration-75 ease-out" leaveFrom="transform scale-100 opacity-100" leaveTo="transform scale-95 opacity-0">
                            <Disclosure.Panel className="px-4 pt-4 pb-2 text-sm text-gray-400">
                                {props.content}
                            </Disclosure.Panel>
                        </Transition>
                    </>
                )}
            </Disclosure>
        </div>
    )
}
