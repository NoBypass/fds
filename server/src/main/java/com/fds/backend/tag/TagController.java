package com.fds.backend.tag;;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping(TagController.PATH)
public class TagController {
    public static final String PATH = "/tag";

    TagService tagService;

    @Autowired
    public TagController(TagService tagService) {
        this.tagService = tagService;
    }

    @DeleteMapping("{id}")
    public void deleteById(@PathVariable Integer id) {
        tagService.deleteById(id);
    }

    @GetMapping("{id}")
    public ResponseEntity<?> findById(@PathVariable Integer id) {
        return ResponseEntity.ok(tagService.findById(id));
    }

    @PostMapping
    public ResponseEntity<?> save(@RequestBody Tag tag) {
        return ResponseEntity.ok(tagService.save(tag));
    }

    @PatchMapping("{id}")
    public ResponseEntity<?> update(@PathVariable Integer id, @RequestBody TagRequestDTO tag) {
        return ResponseEntity.ok(tagService.update(tag, id));
    }

    @GetMapping
    public ResponseEntity<?> findTags(@RequestParam(value = "name", required = false) String name) {
        return ResponseEntity.ok(tagService.findTags(name));
    }
}
