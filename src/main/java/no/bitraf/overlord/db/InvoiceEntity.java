package no.bitraf.overlord.db;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table(name = "invoices")
public class InvoiceEntity {

    @Id
    @Column(name = "id")
    protected int id;

    @Column(name = "text")
    protected String text;
}
