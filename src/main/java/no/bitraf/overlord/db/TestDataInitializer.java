package no.bitraf.overlord.db;

import javax.annotation.PostConstruct;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public final class TestDataInitializer
{
    @Autowired
    protected MemberRepository repository;

    @PostConstruct
    public void initialize()
    {
        addMember( 1, "Erik Sunde", "esu@enonic.com" );
        addMember( 2, "Sten Roger Sandvik", "srs@enonic.com" );
    }

    private void addMember( final int id, final String firstName, final String email )
    {
        final MemberEntity entity = new MemberEntity();
        entity.id = id;
        entity.fullName = firstName;
        entity.email = email;

        this.repository.save( entity );
    }
}
