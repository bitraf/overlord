package no.bitraf.overlord.rest.member;

import com.fasterxml.jackson.annotation.JsonProperty;

public final class MemberDetails
{
    @JsonProperty
    protected int id;

    @JsonProperty
    protected String fullName;

    @JsonProperty
    protected String email;

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
}
