package no.bitraf.overlord.rest.member;

import com.fasterxml.jackson.annotation.JsonProperty;

public final class MemberInfo
{
    @JsonProperty
    protected int id;

    @JsonProperty
    protected String fullName;

    public int getId()
    {
        return id;
    }

    public String getFullName()
    {
        return fullName;
    }
}
