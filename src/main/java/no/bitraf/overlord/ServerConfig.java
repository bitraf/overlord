package no.bitraf.overlord;

import com.fasterxml.jackson.annotation.JsonProperty;

import io.dropwizard.Configuration;

public final class ServerConfig
    extends Configuration
{
    @JsonProperty
    private int port;

    public int getPort()
    {
        return port;
    }

    public void setPort( final int port )
    {
        this.port = port;
    }
}
