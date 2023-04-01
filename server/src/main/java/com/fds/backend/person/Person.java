package com.fds.backend.person;

import com.fds.backend.discordUser.DiscordUser;

import javax.persistence.*;
import java.util.HashSet;
import java.util.Objects;
import java.util.Set;

@Entity
public class Person {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;

    @Column(unique = true)
    private String username;

    public Set<DiscordUser> getItems() {
        return discordUsers;
    }

    public void setItems(Set<DiscordUser> discordUsers) {
        this.discordUsers = discordUsers;
    }

    @OneToMany(mappedBy = "person", fetch = FetchType.LAZY)
    private Set<DiscordUser> discordUsers = new HashSet<>();

    private String password;

    public Person() {
    }

    public Person(Integer id) {
        this.id = id;
    }

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (o == null || getClass() != o.getClass()) {
            return false;
        }
        Person that = (Person) o;
        return Objects.equals(id, that.id);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id);
    }

}
