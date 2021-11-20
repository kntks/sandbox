import asyncio
import time

from main import stop_watch


async def sync_sleep(x: int):
    time.sleep(x)


# time.sleepとasyncio.sleepの比較
@stop_watch
def time_sleep():
    loop = asyncio.get_event_loop()
    tasks = [loop.create_task(sync_sleep(x)) for x in range(5)]
    loop.run_until_complete(asyncio.wait(tasks))


@stop_watch
def asyncio_sleep():
    loop = asyncio.get_event_loop()
    tasks = [loop.create_task(asyncio.sleep(x)) for x in range(5)]
    loop.run_until_complete(asyncio.wait(tasks))


def demo1():
    print("demo1: time.sleep and asyncio.sleep")
    time_sleep()
    asyncio_sleep()


if __name__ == "__main__":
    demo1()
