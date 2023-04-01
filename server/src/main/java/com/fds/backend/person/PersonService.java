package com.fds.backend.person;

import com.fds.backend.security.AuthRequestDTO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import javax.persistence.EntityNotFoundException;
import java.util.List;

@Service
public class PersonService {
    private final PersonRepository personRepository;
    private final PasswordEncoder passwordEncoder;

    @Autowired
    public PersonService(PersonRepository personRepository, PasswordEncoder passwordEncoder) {
        this.personRepository = personRepository;
        this.passwordEncoder = passwordEncoder;
    }

    public PersonResponseDTO create(AuthRequestDTO authRequestDTO) {
        Person person = new Person();
        person.setPassword(passwordEncoder.encode(authRequestDTO.getPassword()));
        person.setUsername(authRequestDTO.getUsername());
        return PersonMapper.toResponseDTO(personRepository.save(person));
    }

    public List<PersonResponseDTO> findAll() {
        return personRepository.findAll().stream().map(PersonMapper::toResponseDTO).toList();
    }

    public PersonResponseDTO findById(Integer id) {
        return PersonMapper.toResponseDTO(personRepository.findById(id).orElseThrow(EntityNotFoundException::new));
    }

    public PersonResponseDTO findByUsername(String username) {
        return PersonMapper.toResponseDTO(personRepository.findByUsername(username));
    }

    public PersonResponseDTO update(PersonRequestDTO personRequestDTO, Integer id) {
        Person existing = personRepository.findById(id).orElseThrow(EntityNotFoundException::new);
        mergePersons(existing, PersonMapper.fromRequestDTO(personRequestDTO));
        return PersonMapper.toResponseDTO(personRepository.save(existing));
    }

    public void deleteById(Integer id) {
        personRepository.deleteById(id);
    }

    private void mergePersons(Person existing, Person changing) {
        if (changing.getUsername() != null) {
            existing.setUsername(changing.getUsername());
        }
        if (changing.getPassword() != null) {
            String newPassword = passwordEncoder.encode(changing.getPassword());
            existing.setPassword(newPassword);
        }
    }
}