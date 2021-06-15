import asyncio
import logging
import ssl
import pathlib

from nats.aio.client import Client as NATS


async def run():
    logging.debug("RUN: Start")

    certs_path = pathlib.Path.home() / 'fwcerts'

    ssl_ctx = ssl.create_default_context(purpose=ssl.Purpose.SERVER_AUTH)
    ssl_ctx.load_cert_chain(
        certfile=certs_path / 'server.crt',
        keyfile=certs_path / 'server.key',
    )

    closed = asyncio.Future()

    async def handle_msg(msg):
        nonlocal closed
        logging.info('Msg (%s, %s)', msg.subject, msg.data.decode())
        if msg.subject == 'chat.close':
            closed.set_result(True)

    nc = NATS()
    await nc.connect("nats://127.0.0.1:4222", tls=ssl_ctx, tls_hostname='massimo-mbp.fwx.one')

    await nc.subscribe('chat.*', cb=handle_msg)

    await asyncio.wait_for(closed, None)

    logging.debug("RUN: Stop")

    await nc.drain()


if __name__ == '__main__':
    logging.basicConfig(
        level=logging.DEBUG
    )
    logging.info("hello")
    
    asyncio.run(run())
