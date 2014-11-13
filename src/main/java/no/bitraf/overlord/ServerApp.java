package no.bitraf.overlord;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;

@Configuration
@EnableAutoConfiguration
@ComponentScan
public class ServerApp
{
    public static void main( final String... args )
        throws Exception
    {
        final SpringApplication app = new SpringApplication( ServerApp.class );
        app.setHeadless( true );
        app.setShowBanner( true );
        app.run( args );
    }
}
