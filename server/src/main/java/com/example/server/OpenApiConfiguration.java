package com.example.server;

import com.example.server.security.SecurityConstants;
import io.swagger.v3.oas.annotations.OpenAPIDefinition;
import io.swagger.v3.oas.annotations.info.Info;
import io.swagger.v3.oas.annotations.security.SecurityRequirement;
import org.springframework.context.annotation.Configuration;


@Configuration
@OpenAPIDefinition(info = @Info(title = "The Todolist API Backend built @ Bbc", version = "1.0.0"),
        security = @SecurityRequirement(name = SecurityConstants.AUTHORIZATION_HEADER_NAME)
)
public class OpenApiConfiguration {
}
