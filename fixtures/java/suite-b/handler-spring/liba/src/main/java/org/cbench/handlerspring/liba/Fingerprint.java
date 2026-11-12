package org.cbench.handlerspring.liba;

import org.cbench.handlerspring.libb.Digests;

public final class Fingerprint {
    private Fingerprint() {
    }

    public static byte[] of(byte[] data) throws Exception {
        return Digests.digest(data);
    }
}
