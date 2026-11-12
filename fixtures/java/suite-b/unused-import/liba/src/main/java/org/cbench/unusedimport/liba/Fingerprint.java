package org.cbench.unusedimport.liba;

import org.cbench.unusedimport.libb.Digests;

public final class Fingerprint {
    public static final String VERSION = "1.0.0";

    private Fingerprint() {
    }

    public static byte[] of(byte[] data) throws Exception {
        return Digests.digest(data);
    }
}
