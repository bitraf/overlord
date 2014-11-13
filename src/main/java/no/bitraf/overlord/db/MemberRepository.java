package no.bitraf.overlord.db;

import org.springframework.data.repository.PagingAndSortingRepository;

public interface MemberRepository
    extends PagingAndSortingRepository<MemberEntity, Integer>
{
}
