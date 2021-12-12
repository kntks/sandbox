from google.cloud import securitycenter

organization_id=
client = securitycenter.SecurityCenterClient()

project_filter = (
    "security_center_properties.resource_type= \"google.cloud.bigquery.Dataset\""
)
asset_iterator = client.list_assets(
    request={"parent": f"organizations/{organization_id}", "filter": project_filter}
)

from itertools import groupby

sort = sorted(asset_iterator, key=lambda x: x.asset.resource_properties["location"])
for group, dataset in groupby(sort, lambda x: x.asset.resource_properties["location"]):
    print(f"======== {group} ======")
    for d in dataset:
        print(d.asset.resource_properties["id"])
