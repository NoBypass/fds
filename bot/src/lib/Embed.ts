import {Channel, CommandInteraction, EmbedData, EmbedField, TextChannel} from 'discord.js'

export default class Embed {
    private readonly interaction: CommandInteraction | null = null
    private readonly channel: TextChannel | null = null
    private embed: EmbedData = {
        color: 0x2F3136,
        footer: {
            iconURL: 'https://cdn.discordapp.com/avatars/672835870080106509/d200b738f793b6ba3fce6f207cac8b6b.webp',
            text: 'Bot by NoBypass'
        },
        fields: []
    }
    constructor(props?: { interaction?: CommandInteraction, channel?: TextChannel }) {
        if (props) {
            if (props.interaction) this.interaction = props.interaction
            if (props.channel) this.channel = props.channel
        }
    }

    public color(hex: number) {
        this.embed.color = hex
    }
    public description(text: string) {
        this.embed.description = text
    }
    public title(text: string) {
        this.embed.title = text
    }
    public titleUrl(url: string) {
        this.embed.url = url
    }
    public thumbnail(url: string) {
        this.embed.thumbnail = {
            url
        }
    }
    public image(url: string) {
        this.embed.image = {
            url
        }
    }
    public time(time?: string) {
        this.embed.timestamp = time == null? new Date().toISOString(): time
    }
    public field(props: EmbedField) {
        this.embed.fields?.push(props)
    }
    public fields(props: EmbedField[]) {
        props.forEach((prop) => {
            this.field(prop)
        })
    }
    public get() {
        return this.embed
    }
    public send() {
        if (this.channel) this.channel.send({ embeds: [this.embed as any] })
        else throw new Error('Cannot send message, no channel was given')
    }
    public reply() {
        if (this.interaction) this.interaction.reply({ embeds: [this.embed as any] })
        else throw new Error('Cannot reply, no interaction was given')
    }
}