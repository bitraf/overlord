package no.bitraf.overlord;

import io.dropwizard.Application;
import io.dropwizard.setup.Bootstrap;
import io.dropwizard.setup.Environment;

public final class ServerApp
    extends Application<ServerConfig>
{
    @Override
    public void initialize( final Bootstrap<ServerConfig> bootstrap )
    {
    }

    @Override
    public void run( final ServerConfig configuration, final Environment environment )
        throws Exception
    {
    }

    public static void main( final String... args )
        throws Exception
    {
        new ServerApp().run( args );
    }
}
