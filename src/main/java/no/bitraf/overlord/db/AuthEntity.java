package no.bitraf.overlord.db;

import javax.persistence.*;

@Entity
@Table(name = "auth")
public class AuthEntity {

    @Id
    @Column(name = "account")
    private int account;

    @Column(name = "realm")
    private String realm;

    @Column(name = "data")
    private String data;

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
