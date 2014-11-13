package no.bitraf.overlord.db;

import java.util.Date;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table(name = "members")
public final class MemberEntity
{
    @Id
    @Column(name = "id")
    protected int id;

    @Column(name = "timestamp")
    protected Date timestamp;

    @Column(name = "full_name")
    protected String fullName;

    @Column(name = "email")
    protected String email;

    @Column(name = "type")
    protected String type;

    @Column(name = "account")
    protected int account;

    @Column(name = "organization")
    protected String organization;

    @Column(name = "price")
    protected int price;

    @Column(name = "recurrence")
    protected String recurrence;

    @Column(name = "flag")
    protected String flag;

    public int getId()
    {
        return id;
    }

    public Date getTimestamp()
    {
        return timestamp;
    }

    public String getFullName()
    {
        return fullName;
    }

    public String getEmail()
    {
        return email;
    }

    public String getType()
    {
        return type;
    }

    public int getAccount()
    {
        return account;
    }

    public String getOrganization()
    {
        return organization;
    }

    public int getPrice()
    {
        return price;
    }

    public String getRecurrence()
    {
        return recurrence;
    }

    public String getFlag()
    {
        return flag;
    }
}

