package com.fds.backend.discordUser;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import javax.persistence.EntityNotFoundException;
import java.util.List;

@Service
public class DiscordUserService {
    @Autowired
    public DiscordUserService(DiscordUserRepository itemRepository) {
        this.itemRepository = itemRepository;
    }

    private final DiscordUserRepository itemRepository;

    public DiscordUserResponseDTO findById(Integer id) {
        return DiscordUserMapper.toDTO(itemRepository.findById(id).orElseThrow(EntityNotFoundException::new));
    }

    public DiscordUserResponseDTO insert(DiscordUserRequestDTO discordUserRequestDTO) {
        return DiscordUserMapper.toDTO(itemRepository.save(DiscordUserMapper.fromDTO(discordUserRequestDTO)));
    }

    public DiscordUserResponseDTO update(DiscordUserRequestDTO discordUserRequestDTO, Integer itemId) {
        DiscordUser existingDiscordUser = itemRepository.findById(itemId).orElseThrow(EntityNotFoundException::new);
        DiscordUser changingDiscordUser = DiscordUserMapper.fromDTO(discordUserRequestDTO);
        return DiscordUserMapper.toDTO(itemRepository.save(mergeItems(existingDiscordUser, changingDiscordUser)));
    }

    private DiscordUser mergeItems(DiscordUser existingDiscordUser, DiscordUser changingDiscordUser) {
        if (changingDiscordUser.getDescription() != null) existingDiscordUser.setDescription(changingDiscordUser.getDescription());
        if (changingDiscordUser.getName() != null) existingDiscordUser.setName(changingDiscordUser.getName());
        if (changingDiscordUser.getPerson() != null) existingDiscordUser.setPerson(changingDiscordUser.getPerson());
        if (changingDiscordUser.getDoneAt() != null) existingDiscordUser.setDoneAt(changingDiscordUser.getDoneAt());
        return existingDiscordUser;
    }

    public void deleteById(Integer id) {
        itemRepository.deleteById(id);
    }

    public List<DiscordUser> findByName(String name) {
        return itemRepository.findByName(name);
    }
}