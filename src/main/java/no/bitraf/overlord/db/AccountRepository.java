package no.bitraf.overlord.db;

import org.springframework.data.repository.PagingAndSortingRepository;

public interface AccountRepository
    extends PagingAndSortingRepository<AccountEntity, Integer>
{
}
