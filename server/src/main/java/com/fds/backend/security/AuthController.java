package com.fds.backend.security;

import com.fds.backend.person.PersonResponseDTO;
import com.fds.backend.person.PersonService;
import com.nimbusds.jose.JOSEException;
import com.nimbusds.jose.JWSAlgorithm;
import com.nimbusds.jose.JWSHeader;
import com.nimbusds.jose.JWSSigner;
import com.nimbusds.jose.crypto.MACSigner;
import com.nimbusds.jwt.JWTClaimsSet;
import com.nimbusds.jwt.SignedJWT;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.responses.ApiResponses;
import io.swagger.v3.oas.annotations.security.SecurityRequirements;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.User;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.server.ResponseStatusException;

import javax.validation.Valid;
import java.util.Date;

@RestController
@RequestMapping(AuthController.PATH)
public class AuthController {
    public static final String PATH = "/auth";

    private final PersonService personService;
    private final AuthenticationManager authenticationManager;

    @Autowired
    public AuthController(PersonService personService, AuthenticationManager authenticationManager) {
        this.personService = personService;
        this.authenticationManager = authenticationManager;
    }

    private static String generateJwtToken(String userName) {
        try {
            JWSSigner signer = new MACSigner(SecurityConstants.SECRET_KEY_SPEC);
            JWTClaimsSet claimsSet = new JWTClaimsSet.Builder()
                    .subject(userName)
                    .expirationTime(new Date(System.currentTimeMillis() + SecurityConstants.EXPIRATION_TIME))
                    .build();

            SignedJWT jwt = new SignedJWT(new JWSHeader(JWSAlgorithm.HS256), claimsSet);
            jwt.sign(signer);
            return jwt.serialize();

        } catch (JOSEException e) {
            throw new RuntimeException(e);
        }
    }

    @PostMapping("/signup")
    @Operation(summary = "Create a new user")
    @ApiResponses(value = {
            @ApiResponse(responseCode = "201", description = "User was created successfully",
                    content = @Content(schema = @Schema(implementation = AuthResponseDTO.class))),
            @ApiResponse(responseCode = "409", description = "User could not be created, username already in use",
                    content = @Content)
    })
    @SecurityRequirements //no security here, default is BEARER
    public ResponseEntity<?> signUp(@Valid @RequestBody AuthRequestDTO authRequestDTO) {
        try {
            PersonResponseDTO personResponseDTO = personService.create(authRequestDTO);

            AuthResponseDTO authResponseDTO = new AuthResponseDTO(personResponseDTO.getId(), personResponseDTO.getUsername());

            return ResponseEntity.status(201).body(authResponseDTO);
        } catch (DataIntegrityViolationException e) {
            throw new ResponseStatusException(HttpStatus.CONFLICT, "User could not be created, username already in use");
        }
    }

    @PostMapping("/signin")
    @Operation(summary = "Receive a token for BEARER authorization")
    @SecurityRequirements //no security here, default is BEARER
    @ApiResponses(value = {
            @ApiResponse(responseCode = "200", description = "Login successful",
                    content = @Content(schema = @Schema(implementation = JwtResponseDTO.class))),
            @ApiResponse(responseCode = "401", description = "Invalid credentials",
                    content = @Content)
    })
    public ResponseEntity<?> signIn(@RequestBody AuthRequestDTO authenticationDTO) {
        try {
            UsernamePasswordAuthenticationToken authenticationToken =
                    new UsernamePasswordAuthenticationToken(
                            authenticationDTO.getUsername(), authenticationDTO.getPassword());

            Authentication authentication = authenticationManager.authenticate(authenticationToken);

            SecurityContextHolder.getContext().setAuthentication(authentication);

            User user = (User) authentication.getPrincipal();
            PersonResponseDTO person = personService.findByUsername(user.getUsername());

            String jwt = generateJwtToken(user.getUsername());

            JwtResponseDTO response = new JwtResponseDTO(jwt,
                    person.getId(),
                    person.getUsername());

            return ResponseEntity.ok(response);
        } catch (BadCredentialsException exception) {
            throw new ResponseStatusException(HttpStatus.UNAUTHORIZED, "Invalid credentials");
        }

    }

}
