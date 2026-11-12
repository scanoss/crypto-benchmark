package org.cbench.wrapper.app;

import java.nio.charset.StandardCharsets;
import java.util.HexFormat;
import org.cbench.wrapper.liba.Protector;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) throws Exception {
        System.out.println(HexFormat.of().formatHex(run("hello".getBytes(StandardCharsets.UTF_8))));
    }

    static byte[] run(byte[] data) throws Exception {
        return Protector.protect(data);
    }
}
