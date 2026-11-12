package org.cbench.handlerspring.app;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.Map;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;

@SpringBootApplication
public class Main {
    public static void main(String[] args) throws Exception {
        boolean serve = args.length > 0 && args[0].equals("serve");
        SpringApplication app = new SpringApplication(Main.class);
        app.setDefaultProperties(Map.of(
                "spring.main.banner-mode", "off",
                "logging.level.root", "warn",
                "server.address", "127.0.0.1",
                "server.port", serve ? "8080" : "0"));
        ConfigurableApplicationContext context = app.run(args);
        if (serve) {
            return;
        }
        // Without "serve", answer one request on an ephemeral port and exit.
        String port = context.getEnvironment().getProperty("local.server.port");
        HttpResponse<String> response = HttpClient.newHttpClient().send(
                HttpRequest.newBuilder(URI.create("http://127.0.0.1:" + port + "/fingerprint?data=hello")).build(),
                HttpResponse.BodyHandlers.ofString());
        System.out.println(response.statusCode() + " " + response.body());
        System.exit(SpringApplication.exit(context));
    }
}
