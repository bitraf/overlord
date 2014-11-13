package no.bitraf.overlord.rest.member;

import com.fasterxml.jackson.annotation.JsonProperty;

import no.bitraf.overlord.db.MemberEntity;

public final class MemberDetails
{
    @JsonProperty
    protected int id;

    // TODO: Should not be exposed directly. Only here to see raw data from database.
    @JsonProperty
    protected MemberEntity raw;
}
