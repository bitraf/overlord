package no.bitraf.overlord.db;

import org.springframework.data.repository.PagingAndSortingRepository;

public interface MembershipInfoRepository
        extends PagingAndSortingRepository<MembershipInfoEntity, Integer> {
}
