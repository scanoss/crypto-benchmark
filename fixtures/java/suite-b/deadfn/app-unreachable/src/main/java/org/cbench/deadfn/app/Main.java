package org.cbench.deadfn.app;

import java.nio.charset.StandardCharsets;
import java.util.HexFormat;
import org.cbench.deadfn.liba.Api;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) throws Exception {
        System.out.println(HexFormat.of().toHexDigits((byte) Api.checksum("hello".getBytes(StandardCharsets.UTF_8))));
    }
}
