package no.bitraf.overlord.db;


import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import java.util.Date;

@Entity
@Table(name = "fiken_faktura")
public class FikenFakturaEntity {

    @Id
    @Column(name = "id")
    protected int id;

    @Column(name = "last_date_billed")
    protected Date lastDateBilled;

    @Column(name = "fiken_kundenummer")
    protected int fikenKundenummer;

    @Column(name = "account")
    protected int account;

    public int getId() {
        return id;
    }

    public Date getLastDateBilled() {
        return lastDateBilled;
    }

    public int getFikenKundenummer() {
        return fikenKundenummer;
    }

    public int getAccount() {
        return account;
    }
}
