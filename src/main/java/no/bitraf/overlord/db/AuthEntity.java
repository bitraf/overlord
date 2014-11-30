package no.bitraf.overlord.db;

import javax.persistence.*;

@Entity
@Table(name = "auth")
public class AuthEntity {

    @Id
    @Column(name = "account")
    protected int account;

    @Column(name = "realm")
    protected String realm;

    @Column(name = "data")
    protected String data;

    public int getAccount() {
        return account;
    }

    public String getRealm() {
        return realm;
    }

    public String getData() {
        return data;
    }
}
