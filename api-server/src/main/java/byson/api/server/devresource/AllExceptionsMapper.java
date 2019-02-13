package byson.api.server.devresource;

import byson.api.server.BasicApiError;
import org.apache.commons.lang3.exception.ExceptionUtils;

import javax.ws.rs.WebApplicationException;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import javax.ws.rs.ext.ExceptionMapper;
import javax.ws.rs.ext.Provider;
import java.util.HashMap;
import java.util.Map;


@Provider
public class AllExceptionsMapper implements
        ExceptionMapper<Exception> {
    @Override
    public Response toResponse(Exception ex) {
        if (ex instanceof WebApplicationException) {
            return ((WebApplicationException) ex).getResponse();
        }
        Map<String, Object> responseContent = new HashMap<>();
        responseContent.put("error", new BasicApiError(
                ex.getMessage(),
                ex.getClass().getName(),
                ExceptionUtils.getStackTrace(ex)
        ));
        return Response
                .status(500)
                .entity(responseContent)
                .type(MediaType.APPLICATION_JSON)
                .build();
    }
}
