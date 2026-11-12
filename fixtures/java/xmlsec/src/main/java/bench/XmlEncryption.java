package bench;

import java.io.ByteArrayInputStream;
import java.nio.charset.StandardCharsets;
import java.security.KeyPair;
import java.security.KeyPairGenerator;
import javax.crypto.KeyGenerator;
import javax.crypto.SecretKey;
import javax.xml.parsers.DocumentBuilderFactory;
import org.apache.xml.security.Init;
import org.apache.xml.security.encryption.EncryptedData;
import org.apache.xml.security.encryption.EncryptedKey;
import org.apache.xml.security.encryption.XMLCipher;
import org.apache.xml.security.keys.KeyInfo;
import org.apache.xml.security.utils.EncryptionConstants;
import org.w3c.dom.Document;
import org.w3c.dom.Element;

public class XmlEncryption {
    private static final String XML = "<patient><name>Jane Doe</name><ssn>123-45-6789</ssn></patient>";

    public static void main(String[] args) throws Exception {
        Init.init();

        KeyGenerator keyGenerator = KeyGenerator.getInstance("AES");
        keyGenerator.init(256);
        SecretKey contentKey = keyGenerator.generateKey();

        KeyPairGenerator kpg = KeyPairGenerator.getInstance("RSA");
        kpg.initialize(3072);
        KeyPair keyPair = kpg.generateKeyPair();

        DocumentBuilderFactory dbf = DocumentBuilderFactory.newInstance();
        dbf.setNamespaceAware(true);
        Document doc = dbf.newDocumentBuilder().parse(new ByteArrayInputStream(XML.getBytes(StandardCharsets.UTF_8)));
        Element ssn = (Element) doc.getElementsByTagName("ssn").item(0);

        XMLCipher keyCipher = XMLCipher.getInstance(XMLCipher.RSA_OAEP);
        keyCipher.init(XMLCipher.WRAP_MODE, keyPair.getPublic());
        EncryptedKey encryptedKey = keyCipher.encryptKey(doc, contentKey);

        XMLCipher xmlCipher = XMLCipher.getInstance(XMLCipher.AES_256_GCM);
        xmlCipher.init(XMLCipher.ENCRYPT_MODE, contentKey);
        EncryptedData encryptedData = xmlCipher.getEncryptedData();
        KeyInfo keyInfo = new KeyInfo(doc);
        keyInfo.add(encryptedKey);
        encryptedData.setKeyInfo(keyInfo);
        xmlCipher.doFinal(doc, ssn, true);

        Element encryptedElement = (Element) doc.getElementsByTagNameNS(
                EncryptionConstants.EncryptionSpecNS, EncryptionConstants._TAG_ENCRYPTEDDATA).item(0);
        XMLCipher decryptCipher = XMLCipher.getInstance();
        decryptCipher.init(XMLCipher.DECRYPT_MODE, null);
        decryptCipher.setKEK(keyPair.getPrivate());
        decryptCipher.doFinal(doc, encryptedElement);

        String recovered = doc.getElementsByTagName("ssn").item(0).getTextContent();
        System.out.println("xml encryption " + recovered);
    }
}
