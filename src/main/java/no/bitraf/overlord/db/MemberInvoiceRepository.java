package no.bitraf.overlord.db;


import org.springframework.data.repository.PagingAndSortingRepository;

public interface MemberInvoiceRepository
        extends PagingAndSortingRepository<MemberInvoiceEntity, Integer> {
}
