package byson.api.core;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;


public class Utils {
    public static String intersperseString(Iterable objs, String separator) {
        StringBuilder s = new StringBuilder();
        boolean first = true;
        for (Object obj : objs) {
            if (first) {
                s.append(obj.toString());
                first = false;
            } else {
                s.append(separator);
                s.append(obj.toString());
            }
        }
        return s.toString();
    }


    public static String generateCommaSeparatedString(Iterable objs) {
        return intersperseString(objs, ",");
    }


    /**
     * A very simple and naive - but null-safe - function to parse comma separated lists.
     */
    public static List<Integer> parseIntList(String commaSeparatedList) {
        return parseStringList(commaSeparatedList).stream()
                .map(Integer::parseInt)
                .collect(Collectors.toList());
    }


    /**
     * A very simple and naive - but null-safe - function to parse comma separated lists.
     */
    public static List<Long> parseLongList(String commaSeparatedList) {
        return parseStringList(commaSeparatedList).stream()
                .map(Long::parseLong)
                .collect(Collectors.toList());
    }


    /**
     * A very simple and naive - but null-safe - function to parse comma separated lists.
     */
    public static List<String> parseStringList(String commaSeparatedList) {
        if (commaSeparatedList != null) {
            return Arrays.asList(commaSeparatedList.split(","));
        }
        return new ArrayList<>();
    }
}
