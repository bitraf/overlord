package no.bitraf.overlord.rest.authLog;

import no.bitraf.overlord.db.AuthLogEntity;
import no.bitraf.overlord.db.AuthLogRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/authlog")
public class AuthLogResource {
    @Autowired
    public AuthLogRepository repository;

    @RequestMapping(method = RequestMethod.GET)
    public Iterable<AuthLogEntity> getAll() { return this.repository.findAll(); }
}
