package no.bitraf.overlord.rest.member;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import com.google.common.collect.Lists;

import no.bitraf.overlord.db.MemberEntity;
import no.bitraf.overlord.db.MemberRepository;

@RestController
@RequestMapping("/member")
public final class MemberResource
{
    @Autowired
    protected MemberRepository repository;

    @RequestMapping(method = RequestMethod.GET)
    public List<MemberInfo> getAll()
    {
        final List<MemberInfo> result = Lists.newArrayList();
        for ( final MemberEntity entity : this.repository.findAll() )
        {
            result.add( toMemberInfo( entity ) );
        }

        return result;
    }

    @RequestMapping(value = "{id}", method = RequestMethod.GET)
    public MemberDetails getDetails( @PathVariable(value = "id") final int id )
    {
        final MemberEntity entity = this.repository.findOne( id );
        if ( entity != null )
        {
            return toMemberDetails( entity );
        }

        throw new NotFoundException( "Member " + id + " not found" );
    }

    private MemberInfo toMemberInfo( final MemberEntity entity )
    {
        final MemberInfo result = new MemberInfo();
        result.id = entity.getId();
        result.fullName = entity.getFullName();
        return result;
    }

    private MemberDetails toMemberDetails( final MemberEntity entity )
    {
        final MemberDetails result = new MemberDetails();
        result.id = entity.getId();
        result.raw = entity;
        return result;
    }
}
