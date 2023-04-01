package com.fds.backend.discordUser;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping(DiscordUserController.PATH)
public class DiscordUserController {
    public static final String PATH = "/discord/user";

    private final DiscordUserService discordUserService;

    @Autowired
    public DiscordUserController(DiscordUserService discordUserService) {
        this.discordUserService = discordUserService;
    }

    @GetMapping("{id}")
    public ResponseEntity<?> findById(@PathVariable Integer id) {
        return ResponseEntity.ok(DiscordUserMapper.fromDTO(discordUserService.findById(id)));
    }

    @PostMapping
    public void save(@RequestBody DiscordUser discordUser) {
        discordUserService.insert(DiscordUserMapper.toDTO(discordUser));
    }

    @DeleteMapping("{id}")
    public void deleteById(@PathVariable Integer id) {
        discordUserService.deleteById(id);
    }

    @GetMapping
    public ResponseEntity<?> findItems(String name) {
        return ResponseEntity.ok(discordUserService.findItems(name));
    }
}


