package no.bitraf.overlord.rest.door;

import com.fasterxml.jackson.annotation.JsonProperty;

public final class DoorStatus
{
    @JsonProperty
    private boolean locked;

    public boolean isLocked()
    {
        return locked;
    }

    public void setLocked( final boolean locked )
    {
        this.locked = locked;
    }
}
