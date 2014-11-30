package no.bitraf.overlord.db;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import java.util.Date;

@Entity
@Table(name = "transactions")
public class TransactionEntity {

    @Id
    @Column(name = "id")
    protected int id;

    @Column(name = "date")
    protected Date date;

    @Column(name = "reason")
    protected String reason;

    public int getId() {
        return id;
    }

    public Date getDate() {
        return date;
    }

    public String getReason() {
        return reason;
    }
}
