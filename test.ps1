# param(
# 	$resName,
# 	$groupName,
# 	$clusterNetwork,
# 	$cluster
# )
$ErrorActionPreference = [System.Management.Automation.ActionPreference]::Stop

$resName = "default%example-lb"
$groupName = "g2"
$clusterNetwork = "Cluster Network 2"
$cluster = "10.231.121.132"
# $vipToClone = "Cluster IP Address"

# $cVip = get-clusterresource $vipToClone -cluster $cluster
# $cVipParam = $cVip | get-clusterparameter -cluster $cluster
write-information $resName
write-information $cluster
write-information $clusterNetwork 
write-information $groupName 

try{
	$r = Add-ClusterResource $resName -cluster $cluster -group $groupName -Type "IP Address" 
	}
	catch{

echo $_
echo "failed"

	
	}
$r = get-clusterresource $resName -cluster $cluster
$r.description = "Load Balancer"
#cluster will fix up the subnet :-)
$r |set-clusterparameter  -cluster $cluster  -Multiple @{"Network"=$clusterNetwork;"EnableDhcp"=1} -ErrorAction SilentlyContinue

$r = $r | start-clusterresource -wait 30 -cluster $cluster
$r = $r | stop-clusterresource -wait 30 -cluster $cluster


($r | get-clusterparameter -cluster $cluster -name "Address").Value
