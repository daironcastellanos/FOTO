import * as functions from 'firebase-functions';
import express from 'express';
import next from 'next';

const dev = process.env.NODE_ENV !== 'production';
const app = next({ dev, conf: { distDir: 'public/.next' } });
const handle = app.getRequestHandler();

const server = express();

server.get('*', (req, res) => {
  return handle(req, res);
});

export const nextApp = functions.https.onRequest(server);

