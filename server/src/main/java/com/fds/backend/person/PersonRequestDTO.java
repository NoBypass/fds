package com.fds.backend.person;

import org.hibernate.validator.constraints.Length;

import javax.validation.constraints.NotBlank;
import java.util.Objects;

public class PersonRequestDTO {
    @NotBlank(message = "must not be blank")
    private String username;

    @NotBlank(message = "must not be blank")
    @Length(min = 6, max = 255, message = "length must be between 6 and 255")
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
        if (!(obj instanceof PersonRequestDTO personRequestDTO)) {
            return false;
        }

        return username.equals(personRequestDTO.username)
                && password.equals(personRequestDTO.password);
    }

    @Override
    public int hashCode() {
        return Objects.hash(username, password);
    }
}
