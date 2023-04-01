package com.fds.backend.security;

import java.util.Objects;

public class JwtResponseDTO extends AuthResponseDTO {
    private String accessToken;
    private String tokenType = "Bearer";

    public JwtResponseDTO(String accessToken, Integer id, String username) {
        super(id, username);
        this.accessToken = accessToken;
    }

    public String getAccessToken() {
        return accessToken;
    }

    public void setAccessToken(String accessToken) {
        this.accessToken = accessToken;
    }

    public String getTokenType() {
        return tokenType;
    }

    public void setTokenType(String tokenType) {
        this.tokenType = tokenType;
    }

    @Override
    public boolean equals(Object obj) {
        if (this == obj) {
            return true;
        }
        if (!(obj instanceof JwtResponseDTO jwtResponseDTO)) {
            return false;
        }

        return super.equals(obj)
                && accessToken.equals(jwtResponseDTO.accessToken);
    }

    @Override
    public int hashCode() {
        return Objects.hash(super.hashCode(), accessToken);
    }
}
