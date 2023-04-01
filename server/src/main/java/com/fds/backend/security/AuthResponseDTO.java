package com.fds.backend.security;

import java.util.Objects;

public class AuthResponseDTO {
    private Integer id;
    private String username;

    public AuthResponseDTO() {
    }


    public AuthResponseDTO(Integer id, String username) {
        this.id = id;
        this.username = username;
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

    @Override
    public boolean equals(Object obj) {
        if (this == obj) {
            return true;
        }
        if (!(obj instanceof AuthResponseDTO authResponseDTO)) {
            return false;
        }

        return id.equals(authResponseDTO.id)
                && username.equals(authResponseDTO.username);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id, username);
    }
}
