package org.cbench.transitive.liba;

import org.cbench.transitive.libb.Digests;

public final class Fingerprint {
    private Fingerprint() {
    }

    public static byte[] of(byte[] data) throws Exception {
        return Digests.digest(data);
    }
}
