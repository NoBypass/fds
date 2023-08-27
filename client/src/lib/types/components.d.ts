export type ColorProp = 'primary' | 'secondary' | 'success' | 'danger' | 'warning' | 'info'

export type Alert = {
    id: number
    message: string
    color: ColorProp
    icon: boolean
    stage: number
}
