package no.bitraf.overlord.db;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import java.math.BigDecimal;
import java.util.Date;

@Entity
@Table(name = "billing_runs")
public class BillingRunEntity {

    @Id
    @Column(name = "date")
    protected Date date;

    @Column(name = "account")
    protected int account;

    @Column(name = "amount")
    protected BigDecimal amount;

    public Date getDate() {
        return date;
    }

    public int getAccount() {
        return account;
    }

    public BigDecimal getAmount() {
        return amount;
    }
}
