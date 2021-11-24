import time
import asyncio
from google.cloud import resourcemanager
from google.cloud import bigquery
from typing import List

from google.cloud.resourcemanager_v3.services.folders.pagers import ListFoldersAsyncPager
from google.cloud.resourcemanager_v3.services.projects.pagers import ListProjectsAsyncPager


async def worker(name, folder_queue: asyncio.Queue, project_queue: asyncio.Queue):
    folder_client = resourcemanager.FoldersAsyncClient()
    project_client = resourcemanager.ProjectsAsyncClient()
    while True:
        parent: str = await folder_queue.get()

        folder_pager = await folder_client.list_folders(parent=parent)
        async for page in folder_pager.pages:
            for f in page.folders:
                print("folder:", f.display_name)
                folder_queue.put_nowait(f.name)

        project_pager = await project_client.list_projects(parent=parent)
        async for page in project_pager.pages:
            for p in page.projects:
                print("project:", p.display_name)
                project_queue.put_nowait({"name": p.name, "display_name": p.display_name, "project_id": p.project_id})
       
        folder_queue.task_done()

        # queue.put_nowait(1.0)
        # print(f'{name} has slept for {sleep_for:.2f} seconds')
        print(f"====={name} end =======")

async def main(organization_id: str):
    folder_queue = asyncio.Queue()
    project_queue = asyncio.Queue()

    await folder_queue.put(organization_id)
    tasks = []
    for i in range(3):
        task = asyncio.create_task(worker(f'worker-{i}', folder_queue, project_queue))
        tasks.append(task)

    await folder_queue.join()
  
    for task in tasks:
        task.cancel()
    await asyncio.gather(*tasks, return_exceptions=True)


if __name__ == "__main__":
  organization_id = "organizations/"
  started_at = time.monotonic()
  asyncio.run(main(organization_id))
  total_slept_for = time.monotonic() - started_at
  print(f'{total_slept_for:.2f} seconds')
  
  