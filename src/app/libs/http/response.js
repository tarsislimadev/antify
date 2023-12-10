import { Request } from './request.js'

import messages from './response.messages.js'

export class Response {
  status = '200'
  content_type = 'text/plain'
  headers = new Headers()
  body = ''

  request = null

  constructor(request = new Request('')) {
    this.request = request
  }

  getStatus() {
    return this.status
  }

  getStatusMessage() {
    const message = messages[this.getStatus()]

    if (!message) return 'Error'

    return message
  }

  getProtocol() {
    return this.request.protocol
  }

  getFirstLine() {
    const first = []
    first.push(this.getProtocol())
    first.push(this.getStatus())
    first.push(this.getStatusMessage())
    return first.join(' ')
  }

  getBodyString() {
    return this.body.toString()
  }

  setJSON(json = {}, status = '200') {
    this.status = status
    this.headers.set('Content-Type', 'application/json')
    this.body = JSON.stringify(json)
    return this
  }

  toString() {
    return [
      this.getFirstLine(),
      // this.headers.toString(),
      '',
      this.getBodyString(),
      '',
    ].join('\r\n')
  }
}
