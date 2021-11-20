import asyncio
from google.cloud import resourcemanager
from google.cloud import bigquery
from typing import List

from google.cloud.resourcemanager_v3.services.folders.pagers import ListFoldersAsyncPager
from google.cloud.resourcemanager_v3.services.projects.pagers import ListProjectsAsyncPager


async def list_all_projects(organization: str = None):
    folder_client = resourcemanager.FoldersAsyncClient()
    project_client = resourcemanager.ProjectsAsyncClient()

    queue = asyncio.Queue()
    queue.put_nowait(organization)
    all_projects = []
    while not queue.empty():
        parent: str = await queue.get()
        print(f"=========  parent: {parent} start =================")
        folder_pager = await folder_client.list_folders(parent=parent)
        project_pager = await project_client.list_projects(parent=parent)

        async for page in folder_pager.pages:
            for f in page.folders:
                print("folder:", f.display_name)
                queue.put_nowait(f.name)

        async for page in project_pager.pages:
            for p in page.projects:
                print("project:", p.display_name)
                all_projects.append({"name": p.name, "display_name": p.display_name, "project_id": p.project_id})
                # queue.put_nowait(p.name)
        print(f"=========  parent: {parent} end =================")
        
    return all_projects

bq_client = bigquery.Client()

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
    print(result)

async def main():
    project_list = await list_all_projects("organizations/1045933413058")
    await print_datasets(project_list=project_list)


if __name__ == "__main__":
    asyncio.run(main())
