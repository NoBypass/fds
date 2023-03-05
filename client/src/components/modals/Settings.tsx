import React from "react"
import { Foldable } from "../Foldable"

export function UserSettings() {
    const themeContent = (
        <p>test</p>
    )

    return (
        <>
            <Foldable title={'Account'} content={themeContent} />
            <Foldable title={'Manage public hypixel stats'} content={themeContent} />
            <Foldable title={'Default values'} content={themeContent} />
            <Foldable title={'Link 2nd Account'} content={themeContent} />
            <Foldable title={'Functional customization'} content={themeContent} />
            <Foldable title={'Theme'} content={themeContent} />
        </>
    )
}