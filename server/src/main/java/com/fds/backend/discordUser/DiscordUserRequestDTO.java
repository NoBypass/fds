package com.fds.backend.discordUser;

import javax.validation.constraints.NotBlank;
import java.sql.Timestamp;
import java.util.Objects;

public class DiscordUserRequestDTO {
    @NotBlank
    private Integer level;
    @NotBlank
    private Integer overflowXp;
    @NotBlank
    private Integer dailiesStreak;
    @NotBlank
    private Integer xpFromDailies;
    @NotBlank
    private Timestamp lastDailyClaimed;
    @NotBlank
    private Integer minutesSpentInVc;
    @NotBlank
    private Integer messagesSent;
    @NotBlank
    private Integer hypixelPlayerId;

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        DiscordUserRequestDTO that = (DiscordUserRequestDTO) o;
        return Objects.equals(level, that.level) && Objects.equals(overflowXp, that.overflowXp) && Objects.equals(dailiesStreak, that.dailiesStreak) && Objects.equals(xpFromDailies, that.xpFromDailies) && Objects.equals(lastDailyClaimed, that.lastDailyClaimed) && Objects.equals(minutesSpentInVc, that.minutesSpentInVc) && Objects.equals(messagesSent, that.messagesSent) && Objects.equals(hypixelPlayerId, that.hypixelPlayerId);
    }

    @Override
    public int hashCode() {
        return Objects.hash(level, overflowXp, dailiesStreak, xpFromDailies, lastDailyClaimed, minutesSpentInVc, messagesSent, hypixelPlayerId);
    }

    public Integer getLevel() {
        return level;
    }

    public void setLevel(Integer level) {
        this.level = level;
    }

    public Integer getOverflowXp() {
        return overflowXp;
    }

    public void setOverflowXp(Integer overflowXp) {
        this.overflowXp = overflowXp;
    }

    public Integer getDailiesStreak() {
        return dailiesStreak;
    }

    public void setDailiesStreak(Integer dailiesStreak) {
        this.dailiesStreak = dailiesStreak;
    }

    public Integer getXpFromDailies() {
        return xpFromDailies;
    }

    public void setXpFromDailies(Integer xpFromDailies) {
        this.xpFromDailies = xpFromDailies;
    }

    public Timestamp getLastDailyClaimed() {
        return lastDailyClaimed;
    }

    public void setLastDailyClaimed(Timestamp lastDailyClaimed) {
        this.lastDailyClaimed = lastDailyClaimed;
    }

    public Integer getMinutesSpentInVc() {
        return minutesSpentInVc;
    }

    public void setMinutesSpentInVc(Integer minutesSpentInVc) {
        this.minutesSpentInVc = minutesSpentInVc;
    }

    public Integer getMessagesSent() {
        return messagesSent;
    }

    public void setMessagesSent(Integer messagesSent) {
        this.messagesSent = messagesSent;
    }

    public Integer getHypixelPlayerId() {
        return hypixelPlayerId;
    }

    public void setHypixelPlayerId(Integer hypixelPlayerId) {
        this.hypixelPlayerId = hypixelPlayerId;
    }
}
