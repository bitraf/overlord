package no.bitraf.overlord.db;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import java.math.BigDecimal;
import java.util.Date;

@Entity
@Table(name = "member_invoices")
public class MemberInvoiceEntity {

    @Id
    @Column(name = "id")
    protected int id;

    @Column(name = "date")
    protected Date date;

    @Column(name = "pay_by")
    protected Date payBy;

    @Column(name = "amount")
    protected BigDecimal amount;

    @Column(name = "bilag")
    protected int bilag;

    public int getId() {
        return id;
    }

    public Date getDate() {
        return date;
    }

    public Date getPayBy() {
        return payBy;
    }

    public BigDecimal getAmount() {
        return amount;
    }

    public int getBilag() {
        return bilag;
    }
}
