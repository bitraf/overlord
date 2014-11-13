package no.bitraf.overlord.rest.member;

import com.fasterxml.jackson.annotation.JsonProperty;

import no.bitraf.overlord.db.MemberEntity;

public final class MemberDetails
{
    @JsonProperty
    protected int id;

    @JsonProperty
    protected MemberEntity raw;

    /**
     *
     * account.type = ['user' or 'product']
     *
     *
     *
     *
     */
}
