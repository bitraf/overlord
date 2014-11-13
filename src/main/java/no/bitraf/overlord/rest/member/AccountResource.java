package no.bitraf.overlord.rest.member;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import no.bitraf.overlord.db.AccountEntity;
import no.bitraf.overlord.db.AccountRepository;

@RestController
@RequestMapping("/account")
public final class AccountResource
{
    @Autowired
    protected AccountRepository repository;

    @RequestMapping(method = RequestMethod.GET)
    public Iterable<AccountEntity> getAll()
    {
        return this.repository.findAll();
    }

    @RequestMapping(value = "{id}", method = RequestMethod.GET)
    public AccountEntity getDetails( @PathVariable(value = "id") final int id )
    {
        final AccountEntity entity = this.repository.findOne( id );
        if ( entity != null )
        {
            return entity;
        }

        throw new NotFoundException( "Account " + id + " not found" );
    }
}
