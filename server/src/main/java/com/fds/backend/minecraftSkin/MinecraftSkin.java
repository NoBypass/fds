package com.fds.backend.minecraftSkin;

import com.fds.backend.mojangUser.MojangUser;

import javax.persistence.*;
import java.util.HashSet;
import java.util.Objects;
import java.util.Set;

@Entity
public class MinecraftSkin {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;
    private String skinBase64;
    @OneToMany(mappedBy = "minecraftSkin", fetch = FetchType.LAZY)
    private Set<MojangUser> mojangUsers = new HashSet<>();

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        MinecraftSkin that = (MinecraftSkin) o;
        return id == that.id && Objects.equals(skinBase64, that.skinBase64) && Objects.equals(mojangUsers, that.mojangUsers);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id, skinBase64, mojangUsers);
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getSkinBase64() {
        return skinBase64;
    }

    public void setSkinBase64(String skinBase64) {
        this.skinBase64 = skinBase64;
    }

    public Set<MojangUser> getMojangUsers() {
        return mojangUsers;
    }

    public void setMojangUsers(Set<MojangUser> mojangUsers) {
        this.mojangUsers = mojangUsers;
    }
}
