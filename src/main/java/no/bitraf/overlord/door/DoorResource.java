package no.bitraf.overlord.door;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

@Path("door")
@Produces(MediaType.APPLICATION_JSON)
public final class DoorResource
{
    @GET
    public DoorStatus status()
    {
        return new DoorStatus();
    }
}
