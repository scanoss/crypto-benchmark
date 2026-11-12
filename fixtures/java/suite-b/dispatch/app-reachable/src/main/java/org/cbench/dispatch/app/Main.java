package org.cbench.dispatch.app;

import java.nio.charset.StandardCharsets;
import java.util.HexFormat;
import org.cbench.dispatch.liba.Hasher;
import org.cbench.dispatch.liba.Hashers;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) throws Exception {
        Hasher h = Hashers.standard();
        System.out.println(HexFormat.of().formatHex(h.sum("hello".getBytes(StandardCharsets.UTF_8))));
    }
}
