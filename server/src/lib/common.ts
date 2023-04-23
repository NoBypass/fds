export const getPropertiesAmount = (obj: any): unknown => {
    return Object.keys(obj).reduce((a, key) => a + obj[key], 0)
}

export const firstLetterUpperCase = (s: string): string => {
    return s.replace(/^\w/, (c) => c.toUpperCase())
}

export const encode = (num: number): string => {
    const alphabet = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
    let encoded = ''
    let quotient = num

    do {
        const remainder = quotient % alphabet.length
        encoded = alphabet[remainder] + encoded
        quotient = Math.floor(quotient / alphabet.length)
    } while (quotient > 0)

    return encoded
}

export const shuffle = (str: string): string => {
    const arr = str.split('');
    for (let i = arr.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [arr[i], arr[j]] = [arr[j], arr[i]];
    }
    return arr.join('');
}

export const generateUUID = (): string => {
    return encode(parseInt(shuffle(`${new Date().getTime().toString().substring(5)}${Math.random().toString().substring(2)}`))).substring(0, 11)
}