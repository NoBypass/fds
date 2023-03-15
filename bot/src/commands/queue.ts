import {
    SlashCommandBuilder,
    ChannelType,
    PermissionFlagsBits, TextChannel
} from "discord.js"
import { SlashCommand } from "../types"
import Embed from "../lib/Embed"
// @ts-ignore
import * as config from '../../config.json'
// @ts-ignore
import { client } from "../../index"
import { getTextChannel } from "../lib/common"

const command : SlashCommand = {
    cooldown: 1, // 60 * 60,
    command: new SlashCommandBuilder()
        .setName('queue')
        .setDescription('Ask the discord to play wth you, select the server you want to play on here')
        .setDefaultMemberPermissions(PermissionFlagsBits.SendMessages)
        .addSubcommandGroup(hypixel =>
            hypixel.setName('hypixel')
                .setDescription('Select what you want to play')
                .addSubcommand(bedwars =>
                    bedwars.setName('bedwars')
                        .setDescription('Ask for BedWars')
                        .addStringOption(option =>
                            option.setName('submode')
                                .setDescription('Select which submode you want to play')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'Solo', value: 'solo' },
                                    { name: 'Doubles', value: 'doubles' },
                                    { name: 'Threes', value: 'threes' },
                                    { name: 'Fours', value: 'fours' },
                                    { name: '4v4', value: '4v4' },
                                    { name: 'Dream Doubles', value: 'dream doubles' },
                                    { name: 'Dream Fours', value: 'dream fours' }
                                )
                        )
                        .addStringOption(option =>
                            option.setName('competitiveness')
                                .setDescription('Select whether you want to play competitive or not')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'Public Ranked', value: 'public ranked' },
                                    { name: 'Private Ranked', value: 'private ranked' },
                                    { name: 'Public Queue', value: 'public queue' }
                                )
                        )
                )
                .addSubcommand(duels =>
                    duels.setName('bridge')
                        .setDescription('Ask for Bridge')
                        .addStringOption(option =>
                            option.setName('submode')
                                .setDescription('Select which submode you want to play')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'Solo', value: 'solo' },
                                    { name: 'Doubles', value: 'doubles' },
                                    { name: 'Threes', value: 'threes' },
                                    { name: 'Fours', value: 'fours' },
                                    { name: 'CTF', value: 'ctf' },
                                    { name: '2v2v2v2', value: '2v2v2v2' },
                                    { name: '3v3v3v3', value: '3v3v3v3' }
                                )
                        )
                        .addStringOption(option =>
                            option.setName('competitiveness')
                                .setDescription('Select whether you want to play competitive or not')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'Public Scrims', value: 'public scrims' },
                                    { name: 'Private Scrims', value: 'private scrims' },
                                    { name: 'Public Queue', value: 'public queue' }
                                )
                        )
                )
                .addSubcommand(other =>
                    other.setName('other')
                        .setDescription('Ask for other gamemodes')
                        .addStringOption(option =>
                            option.setName('submode')
                                .setDescription('Select which submode you want to play')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'SkyWars', value: 'skywars' },
                                    { name: 'Party Games', value: 'party games' },
                                    { name: 'Murder Mystery', value: 'murder mystery' },
                                    { name: 'Other', value: 'other' }
                                )
                        )
                        .addStringOption(option =>
                            option.setName('publicity')
                                .setDescription('Select whether you want to play public games or not')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'Private Games', value: 'private games' },
                                    { name: 'Public Queue', value: 'public queue' }
                                )
                        )
                )
        )
        .addSubcommandGroup(bucky =>
            bucky.setName('bucky-tour')
                .setDescription('Select what you want to play')
                .addSubcommand(bridge =>
                    bridge.setName('bridge')
                        .setDescription('Ask for to play bridge')
                        .addStringOption(option =>
                            option.setName('submode')
                                .setDescription('Select which submode you want to play')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'Solo', value: 'solo' },
                                    { name: 'Doubles', value: 'doubles' },
                                    { name: 'Threes', value: 'threes' },
                                    { name: 'Fours', value: 'fours' }
                                )
                        )
                        .addStringOption(option =>
                            option.setName('region')
                                .setDescription('Select which region you would prefer')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'NA', value: 'north america' },
                                    { name: 'EU', value: 'europe' },
                                    { name: 'Don\'t care', value: 'dont care' }
                                )
                        )
                )
        )
        .addSubcommandGroup(minemen =>
            minemen.setName('minemen')
                .setDescription('Select what you want to play')
                .addSubcommand(bridge =>
                    bridge.setName('bridge')
                        .setDescription('Ask for to play bridge')
                        .addStringOption(option =>
                            option.setName('submode')
                                .setDescription('Select which submode you want to play')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'Solo', value: 'solo' },
                                    { name: 'Doubles', value: 'doubles' },
                                    { name: 'Other', value: 'other' }
                                )
                        )
                        .addStringOption(option =>
                            option.setName('region')
                                .setDescription('Select which region you would prefer')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'NA', value: 'north america' },
                                    { name: 'EU', value: 'europe' },
                                    { name: 'Don\'t care', value: 'dont care' }
                                )
                        )
                )
                .addSubcommand(bedfight =>
                    bedfight.setName('bedfight')
                        .setDescription('Ask for to play bedfight')
                        .addStringOption(option =>
                            option.setName('submode')
                                .setDescription('Select which submode you want to play')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'Solo', value: 'solo' },
                                    { name: 'Doubles', value: 'doubles' },
                                    { name: 'Other', value: 'other' }
                                )
                        )
                        .addStringOption(option =>
                            option.setName('region')
                                .setDescription('Select which region you would prefer')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'NA', value: 'north america' },
                                    { name: 'EU', value: 'europe' },
                                    { name: 'Don\'t care', value: 'dont care' }
                                )
                        )
                )
                .addSubcommand(uhc =>
                    uhc.setName('uhc')
                        .setDescription('Ask for to play bridge')
                        .addStringOption(option =>
                            option.setName('submode')
                                .setDescription('Select which submode you want to play')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'Solo', value: 'solo' },
                                    { name: 'Doubles', value: 'doubles' },
                                    { name: 'Other', value: 'other' }
                                )
                        )
                        .addStringOption(option =>
                            option.setName('region')
                                .setDescription('Select which region you would prefer')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'NA', value: 'north america' },
                                    { name: 'EU', value: 'europe' },
                                    { name: 'Don\'t care', value: 'dont care' }
                                )
                        )
                )
                .addSubcommand(battlerush =>
                    battlerush.setName('battlerush')
                        .setDescription('Ask for to play bridge')
                        .addStringOption(option =>
                            option.setName('submode')
                                .setDescription('Select which submode you want to play')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'Solo', value: 'solo' },
                                    { name: 'Doubles', value: 'doubles' },
                                    { name: 'Other', value: 'other' }
                                )
                        )
                        .addStringOption(option =>
                            option.setName('region')
                                .setDescription('Select which region you would prefer')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'NA', value: 'north america' },
                                    { name: 'EU', value: 'europe' },
                                    { name: 'Don\'t care', value: 'dont care' }
                                )
                        )
                )
                .addSubcommand(other =>
                    other.setName('other')
                        .setDescription('Ask for to play other games')
                        .addStringOption(option =>
                            option.setName('submode')
                                .setDescription('Select which submode you want to play')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'Solo', value: 'solo' },
                                    { name: 'Doubles', value: 'doubles' },
                                    { name: 'Other', value: 'other' }
                                )
                        )
                        .addStringOption(option =>
                            option.setName('region')
                                .setDescription('Select which region you would prefer')
                                .setRequired(true)
                                .addChoices(
                                    { name: 'NA', value: 'north america' },
                                    { name: 'EU', value: 'europe' },
                                    { name: 'Don\'t care', value: 'dont care' }
                                )
                        )
                )
        ),

    execute: interaction => {
        const channel: TextChannel = getTextChannel(interaction, config.channels.lookingToPlay)
        const server = interaction.options.data[0]
        const mode = server.options?.[0]
        const submode = mode?.options?.filter(option => option.name == 'submode')[0]
        const region = mode?.options?.filter(option => option.name == 'region')[0]
        const competitiveness = mode?.options?.filter(option => option.name == 'competitiveness')[0]

        let publicEmbed = new Embed({ channel })
        publicEmbed.title(`${interaction.user.username} wants to play:`)
        publicEmbed.description(`<@${interaction.user.id}> did \`\`/queue\`\`\n
            **Server:** ${server.name}
            **Mode:** ${mode?.name}
            **Submode:** ${submode?.value}
            ${region != null?
            `**Region:** ${region?.value}`:
            `**Competitiveness** ${competitiveness?.value}`}`)
        channel.send({
            content: `${mode?.name == 'bridge'? `<@${config.roles.bridge}->`: ''}${mode?.name == 'bedwars'? `<@${config.roles.bedwars}->`: ''}`,
            embeds: [publicEmbed.get() as any]
        })

        const privateEmbed = new Embed({ interaction })
        privateEmbed.title(`Created queue request in <#${channel.id}>`)
        interaction.reply({ embeds: [privateEmbed.get() as any], ephemeral: true })
    }
}

export default command