package no.bitraf.overlord.db;

import org.springframework.data.repository.PagingAndSortingRepository;


public interface InvoiceRepository
        extends PagingAndSortingRepository<InvoiceEntity, Integer> {
}
