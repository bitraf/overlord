package no.bitraf.overlord.db;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import java.math.BigDecimal;
import java.util.Date;

@Entity
@Table(name = "events")
public final class EventEntity {

    @Id
    @Column(name = "id")
    protected int id;

    @Column(name = "date")
    protected Date date;

    @Column(name = "type")
    protected String type;

    @Column(name = "account")
    protected int account;

    @Column(name = "amount")
    protected BigDecimal amount;

    public int getId() {
        return id;
    }

    public Date getDate() {
        return date;
    }

    public String getType() {
        return type;
    }

    public int getAccount() {
        return account;
    }

    public BigDecimal getAmount() {
        return amount;
    }
}
