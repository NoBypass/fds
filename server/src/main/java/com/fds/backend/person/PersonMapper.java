package com.fds.backend.person;

import java.util.List;

import com.fds.backend.discordUser.DiscordUser;

public class PersonMapper {
    public static PersonResponseDTO toResponseDTO(Person person) {
        PersonResponseDTO personResponseDTO = new PersonResponseDTO();
        if (person.getItems() != null) {
            List<Integer> itemIds = person
                    .getItems()
                    .stream()
                    .map(DiscordUser::getId)
                    .toList();

            personResponseDTO.setItemIds(itemIds);
        }

        personResponseDTO.setId(person.getId());
        personResponseDTO.setUsername(person.getUsername());

        return personResponseDTO;
    }

    public static Person fromRequestDTO(PersonRequestDTO personRequestDTO) {
        Person person = new Person();

        person.setUsername(personRequestDTO.getUsername());

        return person;
    }

}
