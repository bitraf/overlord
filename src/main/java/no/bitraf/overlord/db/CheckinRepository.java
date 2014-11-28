package no.bitraf.overlord.db;

import org.springframework.data.repository.PagingAndSortingRepository;


public interface CheckinRepository
        extends PagingAndSortingRepository<CheckinEntity, Integer>
{

}
