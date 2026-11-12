package org.cbench.depth3.app;

import java.nio.charset.StandardCharsets;
import java.util.HexFormat;
import org.cbench.depth3.liba.Sealer;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) throws Exception {
        System.out.println("hello");
    }

    static byte[] run(byte[] data) throws Exception {
        return Sealer.seal(data);
    }
}
