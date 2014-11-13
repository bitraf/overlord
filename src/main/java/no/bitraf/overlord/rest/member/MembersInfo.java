package no.bitraf.overlord.rest.member;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonProperty;

public final class MembersInfo
{
    @JsonProperty
    protected int totalPages;

    @JsonProperty
    protected int totalElements;

    @JsonProperty
    protected List<MemberInfo> members;
}
