import asyncio
from os import wait
from main import stop_watch
import random


async def http_call(delay: int):
    await asyncio.sleep(delay)
    return f"{delay} sec delay"


async def gather():
    print("asyncio.gather start")
    delay_list = [random.randint(1, 4) for _ in range(4)]
    print("delay_list:", delay_list)
    calls = [http_call(x) for x in delay_list]
    L = await asyncio.gather(*calls)
    print(L)
    print("asyncio.gather end")


async def wait():
    print("asyncio.wait start")
    delay_list = [random.randint(1, 4) for _ in range(4)]
    print("delay_list:", delay_list)
    loop = asyncio.get_event_loop()
    tasks = [loop.create_task(http_call(x)) for x in delay_list]
    done, _ = await asyncio.wait(tasks)

    # delay_listとtask.result()の順序は保証されない
    for task in done:
        print(task.result())
    print("asyncio.wait end")


async def as_completed():
    print("asyncio.as_completed start")
    for coro in asyncio.as_completed([http_call(x) for x in range(1, 4)]):
        print(await coro)

    print("asyncio.as_completed end")


async def demo2_main():
    await gather()
    await wait()
    await as_completed()


@stop_watch
def demo2():
    asyncio.run(demo2_main())


if __name__ == "__main__":
    demo2()
