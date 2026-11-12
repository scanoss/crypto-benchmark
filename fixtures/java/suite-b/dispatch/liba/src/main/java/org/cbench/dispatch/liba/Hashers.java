package org.cbench.dispatch.liba;

import org.cbench.dispatch.libb.Digests;

public final class Hashers {
    private Hashers() {
    }

    public static Hasher standard() {
        return new Standard();
    }

    public static Hasher plain() {
        return new Plain();
    }

    private static final class Standard implements Hasher {
        @Override
        public byte[] sum(byte[] data) throws Exception {
            return Digests.digest(data);
        }
    }

    private static final class Plain implements Hasher {
        @Override
        public byte[] sum(byte[] data) {
            return data.clone();
        }
    }
}
