package no.bitraf.overlord.db;


import org.springframework.data.repository.PagingAndSortingRepository;

public interface EventRepository
        extends PagingAndSortingRepository<EventEntity, Integer> {
}
