package byson.api.server.core;

import javax.xml.bind.annotation.XmlElement;


public interface ApiError {
    @XmlElement(name = "message")
    String getMessage();
}
