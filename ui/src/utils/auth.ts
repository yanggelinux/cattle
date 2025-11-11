// 访问 token 缓存的 key

function getAccessToken(): string {
  return localStorage.getItem('token') || ''
}

function getAuthorization(): string {
  return localStorage.getItem('authorization') || ''
}

function getIsSuper(): string {
  return localStorage.getItem('isSuper') || '0'
}

function getUserID(): number {
  return localStorage.userID ? Number(localStorage.userID) : 0
}

function getUserName(): string {
  return localStorage.getItem('userName') || ''
}

function getDeptName(): string {
  return localStorage.getItem('deptName') || ''
}

function getRoleDisplayNames(): string {
  return localStorage.getItem('roleDisplayName') || ''
}
function setAccessToken(token: string) {
  localStorage.setItem('token', token)
}
function clearToken() {
  localStorage.removeItem('token')
}

export {
  getAccessToken,
  getAuthorization,
  setAccessToken,
  clearToken,
  getUserID,
  getUserName,
  getDeptName,
  getRoleDisplayNames,
  getIsSuper,
}
