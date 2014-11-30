package no.bitraf.overlord.db;


import org.springframework.data.repository.PagingAndSortingRepository;

public interface TransactionRepository
        extends PagingAndSortingRepository<TransactionEntity, Integer> {
}
