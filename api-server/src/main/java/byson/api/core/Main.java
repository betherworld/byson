package byson.api.core;

import byson.api.server.core.ApiServer;

import java.time.LocalDateTime;


public class Main {
    public static void main(String[] args) {
        System.out.println("Started at: " + LocalDateTime.now());
        new ApiServer(args).start();
    }
}
