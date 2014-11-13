package no.bitraf.overlord.db;

import java.util.Date;
import java.util.List;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.OneToMany;
import javax.persistence.OneToOne;
import javax.persistence.Table;

@Entity
@Table(name = "accounts")
public final class AccountEntity
{
    @Id
    @Column(name = "id")
    protected int id;

    @Column(name = "name")
    protected String name;

    @Column(name = "type")
    protected String type;

    @Column(name = "last_billed")
    protected Date lastBilled;

    @OneToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "id")
    protected AccountAliasEntity alias;

    @OneToMany(fetch = FetchType.LAZY)
    @JoinColumn(name = "account")
    protected List<MemberEntity> members;

    public int getId()
    {
        return id;
    }

    public String getName()
    {
        return name;
    }

    public String getType()
    {
        return type;
    }

    public Date getLastBilled()
    {
        return lastBilled;
    }

    public AccountAliasEntity getAlias()
    {
        return alias;
    }

    public List<MemberEntity> getMembers()
    {
        return members;
    }
}
