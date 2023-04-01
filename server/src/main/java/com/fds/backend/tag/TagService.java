package com.fds.backend.tag;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import javax.persistence.EntityNotFoundException;
import java.util.List;
import java.util.stream.Collectors;

@Service
class TagService {
    @Autowired
    public TagService(TagRepository tagRepository) {
        this.tagRepository = tagRepository;
    }

    private final TagRepository tagRepository;

    public TagResponseDTO findById(Integer id) {
        return TagMapper.toResponseDTO(tagRepository.findById(id).orElseThrow(EntityNotFoundException::new));
    }

    public TagResponseDTO insert(TagRequestDTO tagRequestDTO) {
        return TagMapper.toResponseDTO(tagRepository.save(TagMapper.fromRequestDTO(tagRequestDTO)));
    }

    public void deleteById(Integer id) {
        tagRepository.deleteById(id);
    }

    public TagResponseDTO update(TagRequestDTO tagRequestDTO, Integer tagId) {
        Tag existingTag = tagRepository.findById(tagId).orElseThrow(EntityNotFoundException::new);
        Tag changingTag = TagMapper.fromRequestDTO(tagRequestDTO);
        mergeTags(existingTag, changingTag);
        return TagMapper.toResponseDTO(tagRepository.save(existingTag));
    }

    private void mergeTags(Tag exisitingTag, Tag changingTag) {
        if (changingTag.getId() != null) {
            exisitingTag.setId(changingTag.getId());
        }
    }

    public List<TagResponseDTO> findTags(String name) {
        List<Tag> tags;
        if (name == null) {
            tags = tagRepository.findAll();
        } else {
            tags = tagRepository.findByName(name);
        }
        return tags.stream().map(TagMapper::toResponseDTO).collect(Collectors.toList());
    }

    public Object save(Tag tag) {
        return tagRepository.save(tag);
    }
}