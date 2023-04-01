package com.fds.backend.tag;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.List;

public interface TagRepository extends JpaRepository<Tag, Integer> {
    @Query("SElECT t FROM Tag t where t.name LIKE CONCAT('%', :name, '%')")
    List<Tag> findByName(String name);
}
