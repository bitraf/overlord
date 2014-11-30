package no.bitraf.overlord.db;


import org.springframework.data.repository.PagingAndSortingRepository;

public interface MacRepository
        extends PagingAndSortingRepository<MacEntity, Integer> {
}
