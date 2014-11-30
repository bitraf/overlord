package no.bitraf.overlord.db;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import java.math.BigDecimal;

@Entity
@Table(name = "transaction_lines")
public class TransactionLineEntity {

    @Id
    @Column(name = "transaction")
    protected int transaction;

    @Column(name = "debit_account")
    protected int debitAccount;

    @Column(name = "credit_account")
    protected int creditAccount;

    @Column(name = "amount")
    protected BigDecimal amount;

    @Column(name = "stock")
    protected int stock;

    public int getTransaction() {
        return transaction;
    }

    public int getDebitAccount() {
        return debitAccount;
    }

    public int getCreditAccount() {
        return creditAccount;
    }

    public BigDecimal getAmount() {
        return amount;
    }

    public int getStock() {
        return stock;
    }
}
