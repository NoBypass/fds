import {
    DevicePhoneMobileIcon,
    GlobeAltIcon, MagnifyingGlassIcon,
    UserCircleIcon,
    VideoCameraIcon,
    Bars3Icon
} from '@heroicons/react/24/outline'
import { Modal } from './Modal'
import {useEffect, useState} from "react"
import Login from "./modals/Login";
import SelectionBox from "./SelectionBox";

export default function TopBar() {
    const search = [
        {name: 'Hypixel', icon: <DevicePhoneMobileIcon/>, placeholder: 'for player'},
        {name: 'Web', icon: <GlobeAltIcon/>, placeholder: 'the web'},
        {name: 'YouTube', icon: <VideoCameraIcon/>, placeholder: 'on Youtube'}
    ]
    const [title, setTitle] = useState('Login / Signup')
    const [userModalOpen, setUserModalOpen] = useState(false)
    const [selectedSearch, setSelectedSearch] = useState(search[0])

    const options = ['Home', 'Player', 'Leaderboard', 'Graphs / Comparison']

    // TODO: If valid session token, open settings menu, otherwise open login/signup screen
    // TODO: Add handling for themes

    return (
        <>
            <nav className={'border-b border-b-slate-800 h-16 w-full flex justify-around py-2 items-center min-[1625px]:px-48 min-[1515px]:px-32 min-[1375]:px-16 px-4'}>
                <div className={'flex flex-row gap-3 items-center h-full'}>
                    <img src={'/src/assets/graphics/logo.svg'} alt="logo" className={'overflow-visible object-cover h-full LOGO'}/>
                    <h1 className={'font-title text-3xl LOGO_FONT'}>FDS</h1>
                </div>
                <form className={'flex flex-row-reverse relative w-128 min-w-0 relative'}>
                    <input type="text" className={'absolute rounded-full bg-blue-600/[.1] hover:bg-blue-600/[.2] border border-blue-600/[.8] placeholder:italic placeholder:opacity-25 pl-2 focus:outline-offset-2 focus:outline-blue-600 focus:outline outline-2 hover:shadow-md duration-100 hover:shadow-blue-600/[.2] h-8 w-full'} placeholder={'Search ' + selectedSearch.placeholder}/>
                    <SelectionBox selected={selectedSearch} setSelected={setSelectedSearch} options={search} />
                    <button className={'h-8 w-10 -skew-x-20 border border-blue-600/[.7] hover:shadow-md duration-100 hover:shadow-blue-600/[.2] hover:bg-blue-600/[.2] flex justify-center items-center'}>
                        <MagnifyingGlassIcon className={'skew-x-20 h-5 w-5'}/>
                    </button>
                </form>
                <ul className={'font-semibold flex-row items-center flex gap-10 LIST_GAP'}>
                    {options.map((x, i) => <li key={i} className={'TOPBAR_LINK'}><p>{x}</p></li>)}
                    <li className={'TOPBAR_LINK_BURGER hidden'}>
                        <Bars3Icon className={'text-white h-8'}/>
                    </li>
                    <li>
                        <div className={'w-px h-8 border-l border-slate-800 TOPBAR_LINK'}/>
                    </li>
                    <li className={'flex items-center TOPBAR_LINK'}>
                        <button onClick={() => setUserModalOpen(true)}>
                            <UserCircleIcon className={'text-white h-8'}/>
                        </button>
                    </li>
                </ul>
            </nav>
            <Modal content={<Login setTitle={setTitle} />} open={userModalOpen} callback={setUserModalOpen} title={title} />
        </>
    )
}