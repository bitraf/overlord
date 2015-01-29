package no.bitraf.overlord.rest.door;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

@RestController
public final class DoorResource
{
    @RequestMapping(value = "/v1/door", method = RequestMethod.POST)
    public void open()
    {
    }
}
