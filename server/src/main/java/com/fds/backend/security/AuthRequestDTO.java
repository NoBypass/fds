package com.fds.backend.security;

import org.hibernate.validator.constraints.Length;

import javax.validation.constraints.NotBlank;
import java.util.Objects;

public class AuthRequestDTO {
    @NotBlank(message = "username must not be empty")
    private String username;
    @NotBlank(message = "password must not be empty")
    @Length(min = 8, max = 255, message = "length must be between 8 and 255")
    private String password;

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
    public boolean equals(Object obj) {
        if (this == obj) {
            return true;
        }
        if (!(obj instanceof AuthRequestDTO authRequestDTO)) {
            return false;
        }

        return username.equals(authRequestDTO.username)
                && password.equals(authRequestDTO.password);
    }

    @Override
    public int hashCode() {
        return Objects.hash(username, password);
    }
}
