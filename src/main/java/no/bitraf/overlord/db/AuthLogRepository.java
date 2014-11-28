package no.bitraf.overlord.db;

import org.springframework.data.repository.PagingAndSortingRepository;

public interface AuthLogRepository
        extends PagingAndSortingRepository<AuthLogEntity, Integer> {
}
