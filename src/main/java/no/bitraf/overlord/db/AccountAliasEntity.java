package no.bitraf.overlord.db;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table(name = "account_aliases")
public final class AccountAliasEntity
{
    @Id
    @Column(name = "account")
    protected int account;

    @Column(name = "alias")
    protected String alias;

    public int getAccount()
    {
        return account;
    }

    public String getAlias()
    {
        return alias;
    }
}
