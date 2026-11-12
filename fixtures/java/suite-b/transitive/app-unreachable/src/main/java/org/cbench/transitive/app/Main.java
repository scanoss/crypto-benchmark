package org.cbench.transitive.app;

import java.nio.charset.StandardCharsets;
import java.util.HexFormat;
import org.cbench.transitive.liba.Fingerprint;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) throws Exception {
        System.out.println("hello");
    }

    static byte[] run(byte[] data) throws Exception {
        return Fingerprint.of(data);
    }
}
