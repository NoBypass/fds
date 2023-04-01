package com.fds.backend.discordUser;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import java.util.List;

public interface DiscordUserRepository extends JpaRepository<DiscordUser, Integer> {
    @Query("SELECT d FROM DiscordUser d WHERE d.hypixelPlayer IN "
            + "(SELECT h FROM HypixelPlayer h INNER JOIN MojangUser m ON h.mojangUser = m WHERE m.name LIKE CONCAT('%', :name, '%'))")
    List<DiscordUser> findByName(String name);
}