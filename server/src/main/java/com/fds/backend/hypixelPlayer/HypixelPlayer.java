package com.fds.backend.hypixelPlayer;

import com.fds.backend.discordUser.DiscordUser;
import com.fds.backend.mojangUser.MojangUser;

import javax.persistence.*;
import java.sql.Timestamp;
import java.util.Objects;

@Entity
public class HypixelPlayer {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;

    @Column(columnDefinition = "json")
    private String latestPlayerStat;
    private Timestamp latestLookup;
    private Timestamp firstLookup;
    private Boolean tracking;
    @OneToOne(mappedBy = "hypixelPlayer", fetch = FetchType.LAZY)
    private DiscordUser discordUser;
    @OneToOne
    @JoinColumn(name = "mojangUser_id")
    private MojangUser mojangUser;

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        HypixelPlayer that = (HypixelPlayer) o;
        return id == that.id && Objects.equals(latestPlayerStat, that.latestPlayerStat) && Objects.equals(latestLookup, that.latestLookup) && Objects.equals(firstLookup, that.firstLookup) && Objects.equals(tracking, that.tracking) && Objects.equals(discordUser, that.discordUser) && Objects.equals(mojangUser, that.mojangUser);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id, latestPlayerStat, latestLookup, firstLookup, tracking, discordUser, mojangUser);
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getLatestPlayerStat() {
        return latestPlayerStat;
    }

    public void setLatestPlayerStat(String latestPlayerStat) {
        this.latestPlayerStat = latestPlayerStat;
    }

    public Timestamp getLatestLookup() {
        return latestLookup;
    }

    public void setLatestLookup(Timestamp latestLookup) {
        this.latestLookup = latestLookup;
    }

    public Timestamp getFirstLookup() {
        return firstLookup;
    }

    public void setFirstLookup(Timestamp firstLookup) {
        this.firstLookup = firstLookup;
    }

    public Boolean getTracking() {
        return tracking;
    }

    public void setTracking(Boolean tracking) {
        this.tracking = tracking;
    }

    public DiscordUser getDiscordUser() {
        return discordUser;
    }

    public void setDiscordUser(DiscordUser discordUser) {
        this.discordUser = discordUser;
    }

    public MojangUser getMojangUser() {
        return mojangUser;
    }

    public void setMojangUser(MojangUser mojangUser) {
        this.mojangUser = mojangUser;
    }
}
