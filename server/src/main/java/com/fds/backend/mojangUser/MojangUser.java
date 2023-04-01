package com.fds.backend.mojangUser;

import com.fds.backend.hypixelPlayer.HypixelPlayer;
import com.fds.backend.minecraftSkin.MinecraftSkin;

import javax.persistence.*;
import java.util.Objects;

@Entity
public class MojangUser {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;
    private String uuid;
    private String name;
    @ManyToOne
    @JoinColumn(name = "minecraftSkin_id")
    private MinecraftSkin minecraftSkin;
    @OneToOne(mappedBy = "mojangUser", fetch = FetchType.LAZY)
    private HypixelPlayer hypixelPlayer;

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        MojangUser that = (MojangUser) o;
        return id == that.id && Objects.equals(uuid, that.uuid) && Objects.equals(name, that.name) && Objects.equals(minecraftSkin, that.minecraftSkin) && Objects.equals(hypixelPlayer, that.hypixelPlayer);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id, uuid, name, minecraftSkin, hypixelPlayer);
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getUuid() {
        return uuid;
    }

    public void setUuid(String uuid) {
        this.uuid = uuid;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public MinecraftSkin getMinecraftSkin() {
        return minecraftSkin;
    }

    public void setMinecraftSkin(MinecraftSkin minecraftSkin) {
        this.minecraftSkin = minecraftSkin;
    }

    public HypixelPlayer getHypixelPlayer() {
        return hypixelPlayer;
    }

    public void setHypixelPlayer(HypixelPlayer hypixelPlayer) {
        this.hypixelPlayer = hypixelPlayer;
    }
}
