package no.bitraf.overlord.db;


import org.springframework.data.repository.PagingAndSortingRepository;

public interface TransactionLineRepository
        extends PagingAndSortingRepository<TransactionLineEntity, Integer> {
}
