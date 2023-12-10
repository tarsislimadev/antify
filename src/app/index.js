import netPkg from 'net'
import { Request, Response } from './libs/http/index.js'
import { PORT } from './config.js'

import app from './app.js'

const server = netPkg.createServer((socket) => {
  let lastData = Date.now()

  socket.on('data', (data) => {
    lastData = Date.now()
    const dataStr = data.toString()
    const req = new Request(dataStr)
    const res = new Response(req)
    socket.write(app.run(req, res).toString())
  })

  setInterval(() => Date.now() - lastData > 1000 ? socket.end() : null, 1000)
})

server.listen(PORT, () => console.log(`listening on port ${PORT}`))
