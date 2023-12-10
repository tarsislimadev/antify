import { Router } from './libs/router/index.js'

const app = new Router()

app.get('/', (_, res) => res.setJSON({ date: Date.now() }))

export default app
