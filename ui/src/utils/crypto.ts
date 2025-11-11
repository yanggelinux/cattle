const sharedKeyBase64 = 'MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODA5MTI=' // 后端同一份 key
const encoder = new TextEncoder()
const decoder = new TextDecoder()

async function getCryptoKey(): Promise<CryptoKey> {
  const rawKey = Uint8Array.from(atob(sharedKeyBase64), (c) => c.charCodeAt(0))
  return await crypto.subtle.importKey('raw', rawKey, { name: 'AES-GCM' }, false, [
    'encrypt',
    'decrypt',
  ])
}

export async function encryptPassword(password: string): Promise<{ iv: string; cipher: string }> {
  const key = await getCryptoKey()
  const iv = crypto.getRandomValues(new Uint8Array(12)) // GCM 标准使用 12 字节 IV
  const encodedPassword = encoder.encode(password)

  const ciphertext = await crypto.subtle.encrypt({ name: 'AES-GCM', iv: iv }, key, encodedPassword)

  return {
    iv: btoa(String.fromCharCode(...iv)),
    cipher: btoa(String.fromCharCode(...new Uint8Array(ciphertext))),
  }
}

export async function decryptPassword(ivB64: string, cipherB64: string): Promise<string> {
  const key = await getCryptoKey()
  const iv = Uint8Array.from(atob(ivB64), (c) => c.charCodeAt(0))
  const ciphertext = Uint8Array.from(atob(cipherB64), (c) => c.charCodeAt(0))

  const decrypted = await crypto.subtle.decrypt({ name: 'AES-GCM', iv: iv }, key, ciphertext)

  return decoder.decode(decrypted)
}
