import { Dialog, Transition  } from '@headlessui/react'
import { XMarkIcon } from '@heroicons/react/24/outline'
import { Fragment } from "react";

type Props = {
    content: JSX.Element
    open: boolean
    callback: (open: boolean) => void
    title?: string
}

export function Modal({ content, open, callback, title }: Props) {
    return (
        <Transition appear show={open} as={Fragment}>
            <Dialog as="div" className="relative z-10" onClose={() => callback(false)}>
                <Transition.Child as={Fragment} enter="ease-out duration-300" enterFrom="opacity-0" enterTo="opacity-100" leave="ease-in duration-200" leaveFrom="opacity-100" leaveTo="opacity-0">
                    <div className="fixed inset-0 bg-black bg-opacity-25" />
                </Transition.Child>

                <div className="fixed inset-0 overflow-y-auto">
                    <div className="flex min-h-full items-center justify-center p-4 text-center">
                        <Transition.Child as={Fragment} enter="ease-out duration-300" enterFrom="opacity-0 scale-95" enterTo="opacity-100 scale-100" leave="ease-in duration-200" leaveFrom="opacity-100 scale-100" leaveTo="opacity-0 scale-95">
                            <Dialog.Panel className="w-full max-w-md transform overflow-hidden rounded-2xl bg-slate-900 p-4 text-left align-middle shadow-xl transition-all border border-slate-800">
                                <div className="text-lg font-medium leading-6 flex justify-between items-center">
                                    <h2 className={'font-bold'}>{title}</h2>
                                    <button className={'rounded-lg bg-red-600/[.2] w-8 h-8 flex justify-center items-center hover:bg-red-600/[.4] duration-150 border border-red-600/[.4]'} onClick={() => callback(false)} >
                                        <XMarkIcon className={'h-6 text-red-700 '} strokeWidth="3" />
                                    </button>
                                </div>
                                {content}
                            </Dialog.Panel>
                        </Transition.Child>
                    </div>
                </div>
            </Dialog>
        </Transition>
    );
}