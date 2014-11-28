package no.bitraf.overlord.db;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "auth_log")
public class AuthLogEntity {

    @Column(name = "host")
    protected String host;

    @Column(name = "account")
    protected Integer account; // Set as Integer instead of int because value might be null

    @Id
    @Column(name = "date")
    protected Date date;

    @Column(name = "realm")
    protected String realm;

    public String getHost() {
        return host;
    }

    public Integer getAccount() {
        return account;
    }

    public Date getDate() {
        return date;
    }

    public String getRealm() {
        return realm;
    }
}
