package byson.api.server.resource;

import byson.api.dto.BisonToken;

import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import java.sql.SQLException;
import java.time.LocalDateTime;
import java.util.UUID;


@Path("/bison")
@Produces({MediaType.APPLICATION_JSON})
public class BisonResource {


    @GET
    public BisonToken getArticles(
            @DefaultValue("") @QueryParam("uuid") String uuid
    ) throws SQLException {
        /* Get the objectives that should be computed */
        return new BisonToken(UUID.randomUUID().toString(), "Franku", LocalDateTime.of(2015, 4, 20, 10, 54));
    }
}
