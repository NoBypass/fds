package com.fds.backend.tag;

import com.fds.backend.discordUser.DiscordUser;

import java.util.ArrayList;
import java.util.List;

public class TagMapper {
    public static TagResponseDTO toResponseDTO(Tag tag) {
        TagResponseDTO tagResponseDTO = new TagResponseDTO();
        tagResponseDTO.setId(tag.getId());
        List<Integer> itemIds = new ArrayList<Integer>();
        for (DiscordUser discordUser : tag.getLinkedItems()) itemIds.add(discordUser.getId());
        tagResponseDTO.setItemIds(itemIds);
        return tagResponseDTO;
    }

    public static Tag fromRequestDTO(TagRequestDTO tagRequestDTO) {
        Tag tag = new Tag();
        tag.setName(tag.getName());
        return tag;
    }
}