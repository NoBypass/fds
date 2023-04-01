package com.fds.backend.tag;

import com.fds.backend.discordUser.DiscordUser;

import javax.persistence.*;
import java.util.Objects;
import java.util.Set;

@Entity
public class Tag {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;
    @ManyToMany(mappedBy = "tags")
    private Set<DiscordUser> linkedDiscordUsers;
    @Column(unique = true)
    private String name;

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (!(o instanceof Tag tag)) return false;
        return Objects.equals(id, tag.id) && Objects.equals(linkedDiscordUsers, tag.linkedDiscordUsers) && Objects.equals(name, tag.name);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id, linkedDiscordUsers, name);
    }

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public Set<DiscordUser> getLinkedItems() {
        return linkedDiscordUsers;
    }

    public void setLinkedItems(Set<DiscordUser> linkedDiscordUsers) {
        this.linkedDiscordUsers = linkedDiscordUsers;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
