import asyncio
import time
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


bq_client = bigquery.Client()
import itertools
# https://googleapis.dev/python/bigquery/latest/generated/google.cloud.bigquery.dataset.DatasetListItem.html#google.cloud.bigquery.dataset.DatasetListItem
# https://googleapis.dev/python/bigquery/latest/generated/google.cloud.bigquery.dataset.Dataset.html#google.cloud.bigquery.dataset.Dataset
async def list_datasets(project: str = None):
    dataset_list = []
    for dataset_list_item in bq_client.list_datasets(project=project):
        dataset_list.append(bq_client.get_dataset(dataset_list_item.reference))
    return dataset_list


async def print_datasets(project_list):
    tasks = [asyncio.create_task(list_datasets(p["project_id"])) for p in project_list]
    result = await asyncio.gather(*tasks)
    print(list(itertools.chain.from_iterable(result)))

async def main(organization_id: str):
    folder_queue = asyncio.Queue()
    project_queue = asyncio.Queue()

    await folder_queue.put(organization_id)
    tasks = []
    for i in range(3):
        task = asyncio.create_task(worker(f'worker-{i}', folder_queue, project_queue))
        tasks.append(task)
    
    for task in tasks:
        task.cancel()
    # すべてのワーカータスクがキャンセルされるまで待ちます。
    await asyncio.gather(*tasks, return_exceptions=True)
    # await print_datasets(project_list=project_list)

    dataset_tasks = [asyncio.create_task(list_datasets())]

if __name__ == "__main__":
    organization_id = "organizations/1045933413058"
    started_at = time.monotonic()
    asyncio.run(main())
    total_slept_for = time.monotonic() - started_at
    print(f'{total_slept_for:.2f} seconds')
    
