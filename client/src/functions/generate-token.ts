export default function generateToken(len:number) {
    let result = ''
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+"*ç%&/()=!£$'
    for (let i = 0; i < len; i++) result += characters.charAt(Math.floor(Math.random() * characters.length))
    return result
}