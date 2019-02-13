package byson.api.server.core;

import org.eclipse.jetty.server.Connector;
import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.server.ServerConnector;
import org.eclipse.jetty.server.handler.HandlerList;
import org.eclipse.jetty.servlet.FilterHolder;
import org.eclipse.jetty.servlet.ServletContextHandler;
import org.eclipse.jetty.servlet.ServletHolder;
import org.eclipse.jetty.servlets.CrossOriginFilter;
import org.glassfish.jersey.servlet.ServletContainer;

import javax.servlet.DispatcherType;
import java.util.Arrays;
import java.util.EnumSet;
import java.util.List;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;


public class ApiServer {
    public static final ExecutorService threadPool = Executors.newCachedThreadPool();


    private boolean serveSwagger = true;


    public ApiServer(String[] args) {
        if (args.length > 0 && args[0].equals("help")) {
            printHelp();
        }
        List<String> argList = Arrays.asList(args);
        if (argList.contains("--no-swagger")) {
            this.serveSwagger = false;
        }
    }


    private void printHelp() {
        System.out.println("Usage: ... api_server <options>");
        System.out.println("Available options: --no-swagger");
    }


    public void start() {
        int port = 10050;
        Server server = new Server();

        ServerConnector connector = new ServerConnector(server);
        connector.setPort(port);
        connector.setHost("localhost");
        server.setConnectors(new Connector[]{connector});

        HandlerList handlers = new HandlerList();
        server.setHandler(handlers);

        ServletContextHandler context = new ServletContextHandler(ServletContextHandler.NO_SESSIONS);
        context.setContextPath("/");
        handlers.addHandler(context);

        // Actual API
        String providerPackages = "byson.api.server.resource,byson.api.server.devresource";

        ServletHolder apiServletHolder = context.addServlet(ServletContainer.class, "/*");
        apiServletHolder.setInitOrder(1);
        apiServletHolder.setInitParameter(
                "jersey.config.server.provider.packages",
                providerPackages
        );

        FilterHolder cors = context.addFilter(CrossOriginFilter.class, "/*", EnumSet.of(DispatcherType.REQUEST));
        cors.setInitParameter(CrossOriginFilter.ALLOWED_ORIGINS_PARAM, "*");
        cors.setInitParameter(CrossOriginFilter.ACCESS_CONTROL_ALLOW_ORIGIN_HEADER, "*");
        cors.setInitParameter(CrossOriginFilter.ALLOWED_METHODS_PARAM, "GET,POST,HEAD");
        cors.setInitParameter(CrossOriginFilter.ALLOWED_HEADERS_PARAM,
                "X-Requested-With,Content-Type,Accept,Origin");

        try {
            server.start();
            System.out.println("API Server started at port " + port + ".");
            server.join();
        } catch (Exception ex) {
            System.err.println("The API Server could not be started.");
            ex.printStackTrace();
            System.exit(1);
        }
    }
}
