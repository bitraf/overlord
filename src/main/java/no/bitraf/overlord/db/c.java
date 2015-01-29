package no.bitraf.overlord.db;


import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import java.util.Date;

@Entity
@Table(name = "checkins")
public class CheckinEntity {

    @Id
    @Column(name = "id")
    protected int id;

    @Column(name = "account")
    protected int account;

    @Column(name = "date")
    protected Date date;

    @Column(name = "type")
    protected String type;

    public int getId() {
        return id;
    }

    public int getAccount() {
        return account;
    }

    public Date getDate() {
        return date;
    }

    public String getType() {
        return type;
    }
}
