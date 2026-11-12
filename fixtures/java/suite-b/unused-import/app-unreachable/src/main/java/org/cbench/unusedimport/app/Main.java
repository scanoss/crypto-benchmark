package org.cbench.unusedimport.app;

import java.nio.charset.StandardCharsets;
import java.util.HexFormat;
import org.cbench.unusedimport.liba.Fingerprint;

public final class Main {
    private Main() {
    }

    public static void main(String[] args) throws Exception {
        System.out.println(Fingerprint.VERSION);
    }
}
