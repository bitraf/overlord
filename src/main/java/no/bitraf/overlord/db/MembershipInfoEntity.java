package no.bitraf.overlord.db;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import java.math.BigDecimal;

@Entity
@Table(name = "membership_infos")
public class MembershipInfoEntity {

    @Id
    @Column(name = "name")
    protected String name;


    @Column(name = "price")
    protected BigDecimal price;

    @Column(name = "recurrence")
    protected String recurrence;

    public String getName() {
        return name;
    }

    public BigDecimal getPrice() {
        return price;
    }

    public String getRecurrence() {
        return recurrence;
    }
}
