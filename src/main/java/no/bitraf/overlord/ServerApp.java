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
        SpringApplication.run( ServerApp.class, args );
    }
}
