import os
import asyncio
import time
import requests
from surrealdb import Surreal

max_width = 50

async def main():
    surreal_url = os.getenv("SURREAL_HOST")
    surreal_pass = os.getenv("SURREAL_PWD")
    surreal_user = os.getenv("SURREAL_USER")

    bz_response = requests.get('https://api.hypixel.net/v2/skyblock/bazaar').json()
    total = len(bz_response['products'])
    start_time = time.time()
    i = 0

    print(f'found {total} items on the bazaar')

    async with Surreal(surreal_url) as db:
        await db.signin({"user": surreal_user, "pass": surreal_pass, "NS": "skyblock", "DB": "items"})
        await db.use("skyblock", "items")
        print("connected to database")

        if 'products' in bz_response and isinstance(bz_response['products'], dict):
            for item_key, item_value in bz_response['products'].items():
                name = item_value["product_id"].replace(":", "/")
                await db.create("bz_item", {
                    "id": name,
                    "name": name,
                })
                i+=1
                loading_bar(total, i, item_key)
        else:
            print("Error: 'products' key not found or is not a dictionary")

    print(f'\nfinished in {time.time() - start_time:.2f} seconds')


def loading_bar(total: int, curr: int, curr_name: str):
    curr_fraction = round(max_width / total * curr)
    bar = '\033[92m' + '─' * curr_fraction + '\033[90m' + '─' * (max_width - curr_fraction)
    print(f'\r\033[0m<{bar}\033[0m> {curr}/{total} {curr_name}', end='', flush=True)


if __name__ == "__main__":
    asyncio.run(main())

