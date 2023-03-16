package com.example.server.security;

import javax.crypto.spec.SecretKeySpec;
import java.nio.charset.StandardCharsets;

public class SecurityConstants {
    public static final String SECRET = "Secret Key to generate JWT's (min 256 Bits)";
    public static final String ALGORITHM = "HmacSHA256";
    public static final SecretKeySpec SECRET_KEY_SPEC = new SecretKeySpec(SECRET.getBytes(StandardCharsets.UTF_8), ALGORITHM);
    public static final long EXPIRATION_TIME = 864_000_000L; // 10 days
    public static final String AUTH_URLS = "/auth/*";
    public static final String[] API_DOCUMENTATION_URLS = {
            "/v3/api-docs/**",
            "/swagger-ui.html",
            "/swagger-ui/**"
    };

    public static final String AUTHORIZATION_HEADER_NAME = "Authorization";

}
