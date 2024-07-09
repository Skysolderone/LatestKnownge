import org.springframework.util.Base64Utils;
import javax.crypto.Cipher;
import java.nio.charset.StandardCharsets;
import java.security.KeyFactory;
import java.security.PrivateKey;
import java.security.PublicKey;
import java.security.Signature;
import java.security.spec.PKCS8EncodedKeySpec;
import java.security.spec.X509EncodedKeySpec;
import java.util.HashMap;
import java.util.Iterator;
import java.util.Map;
import java.util.Set;
import java.util.TreeSet;
public class RsaUtil {
// Perform RSA encryption using the public key.
public static String encryptRSA(String plaintext, String publicKey) throws
Exception {
Cipher cipher = Cipher.getInstance("RSA");
cipher.init(Cipher.ENCRYPT_MODE, getPublicKey(publicKey));
byte[] encryptedBytes =
cipher.doFinal(plaintext.getBytes(StandardCharsets.UTF_8));
return Base64Utils.encodeToString(encryptedBytes);
}
// Perform RSA decryption using the private key.
public static String decryptRSA(String encryptedData, String privateKey)
throws Exception {
Cipher cipher = Cipher.getInstance("RSA");
cipher.init(Cipher.DECRYPT_MODE, getPrivateKey(privateKey));
byte[] encryptedBytes = Base64Utils.decodeFromString(encryptedData);
byte[] decryptedBytes = cipher.doFinal(encryptedBytes);
return new String(decryptedBytes, StandardCharsets.UTF_8);
}
public static PublicKey getPublicKey(String key) throws Exception {
byte[] byteKey =
org.apache.commons.codec.binary.Base64.decodeBase64(key);
X509EncodedKeySpec x509EncodedKeySpec = new
X509EncodedKeySpec(byteKey);
KeyFactory keyFactory = KeyFactory.getInstance("RSA");
return keyFactory.generatePublic(x509EncodedKeySpec);
}

public static PrivateKey getPrivateKey(String key) throws Exception {
byte[] byteKey =
org.apache.commons.codec.binary.Base64.decodeBase64(key);
PKCS8EncodedKeySpec x509EncodedKeySpec = new
PKCS8EncodedKeySpec(byteKey);
KeyFactory keyFactory = KeyFactory.getInstance("RSA");
return keyFactory.generatePrivate(x509EncodedKeySpec);
}
public static boolean verify(Map<String, String> map, String publicKey,
String sign) throws Exception {
String srcData = getSignData(map);
byte[] keyBytes = getPublicKey(publicKey).getEncoded();
X509EncodedKeySpec keySpec = new X509EncodedKeySpec(keyBytes);
KeyFactory keyFactory = KeyFactory.getInstance("RSA");
PublicKey key = keyFactory.generatePublic(keySpec);
Signature signature = Signature.getInstance("MD5withRSA");
signature.initVerify(key);
signature.update(srcData.getBytes());
return signature.verify(Base64Utils.decodeFromString(sign));
}
public static String sign(String data, String privateKey) throws Exception
{
byte[] keyBytes = getPrivateKey(privateKey).getEncoded();
PKCS8EncodedKeySpec keySpec = new PKCS8EncodedKeySpec(keyBytes);
KeyFactory keyFactory = KeyFactory.getInstance("RSA");
PrivateKey key = keyFactory.generatePrivate(keySpec);
Signature signature = Signature.getInstance("MD5withRSA");
signature.initSign(key);
signature.update(data.getBytes());
return Base64Utils.encodeToString(signature.sign());
}
public static String getSignData(Map<String, String> params) {
StringBuilder sb = new StringBuilder();
//step1：Sort the parameters
Set<String> keyset = params.keySet();
TreeSet<String> sortSet = new TreeSet<String>();
sortSet.addAll(keyset);
Iterator<String> it = sortSet.iterator();
//step2：Concatenate the keys and values of the parameters
while (it.hasNext()) {
String key = it.next();
String value = params.get(key);
sb.append(key).append(value);
}
return sb.toString();
}
static String publicKey =
"MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0noWhK3+WDdWGh+x1F3T\n" +
"oCuCy7v7KL7gJP4Ir75UJpygkaTxZLybEemmjWB5HQUK0o7G2/ht2YVZiusL4dBP\n" +
"GnRSzILt0ZKMJduzpbsjWW/eZNgb+/INRkifyiS1WYOMnnlsbA3qapFqKFL2KhwU\n" +
"S6vulUS/hYcUrkQejLafp/S1v24pIMD0V0tagpUcD/CnNKPXu/9AcXIl5xXeEpl1\n" +
"GHDn3kMhyuM3hcXhojmoMb6nuAgwGK9zntK8Ip9/ZVmmoepcBKkFv4E3pnZy3ASg\n" +
"vs2v83y/eAhOkUhV+meigk00vYFk464nlLqUN1tHBAnpKptHHCUFlg4CWBnq3v9Q\n" +
"ywIDAQAB";
static String privateKey =
"MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDSehaErf5YN1Ya\n" +
"H7HUXdOgK4LLu/sovuAk/givvlQmnKCRpPFkvJsR6aaNYHkdBQrSjsbb+G3ZhVmK\n" +
"6wvh0E8adFLMgu3Rkowl27OluyNZb95k2Bv78g1GSJ/KJLVZg4yeeWxsDepqkWoo\n" +
"UvYqHBRLq+6VRL+FhxSuRB6Mtp+n9LW/bikgwPRXS1qClRwP8Kc0o9e7/0BxciXn\n" +
"Fd4SmXUYcOfeQyHK4zeFxeGiOagxvqe4CDAYr3Oe0rwin39lWaah6lwEqQW/gTem\n" +
"dnLcBKC+za/zfL94CE6RSFX6Z6KCTTS9gWTjrieUupQ3W0cECekqm0ccJQWWDgJY\n" +
"Gere/1DLAgMBAAECggEBAL4E6RWJ/Alqk/pryOFgf/GntqL2qG6FvNjI29DGNTTS\n" +
"zQgQcLqwqBNI7UvP4Cf56GZl7lSjeZEbGdcYExcgOHo6sJHgVpKRCqQXMJ4cUHKr\n" +
"U42ZqdIRGjjs0g3ujCcvY6GjH+aBvbPOapfTJlioFw6mP3RVp88GP/6Ak1AYhBVd\n" +
"bZ32KI94vuNPbsrwQHlGk0hKpVh0eAeosAbKNIEurultPysNbxm7dYUY6VPtqljR\n" +
"KooB1u37RTSEI+fiBYmjS4b9j4ynGfT0yGG1gYj+ytiJRUPDkoiTT8cCADjona2N\n" +
"qR5LObGKShFmj7SN8wPyNJPTn+JMx5wM7nutiBH7orkCgYEA8/FXKYdC+0dKc6v6\n" +
"qkZf6XA2j2RgdAk8OJwYTA9JjxWX0tgFmFKKaSN9yXihGes441BwL04wi5wvHNw7\n" +
"FUcUSRjfx0Uu8HOzNpqdMVQPQfX1lTDbpm9eHRvjDdcLHlOTCekYQOGVmuE1H3Zw\n" +
"hmLGfgm3flP799Hmaqrc8/glzZcCgYEA3OFMFLYZYQm+5IAQZe8klEC5xa7ccHsO\n" +
"CA0+HqJvh+U+Zy74PCjFfz29jyGQxmz/eygWr5VQoHA4jQmlKqjTpj2VKgCLOjdF\n" +
"aXGoQS/mOjR0M3hZPPHzOzbHsl09p8/96rC9VsiQOF2oqYoNGHr+DwH++qXYsrGI\n" +
"zIin0HX3ZO0CgYBhY1Rc6/c9wjRjAaHNINNhqr7deEFZkeZM42R/2QqQ6Rn3vu5c\n" +
"5XcEinrJWDNY2aOYfKCNAjY3Rl84smOUFxBuLlQIhgI7VLWTcx6WduywdLVanmrS\n" +
"g1ubW2rGN7fkn5DwP/LC1EyZzJccvHgn7n84CkELWhQZ196ZbyVO8R6GeQKBgBD4\n" +
"MispLlv5LrRJbnkWXV4SdvNMEt2FZreRpOMfoaf0Ic41mpasnze+W5DiiEfmWd5x\n" +
"XwTQWOhqlr2nLwxO+iu8cXhPoGKxmmCWfdG8R1jTbNYDef+nqwMymzcF2NXKsfxU\n" +
"5ccEE6hw8aNM2uK1mE043wkMstBPuW3VVec7GO2lAoGAGjm9PSVVI4Idex9BPAR+\n" +
"m4JP949yKh3r0CZOxzAktfLgoal0TqF+LlJ69sKwceMsPdf3wmy2mfDWyF0Qza5L\n" +
"6KSs5irAAUBxwFDo7FKGJu6bPxQAryB2rMYzvScZrGT2ktNjoVPIqaxqNNxXK1PD\n" +
"G+SIwgQWs0CoST8pGsLjXyo=";}

public static String getSignData(Map<String, String> params) {
StringBuilder sb = new StringBuilder();
//step1：Sort the parameters
Set<String> keyset = params.keySet();
TreeSet<String> sortSet = new TreeSet<String>();
sortSet.addAll(keyset);
Iterator<String> it = sortSet.iterator();
//step2：Concatenate the keys and values of the parameters
while (it.hasNext()) {
String key = it.next();
String value = params.get(key);
sb.append(key).append(value);
}
return sb.toString();
}

public static PrivateKey getPrivateKey(String key) throws Exception {
byte[] byteKey =
org.apache.commons.codec.binary.Base64.decodeBase64(key);
PKCS8EncodedKeySpec x509EncodedKeySpec = new
PKCS8EncodedKeySpec(byteKey);
KeyFactory keyFactory = KeyFactory.getInstance("RSA");
return keyFactory.generatePrivate(x509EncodedKeySpec);
}

public static String sign(String data, String privateKey) throws Exception
{
byte[] keyBytes = getPrivateKey(privateKey).getEncoded();
PKCS8EncodedKeySpec keySpec = new PKCS8EncodedKeySpec(keyBytes);
KeyFactory keyFactory = KeyFactory.getInstance("RSA");
PrivateKey key = keyFactory.generatePrivate(keySpec);
Signature signature = Signature.getInstance("MD5withRSA");
signature.initSign(key);
signature.update(data.getBytes());
return Base64Utils.encodeToString(signature.sign());
}
static String privateKey =
"MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDSehaErf5YN1Ya\n" +
"H7HUXdOgK4LLu/sovuAk/givvlQmnKCRpPFkvJsR6aaNYHkdBQrSjsbb+G3ZhVmK\n" +
"6wvh0E8adFLMgu3Rkowl27OluyNZb95k2Bv78g1GSJ/KJLVZg4yeeWxsDepqkWoo\n" +
"UvYqHBRLq+6VRL+FhxSuRB6Mtp+n9LW/bikgwPRXS1qClRwP8Kc0o9e7/0BxciXn\n" +
"Fd4SmXUYcOfeQyHK4zeFxeGiOagxvqe4CDAYr3Oe0rwin39lWaah6lwEqQW/gTem\n" +
"dnLcBKC+za/zfL94CE6RSFX6Z6KCTTS9gWTjrieUupQ3W0cECekqm0ccJQWWDgJY\n" +
"Gere/1DLAgMBAAECggEBAL4E6RWJ/Alqk/pryOFgf/GntqL2qG6FvNjI29DGNTTS\n" +
"zQgQcLqwqBNI7UvP4Cf56GZl7lSjeZEbGdcYExcgOHo6sJHgVpKRCqQXMJ4cUHKr\n" +
"U42ZqdIRGjjs0g3ujCcvY6GjH+aBvbPOapfTJlioFw6mP3RVp88GP/6Ak1AYhBVd\n" +
"bZ32KI94vuNPbsrwQHlGk0hKpVh0eAeosAbKNIEurultPysNbxm7dYUY6VPtqljR\n" +
"KooB1u37RTSEI+fiBYmjS4b9j4ynGfT0yGG1gYj+ytiJRUPDkoiTT8cCADjona2N\n" +
"qR5LObGKShFmj7SN8wPyNJPTn+JMx5wM7nutiBH7orkCgYEA8/FXKYdC+0dKc6v6\n" +
"qkZf6XA2j2RgdAk8OJwYTA9JjxWX0tgFmFKKaSN9yXihGes441BwL04wi5wvHNw7\n" +
"FUcUSRjfx0Uu8HOzNpqdMVQPQfX1lTDbpm9eHRvjDdcLHlOTCekYQOGVmuE1H3Zw\n" +
"hmLGfgm3flP799Hmaqrc8/glzZcCgYEA3OFMFLYZYQm+5IAQZe8klEC5xa7ccHsO\n" +
"CA0+HqJvh+U+Zy74PCjFfz29jyGQxmz/eygWr5VQoHA4jQmlKqjTpj2VKgCLOjdF\n" +
"aXGoQS/mOjR0M3hZPPHzOzbHsl09p8/96rC9VsiQOF2oqYoNGHr+DwH++qXYsrGI\n" +
"zIin0HX3ZO0CgYBhY1Rc6/c9wjRjAaHNINNhqr7deEFZkeZM42R/2QqQ6Rn3vu5c\n" +
"5XcEinrJWDNY2aOYfKCNAjY3Rl84smOUFxBuLlQIhgI7VLWTcx6WduywdLVanmrS\n" +
"g1ubW2rGN7fkn5DwP/LC1EyZzJccvHgn7n84CkELWhQZ196ZbyVO8R6GeQKBgBD4\n" +
"MispLlv5LrRJbnkWXV4SdvNMEt2FZreRpOMfoaf0Ic41mpasnze+W5DiiEfmWd5x\n" +
"XwTQWOhqlr2nLwxO+iu8cXhPoGKxmmCWfdG8R1jTbNYDef+nqwMymzcF2NXKsfxU\n" +
"5ccEE6hw8aNM2uK1mE043wkMstBPuW3VVec7GO2lAoGAGjm9PSVVI4Idex9BPAR+\n" +
"m4JP949yKh3r0CZOxzAktfLgoal0TqF+LlJ69sKwceMsPdf3wmy2mfDWyF0Qza5L\n" +
"6KSs5irAAUBxwFDo7FKGJu6bPxQAryB2rMYzvScZrGT2ktNjoVPIqaxqNNxXK1PD\n" +
"G+SIwgQWs0CoST8pGsLjXyo=";
public static void main(String[] args) throws Exception {
// Generate sign start
Map map = new HashMap();
long l = System.currentTimeMillis();
map.put("timestamp", String.valueOf(l));
map.put("clientId", "1");
map.put("clientUserId", "123");
String sign = sign(getSignData(map), privateKey);
System.out.println(sign);
// Generate sign end
}







