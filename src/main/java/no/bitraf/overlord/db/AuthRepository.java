package no.bitraf.overlord.db;


import org.springframework.data.repository.PagingAndSortingRepository;

public interface AuthRepository
        extends PagingAndSortingRepository<AuthEntity, Integer>
{
}
