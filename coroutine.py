import asyncio

async def say_hello():
    print("hello world")


async def main():
    await say_hello()


#asyncioio.run(main())