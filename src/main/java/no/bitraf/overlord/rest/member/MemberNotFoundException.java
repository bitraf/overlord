package no.bitraf.overlord.rest.member;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(HttpStatus.NOT_FOUND)
public final class MemberNotFoundException
    extends RuntimeException
{
    public MemberNotFoundException( final int id )
    {
        super( "Member " + id + " not found" );
    }
}
