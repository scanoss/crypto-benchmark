package org.cbench.wrapper.libb;

import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

public final class Encoder {
    private Encoder() {
    }

    public static byte[] encode(byte[] data) throws NoSuchAlgorithmException {
        return encodeInner(data);
    }

    private static byte[] encodeInner(byte[] data) throws NoSuchAlgorithmException {
        return hashBlock(data);
    }

    private static byte[] hashBlock(byte[] data) throws NoSuchAlgorithmException {
        MessageDigest md = MessageDigest.getInstance("SHA-384");
        return md.digest(data);
    }
}
