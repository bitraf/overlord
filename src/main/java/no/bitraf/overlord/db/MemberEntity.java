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

    @Column(name = "full_name")
    protected String fullName;

    @Column(name = "email")
    protected String email;

    @Column(name = "date")
    protected Date date;

    @Column(name = "account")
    protected int account;

    @Column(name = "organization")
    protected String organization;

    @Column(name = "price")
    protected int price;

    @Column(name = "flag")
    protected String flag;

    @Column(name = "recurrence")
    protected String recurrence;

    public int getId()
    {
        return id;
    }

    public String getFullName()
    {
        return fullName;
    }

    public String getEmail()
    {
        return email;
    }

    public Date getDate()
    {
        return date;
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

    public String getFlag()
    {
        return flag;
    }

    public String getRecurrence()
    {
        return recurrence;
    }
}

