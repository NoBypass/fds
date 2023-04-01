package com.fds.backend.tag;

import java.util.List;
import java.util.Objects;

public class TagResponseDTO {
    private Integer id;
    private List<Integer> itemIds;

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (!(o instanceof TagResponseDTO that)) return false;
        return Objects.equals(id, that.id) && Objects.equals(itemIds, that.itemIds);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id, itemIds);
    }

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public List<Integer> getItemIds() {
        return itemIds;
    }

    public void setItemIds(List<Integer> itemIds) {
        this.itemIds = itemIds;
    }
}
