package no.bitraf.overlord.rest.checkin;

import no.bitraf.overlord.db.CheckinEntity;
import no.bitraf.overlord.db.CheckinRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;


@RestController
@RequestMapping("/checkins")
public class CheckinResource
{
    @Autowired
    public CheckinRepository repository;

    @RequestMapping(method = RequestMethod.GET)
    public Iterable<CheckinEntity> getAll() { return  this.repository.findAll(); }
}
