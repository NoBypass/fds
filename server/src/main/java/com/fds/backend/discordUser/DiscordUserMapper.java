package com.fds.backend.discordUser;

public class DiscordUserMapper {
    public static DiscordUserResponseDTO toDTO(DiscordUser discordUser) {
        DiscordUserResponseDTO discordUserResponseDTO = new DiscordUserResponseDTO();
        discordUserResponseDTO.setId(discordUser.getId());
        discordUserResponseDTO.setDailiesStreak(discordUser.getDailiesStreak());
        discordUserResponseDTO.setLevel(discordUser.getLevel());
        discordUserResponseDTO.setHypixelPlayerId(discordUser.getHypixelPlayer().getId());
        discordUserResponseDTO.setMessagesSent(discordUser.getMessagesSent());
        discordUserResponseDTO.setLastDailyClaimed(discordUser.getLastDailyClaimed());
        discordUserResponseDTO.setMinutesSpentInVc(discordUser.getMinutesSpentInVc());
        discordUserResponseDTO.setOverflowXp(discordUser.getOverflowXp());
        discordUserResponseDTO.setXpFromDailies(discordUser.getXpFromDailies());
        return discordUserResponseDTO;
    }

    public static DiscordUser fromDTO(DiscordUserRequestDTO discordUserRequestDTO) {
        DiscordUser discordUser = new DiscordUser();
        discordUser.setDailiesStreak(discordUserRequestDTO.getDailiesStreak());
        discordUser.setLevel(discordUserRequestDTO.getLevel());
        discordUser.setLastDailyClaimed(discordUserRequestDTO.getLastDailyClaimed());
        discordUser.setOverflowXp(discordUserRequestDTO.getOverflowXp());
        discordUser.setMessagesSent(discordUserRequestDTO.getMessagesSent());
        discordUser.setXpFromDailies(discordUserRequestDTO.getXpFromDailies());
        discordUser.setMinutesSpentInVc(discordUser.getMinutesSpentInVc());
        return discordUser;
    }
}
