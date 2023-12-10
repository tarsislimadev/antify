// 

export class Request {
  constructor(str) {
    this.method = this.parserMethod(str)
    this.path = this.parsePath(str)
    this.protocol = this.parseProtocol(str)
    this.headers = this.parseHeaders(str)
    this.body = this.parseBody(str)
  }

  getHeadersAndBody(str) {
    const [headers, body = ''] = str.split('\r\n\r\n', 2)
    return [headers, body]
  }

  getLines(str) {
    return str.split('\r\n')
  }

  getFirstLine(str) {
    const lines = this.getLines(str)
    return lines[0]
  }

  parserMethod(str) {
    const [method,] = this.getFirstLine(str).split(' ')
    return method
  }

  parsePath(str) {
    const [,path] = this.getFirstLine(str).split(' ')
    return path
  }

  parseProtocol(str) {
    const [,,protocol] = this.getFirstLine(str).split(' ')
    return protocol
  }

  parseHeaders(str) {
    const [headers,] = this.getHeadersAndBody(str)
    return headers.split('\r\n')
  }

  parseBody(str) {
    const [,body] = this.getHeadersAndBody(str)
    return body.toString()
  }

}
