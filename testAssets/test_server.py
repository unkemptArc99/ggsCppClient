import asyncio
import websockets

async def server(websocket, path):
  data = await websocket.recv()
  await websocket.send("Thanks for connecting " + data)

start_server = websockets.serve(server, "localhost", 8080)
asyncio.get_event_loop().run_until_complete(start_server)
asyncio.get_event_loop().run_forever()