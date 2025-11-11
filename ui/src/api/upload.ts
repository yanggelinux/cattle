import request from '@/utils/request'

const BASE_URL = '/api/v1/upload'

const UploadAPI = {
  uploadFile(uploadData: any) {
    return request({
      url: `${BASE_URL}/file`,
      method: 'post',
      data: uploadData,
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    })
  },
}

export default UploadAPI

export interface UploadFileResult {
  fileName: string
  fileUrl: string
}

export interface UploadFileInfo {
  fieldKey: string
  fieldVal: string
  groupName: string
  idx: number
}
