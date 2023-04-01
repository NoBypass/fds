package com.fds.backend.discordUser;

import com.fds.backend.hypixelPlayer.HypixelPlayer;

import javax.persistence.*;
import java.sql.Timestamp;
import java.util.Objects;

@Entity
public class DiscordUser {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;
    private Integer level;
    private Integer overflowXp;
    private Integer dailiesStreak;
    private Integer xpFromDailies;
    private Timestamp lastDailyClaimed;
    private Integer minutesSpentInVc;
    private Integer messagesSent;
    @OneToOne
    @JoinColumn(name = "hypixelPlayer_id")
    private HypixelPlayer hypixelPlayer;

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        DiscordUser that = (DiscordUser) o;
        return id == that.id && level == that.level && overflowXp == that.overflowXp && dailiesStreak == that.dailiesStreak && xpFromDailies == that.xpFromDailies && minutesSpentInVc == that.minutesSpentInVc && messagesSent == that.messagesSent && Objects.equals(lastDailyClaimed, that.lastDailyClaimed) && Objects.equals(hypixelPlayer, that.hypixelPlayer);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id, level, overflowXp, dailiesStreak, xpFromDailies, lastDailyClaimed, minutesSpentInVc, messagesSent, hypixelPlayer);
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public int getLevel() {
        return level;
    }

    public void setLevel(int level) {
        this.level = level;
    }

    public int getOverflowXp() {
        return overflowXp;
    }

    public void setOverflowXp(int overflowXp) {
        this.overflowXp = overflowXp;
    }

    public int getDailiesStreak() {
        return dailiesStreak;
    }

    public void setDailiesStreak(int dailiesStreak) {
        this.dailiesStreak = dailiesStreak;
    }

    public int getXpFromDailies() {
        return xpFromDailies;
    }

    public void setXpFromDailies(int xpFromDailies) {
        this.xpFromDailies = xpFromDailies;
    }

    public Timestamp getLastDailyClaimed() {
        return lastDailyClaimed;
    }

    public void setLastDailyClaimed(Timestamp lastDailyClaimed) {
        this.lastDailyClaimed = lastDailyClaimed;
    }

    public int getMinutesSpentInVc() {
        return minutesSpentInVc;
    }

    public void setMinutesSpentInVc(int minutesSpentInVc) {
        this.minutesSpentInVc = minutesSpentInVc;
    }

    public int getMessagesSent() {
        return messagesSent;
    }

    public void setMessagesSent(int messagesSent) {
        this.messagesSent = messagesSent;
    }

    public HypixelPlayer getHypixelPlayer() {
        return hypixelPlayer;
    }

    public void setHypixelPlayer(HypixelPlayer hypixelPlayer) {
        this.hypixelPlayer = hypixelPlayer;
    }
}
