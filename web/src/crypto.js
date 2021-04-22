import crypto from 'crypto-js';

export const encrypt = (message) => {
  return crypto.AES.encrypt(
    message,
    process.env.REACT_APP_SECRET,
  ).toString();
}

export const decrypt = (encrypted) => {
  try {
    var bytes  = crypto.AES.decrypt(encrypted, process.env.REACT_APP_SECRET);
    return bytes.toString(crypto.enc.Utf8);
  } catch (err) {
    return {};
  }
}
