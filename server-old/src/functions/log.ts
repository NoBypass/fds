const colors: {[key: string]: any} = {
    s: {
        reset: "\x1b[0m",
        bright: "\x1b[1m",
        dim: "\x1b[2m",
        underscore: "\x1b[4m",
        blink: "\x1b[5m",
        reverse: "\x1b[7m",
        hidden: "\x1b[8m"
    },
    c: {
        black: "\x1b[30m",
        red: "\x1b[31m",
        green: "\x1b[32m",
        yellow: "\x1b[33m",
        blue: "\x1b[34m",
        magenta: "\x1b[35m",
        cyan: "\x1b[36m",
        white: "\x1b[37m",
        gray: "\x1b[90m"
    },
    bg: {
        black: "\x1b[40m",
        red: "\x1b[41m",
        green: "\x1b[42m",
        yellow: "\x1b[43m",
        blue: "\x1b[44m",
        magenta: "\x1b[45m",
        cyan: "\x1b[46m",
        white: "\x1b[47m",
        gray: "\x1b[100m"
    }
}

export default function log(data:string, style?:string, styles?:string[], brackets?:boolean):void {
    const allStyles = getAllStyles()
    if (brackets == undefined) brackets = true
    if (style != undefined && !allStyles.includes(style)) console.log(new Error(`Style ${style} does not exist`))
    if (style != undefined) {
        styles?.forEach(style => {
            if (!allStyles.includes(style)) console.log(new Error(`Style ${style} does not exist`))
        })
    }
    if (style == undefined && styles == undefined) console.log(new Error('Cannot have both style and styles undefined'))

    let toLog = ''
    if (brackets) toLog += '|| '
    if (style != undefined) toLog += getStyle(style)
    if (styles != undefined) {
        styles.forEach(style => {
            toLog += getStyle(style)
        })
    }
    toLog += data + colors.s.reset
    return console.log(toLog)
}

function getStyle(style: string): string | undefined {
    let stylePath = style.split('_');
    if (stylePath[0] in colors) {
        let newColors = colors[stylePath[0]];
        return newColors[stylePath[1]];
    }
    return undefined;
}

function getAllStyles(): string[] {
    let allStyles: string[] = [];
    for (let [colorType, color] of Object.entries(colors)) {
        for (let [colorName] of Object.entries(color)) {
            allStyles.push(colorType + '_' + colorName);
        }
    }
    return allStyles;
}