package com.fds.backend.discordUser;

import java.sql.Timestamp;
import java.util.Objects;

public class DiscordUserResponseDTO extends DiscordUserRequestDTO {
    private Integer id;
    private Integer level;
    private Integer overflowXp;
    private Integer dailiesStreak;
    private Integer xpFromDailies;
    private Timestamp lastDailyClaimed;
    private Integer minutesSpentInVc;
    private Integer messagesSent;
    private Integer hypixelPlayerId;

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        if (!super.equals(o)) return false;
        DiscordUserResponseDTO that = (DiscordUserResponseDTO) o;
        return Objects.equals(id, that.id) && Objects.equals(level, that.level) && Objects.equals(overflowXp, that.overflowXp) && Objects.equals(dailiesStreak, that.dailiesStreak) && Objects.equals(xpFromDailies, that.xpFromDailies) && Objects.equals(lastDailyClaimed, that.lastDailyClaimed) && Objects.equals(minutesSpentInVc, that.minutesSpentInVc) && Objects.equals(messagesSent, that.messagesSent) && Objects.equals(hypixelPlayerId, that.hypixelPlayerId);
    }

    @Override
    public int hashCode() {
        return Objects.hash(super.hashCode(), id, level, overflowXp, dailiesStreak, xpFromDailies, lastDailyClaimed, minutesSpentInVc, messagesSent, hypixelPlayerId);
    }

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
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
