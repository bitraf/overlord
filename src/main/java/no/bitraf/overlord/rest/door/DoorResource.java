package no.bitraf.overlord.rest.door;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public final class DoorResource
{
    @RequestMapping("/door")
    public DoorStatus status()
    {
        final DoorStatus status = new DoorStatus();
        status.setLocked( false );
        return status;
    }
}
