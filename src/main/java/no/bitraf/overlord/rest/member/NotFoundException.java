package no.bitraf.overlord.rest.member;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(HttpStatus.NOT_FOUND)
public final class NotFoundException
    extends RuntimeException
{
    public NotFoundException( final String message )
    {
        super( message );
    }
}

