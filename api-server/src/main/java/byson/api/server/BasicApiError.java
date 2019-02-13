package byson.api.server;

import byson.api.server.core.ApiError;


public class BasicApiError implements ApiError {
    private String message;
    private String className;
    private String stackTrace;


    public BasicApiError(String message, String className, String stackTrace) {
        this.message = message;
        this.className = className;
        this.stackTrace = stackTrace;
    }


    @Override
    public String getMessage() {
        return message;
    }


    public String getClassName() {
        return className;
    }


    public String getStackTrace() {
        return stackTrace;
    }
}
