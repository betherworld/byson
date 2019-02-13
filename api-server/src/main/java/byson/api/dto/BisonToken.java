package byson.api.dto;

import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlTransient;
import java.time.LocalDateTime;


public class BisonToken {
    private String uuid;
    private String name;
    private LocalDateTime firstLocatedOnDate;


    public BisonToken() {
    }


    public BisonToken(String uuid, String name, LocalDateTime firstLocatedOnDate) {
        this.uuid = uuid;
        this.name = name;
        this.firstLocatedOnDate = firstLocatedOnDate;
    }


    public String getUuid() {
        return uuid;
    }


    public BisonToken setUuid(String uuid) {
        this.uuid = uuid;
        return this;
    }


    public String getName() {
        return name;
    }


    public BisonToken setName(String name) {
        this.name = name;
        return this;
    }


    @XmlTransient
    public LocalDateTime getFirstLocatedOnDate() {
        return firstLocatedOnDate;
    }


    @XmlElement(name = "firstLocatedOnDate")
    public String getFirstLocatedOnDateString() {
        return firstLocatedOnDate != null ? firstLocatedOnDate.toString() : null;
    }


    public BisonToken setFirstLocatedOnDate(LocalDateTime firstLocatedOnDate) {
        this.firstLocatedOnDate = firstLocatedOnDate;
        return this;
    }
}
